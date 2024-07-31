package AllTest

import (
	"fmt"
	"github.com/faiface/beep"
	_ "github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"os"
	"testing"
	"time"
)

// 播放时打印时间
func TestPrintTime(t *testing.T) {
	// 打开音乐文件
	file, err := os.Open("../resources/霞据云佩.mp3")
	if err != nil {
		// 处理错误
	}
	defer file.Close()

	// 解码音乐文件
	streamer, format, err := mp3.Decode(file)
	if err != nil {
		// 处理错误
	}

	defer streamer.Close()
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	for {
		select {
		case <-done:
			return
		case <-time.After(time.Second):
			speaker.Lock()
			fmt.Println(format.SampleRate.D(streamer.Position()).Round(time.Second))
			speaker.Unlock()
		}
	}

}

// 暂停和继续播放
func TestPauseAndResume(t *testing.T) {
	// 打开音乐文件
	file, err := os.Open("../resources/霞据云佩.mp3")
	if err != nil {
		// 处理错误
	}
	defer file.Close()

	// 解码音乐文件
	streamer, format, err := mp3.Decode(file)
	if err != nil {
		// 处理错误
	}

	defer streamer.Close()
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, streamer), Paused: false}
	speaker.Play(ctrl)

	for {
		//fmt.Print("Press [ENTER] to pause/resume. ")
		n, err := fmt.Scanln()
		if err != nil {
			// 处理错误

		}
		if n == 0 {
			speaker.Lock()
			ctrl.Paused = !ctrl.Paused
			speaker.Unlock()
		}

	}

}

// 改变音量
func TestChangeVolume(t *testing.T) {
	// 打开音乐文件
	file, err := os.Open("../resources/霞据云佩.mp3")
	if err != nil {
		// 处理错误
	}
	defer file.Close()

	// 解码音乐文件
	streamer, format, err := mp3.Decode(file)
	if err != nil {
		// 处理错误
	}

	defer streamer.Close()
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, streamer), Paused: false}
	volume := &effects.Volume{
		Streamer: ctrl,
		Base:     2,
		Volume:   0,
		Silent:   false,
	}
	speaker.Play(volume)

	for {
		//fmt.Print("Press [ENTER] to pause/resume. ")
		n, err := fmt.Scanln()
		if err != nil {
			// 处理错误

		}
		if n == 103 {
			speaker.Lock()
			volume.Volume += 0.1
			speaker.Unlock()
		}
	}
}
