package main

import "fmt"

func main() {
	//file, err := os.Open("/Users/xuzhi/Documents/work_project/NetCD-go/resources/霞据云佩.mp3")
	//if err != nil {
	//	// 处理错误
	//}
	//defer file.Close()
	//
	//// 解码音乐文件
	//streamer, format, err := mp3.Decode(file)
	//if err != nil {
	//	// 处理错误
	//}
	//
	//defer streamer.Close()
	//speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	//
	//ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, streamer), Paused: false}
	//speaker.Play(ctrl)

	a := 1
	add(&a)
	fmt.Print(a)

}

func add(a *int) {
	addPlus(a)
}

func addPlus(a *int) {
	*a++
}
