package main

import (
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"os"
	"time"
)

func main() {
	// 打开音乐文件
	file, err := os.Open("/Users/xuzhi/Documents/work_project/NetCD-go/resources/霞据云佩.mp3")
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