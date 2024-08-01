package main

import (
	"github.com/k0kubun/go-ansi"
	"github.com/schollz/progressbar/v3"
	"time"
)

func main() {
	//file, err := os.Open("resources/霞据云佩.mp3")
	//if err != nil {
	//	log.Fatal("文件读取错误")
	//}
	//defer file.Close()
	//
	//player := muisc_opr.NewPlayer()
	//
	//player.PlayMp3(file)
	//
	//player.Open()
	//
	//select {}

	//speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	//
	//ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, streamer), Paused: false}
	//speaker.Play(ctrl)
	//
	//select {}

	//a := 1
	//add(&a)
	//fmt.Print(a)

	//opr := make(chan int)
	//go func() {
	//	for {
	//		//fmt.Println("请输入操作：")
	//		//fmt.Println("0. 暂停/播放")
	//		//fmt.Println("1. 下一首")
	//		//fmt.Println("2. 上一首")
	//		//fmt.Println("3. 退出")
	//		fmt.Println("输入")
	//		n, _ := fmt.Scanln()
	//		opr <- n
	//	}
	//}()
	//
	//for {
	//	switch <-opr {
	//	case 0:
	//		fmt.Printf("togglePlay, len(opr)：%v \n", len(opr))
	//	}
	//}
	//
	//select {}

	bar := progressbar.NewOptions(1000,
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(50),
		progressbar.OptionSetDescription("playing..."),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))
	for i := 0; i < 1000; i++ {
		bar.Add(1)
		time.Sleep(5 * time.Millisecond)
	}

}

func add(a *int) {
	addPlus(a)
}

func addPlus(a *int) {
	*a++
}
