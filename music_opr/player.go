package main

import (
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

var targetFormat = beep.Format{
	SampleRate:  44100,
	NumChannels: 2,
	Precision:   2,
}

/*
IterFunc ：定义了一个名为IterFunc的新类型，它是一个函数类型。
在Go中，函数本身也可以被当作类型来使用，这允许我们创建函数类型的变量，这些变量可以存储和传递函数。
用于后续的音频迭代
*/
//type IterFunc func() beep.Streamer

// Player 定义一个播放器
type Player struct {
	ctrl          *beep.Ctrl        // 控制播放
	volume        *effects.Volume   // 音量控制
	streamer      beep.Streamer     // 当前音频流
	currentStream beep.StreamSeeker // 当前音频流的位置
}

// 创建一个播放器
func newPlayer() *Player {
	p := &Player{}
	return p.reset()
}

// 下一首歌的切换逻辑：随机-顺序-循环
func nextSong(currentIndex *int) beep.StreamSeekCloser {
	audioFiles := []string{
		"resources/sound-sculptors.mp3",
		"resources/瑶山遗韵.mp3", // 02:14
		"resources/霞据云佩.mp3", // 02:11
	}

	// 这个函数每次被调用时，都会尝试加载列表中的下一个音频文件
	if *currentIndex >= len(audioFiles)-1 {
		// 如果没有更多的文件，将currentIndex置为-1
		*currentIndex = -1
	}

	*currentIndex++

	// 打开当前索引的音频文件
	file, err := os.Open(audioFiles[*currentIndex])
	if err != nil {
		log.Printf("Failed to open audio file: %v", err)
		return nil
	}

	// 解码音频文件并返回streamer
	streamer, _, err := mp3.Decode(file)
	if err != nil {
		log.Printf("Failed to decode audio file: %v", err)
		return nil
	}
	fmt.Printf("playing... %v \n", strings.Split(audioFiles[*currentIndex], "/")[1])

	return streamer

}

// 播放器切歌逻辑
func (p *Player) changeSong(currentIndex *int) {
	speaker.Clear()

	// 拿到下一次的streamer
	steamer := nextSong(currentIndex)

	// 更新currentStream
	p.currentStream = steamer

	// 更新streamer
	p.streamer = steamer

	p.ctrl = &beep.Ctrl{Streamer: p.streamer}

	speaker.Play(p.ctrl)
}

// 重置播放器：清除当前的音频流并重置streamer。
func (p *Player) reset() *Player {
	speaker.Clear()
	p.streamer = nil
	p.ctrl = nil
	return p
}

// 开启播放器：open方法初始化音频输出设备，并开始播放音频。
func (p *Player) open() *Player {
	err := speaker.Init(targetFormat.SampleRate, targetFormat.NumChannels)
	if err != nil {
		return nil
	}
	speaker.Play(p.ctrl)
	return p
}

// 关闭播放器：停止播放并关闭音频输出设备。
func (p *Player) close() *Player {
	speaker.Clear()
	return p
}

// 控制音乐播放和暂停
func (p *Player) togglePlay() {
	speaker.Lock()
	p.ctrl.Paused = !p.ctrl.Paused
	speaker.Unlock()
}

// 解码一个mp3文件，设置player当前的参数
func (p *Player) playMp3(file io.ReadCloser) beep.Streamer {
	streamer, _, err := mp3.Decode(file)
	if err != nil {
		return nil
	}

	p.streamer = streamer      // 如果没设置会runtime error: invalid memory address or nil pointer dereference
	p.currentStream = streamer // 方便记录currentPosition()
	p.ctrl = &beep.Ctrl{Streamer: p.streamer}
	return streamer
}

// 获取播放的当前时间
func (p *Player) currentPosition() string {
	speaker.Lock()
	pos := targetFormat.SampleRate.D(p.currentStream.Position())
	totalPos := targetFormat.SampleRate.D(p.currentStream.Len())
	speaker.Unlock()
	minutes := pos / time.Minute
	second := pos % time.Minute / time.Second

	totalMinute := totalPos / time.Minute
	totalSecond := totalPos % time.Minute / time.Second
	return fmt.Sprintf("%02d:%02d / %02d:%02d", minutes, second, totalMinute, totalSecond)
}

// 当前音乐是否播放完
func (p *Player) isDone() bool {
	return p.currentStream.Position() == p.currentStream.Len()
}
