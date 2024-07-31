package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	// 打开音乐文件
	file, err := os.Open("resources/sound-sculptors.mp3")
	if err != nil {
		log.Fatal("读取file错误")
	}
	defer file.Close()

	player := newPlayer()

	player.playMp3(file)
	if player.streamer == nil {
		log.Fatal("Failed to decode the audio file")
	}

	player.open() // 初始化音频输出设备

	currentIndex := 0
	go func() {
		for {
			if player.ctrl.Paused != true {
				time.Sleep(1 * time.Second)
				fmt.Print(player.currentPosition(), "\n")
			}
			// 当前音乐播放完，切换下一首
			if player.isDone() {
				fmt.Println("正在切换下一首...")
				player.changeSong(&currentIndex)
			}
		}
	}()

	opr := make(chan int)

	// 起协程获取opr
	go func() {
		for {
			//fmt.Println("请输入操作：")
			//fmt.Println("0. 暂停/播放")
			//fmt.Println("1. 下一首")
			//fmt.Println("2. 上一首")
			//fmt.Println("3. 退出")
			//fmt.Println("输入")
			n, _ := fmt.Scanln()
			opr <- n
		}
	}()

	go func() {
		for {
			n := <-opr
			switch n {
			case 0:
				player.togglePlay()
				if player.ctrl.Paused {
					fmt.Printf("paused...\n")
				} else {
					fmt.Printf("playing...\n")
				}
			}
		}

	}()

	select {}

}
