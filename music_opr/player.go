package main

import (
	"bufio"
	"example.com/m/entity/enum"
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var targetFormat = beep.Format{
	SampleRate:  beep.SampleRate(44100),
	NumChannels: 2,
	Precision:   2,
}

var allSongList []string

func getAllSongList() []string {
	fileHandle, err := os.OpenFile("resources/songList.txt", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal("读取songList文件错误")
	}

	defer fileHandle.Close()

	reader := bufio.NewReader(fileHandle)

	var results []string
	// 按行处理txt
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		results = append(results, string(line))
	}
	return results
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
	playLogic     int               // 播放逻辑：0-顺序播放，1-随机播放，2-单曲循环
}

func getRandomIndex() int {
	return rand.Intn(len(allSongList))
}

// NewPlayer 创建一个播放器
func NewPlayer() *Player {
	p := &Player{}
	allSongList = getAllSongList()
	p.playLogic = enum.ORDER
	return p.reset()
}

// 下一首歌的切换逻辑：随机-顺序-循环
func (p *Player) nextSong(currentIndex *int, isDone int) (beep.StreamSeekCloser, string) {

	// 这个函数每次被调用时，都会尝试加载列表中的下一个音频文件
	if *currentIndex >= len(allSongList)-1 {
		// 如果没有更多的文件，将currentIndex置为-1
		*currentIndex = -1
	}

	// 不是LOOP情况被动切歌才+1
	if isDone == 0 {
		*currentIndex++
	}

	// 打开当前索引的音频文件
	file, err := os.Open(allSongList[*currentIndex])
	if err != nil {
		log.Printf("Failed to open audio file: %v", err)
		return nil, ""
	}

	// 解码音频文件并返回streamer
	streamer, _, err := mp3.Decode(file)
	if err != nil {
		log.Printf("Failed to decode audio file: %v", err)
		return nil, ""
	}

	return streamer, strings.Split(allSongList[*currentIndex], "/")[1]

}

// 上一首的切歌逻辑
func (p *Player) previousSong(currentIndex *int) (beep.StreamSeekCloser, string) {
	// 这个函数每次被调用时，都会尝试加载列表中的下一个音频文件
	if *currentIndex == 0 {
		// 如果没有上一首，将currentIndex置为len(allSongList的长度)
		*currentIndex = len(allSongList)
	}

	*currentIndex--

	// 打开当前索引的音频文件
	file, err := os.Open(allSongList[*currentIndex])
	if err != nil {
		log.Printf("Failed to open audio file: %v", err)
		return nil, ""
	}

	// 解码音频文件并返回streamer
	streamer, _, err := mp3.Decode(file)
	if err != nil {
		log.Printf("Failed to decode audio file: %v", err)
		return nil, ""
	}

	return streamer, strings.Split(allSongList[*currentIndex], "/")[1]
}

// 播放器切歌逻辑
func (p *Player) changeSong(currentIndex *int, changeLogic int) {
	speaker.Clear()

	var streamer beep.StreamSeekCloser
	songName := ""
	// 拿到下一首的streamer

	// RANDOM就对currentIndex随机
	if p.playLogic == enum.RANDOM {
		*currentIndex = getRandomIndex()
	}
	if changeLogic == 0 {
		streamer, songName = p.nextSong(currentIndex, 0)
	} else if changeLogic == 1 {
		streamer, songName = p.previousSong(currentIndex)
	} else if changeLogic == 3 {
		// LOOP情况下被动切换下一首：触发循环
		streamer, songName = p.nextSong(currentIndex, 1)
	}

	// 更新currentStream
	p.currentStream = streamer

	// 更新streamer
	p.streamer = streamer

	p.ctrl = &beep.Ctrl{Streamer: p.streamer}

	length := targetFormat.SampleRate.D(p.currentStream.Len()) / time.Second
	bar = getBar(int(length), songName)
	fmt.Printf("playing %v \n", songName)

	speaker.Play(p.ctrl)
}

// 重置播放器：清除当前的音频流并重置streamer。
func (p *Player) reset() *Player {
	speaker.Clear()
	p.streamer = nil
	p.ctrl = nil
	return p
}

// Open 开启播放器：open方法初始化音频输出设备，并开始播放音频。
func (p *Player) Open() *Player {
	speaker.Init(targetFormat.SampleRate, targetFormat.SampleRate.N(time.Second/10))
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

// PlayMp3 解码一个mp3文件，设置player当前的参数
func (p *Player) PlayMp3(file io.ReadCloser) beep.Streamer {
	streamer, _, err := mp3.Decode(file)
	if err != nil {
		log.Fatal("mp3 decode failed:", err)
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
	// 增加容错，两者不会严格相等
	return (float64(p.currentStream.Position()) / float64(p.currentStream.Len())) > 0.99
}

func (p *Player) changePlayLogic() {
	if p.playLogic == 2 {
		p.playLogic = 0
	} else {
		p.playLogic++
	}
}
