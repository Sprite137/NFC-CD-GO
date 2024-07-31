package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main_() {
	// 打开音乐文件
	file, err := os.Open("resources/jazz-logo.mp3")
	if err != nil {
		log.Fatal("读取file错误")
	}
	defer file.Close()

	// new一个播放器
	player := NewPlayer()

	// 设置播放的音乐
	player.PlayMp3(file)

	// 开始播放
	player.Open()

	// 打印歌曲进度，切换下一首
	currentIndex := 0

	for i, songPath := range allSongList {
		if songPath == file.Name() {
			currentIndex = i
			break
		}
	}
	go func() {
		for {
			if player.ctrl.Paused != true {
				fmt.Print(player.currentPosition(), "\n")
				time.Sleep(1 * time.Second)
			}
			// 当前音乐播放完，切换下一首
			if player.isDone() {
				fmt.Println("正在切换下一首...")
				player.changeSong(&currentIndex, 0)
			}
		}
	}()

	opr := make(chan int)

	// 起协程获取opr
	go func() {
		oprNum := -1

		for {

			//fmt.Println("请输入操作：")
			//fmt.Println("0. 暂停/播放")
			//fmt.Println("1. 下一首")
			//fmt.Println("2. 上一首")
			//fmt.Println("3. 退出")
			//fmt.Println("输入")
			fmt.Scanln(&oprNum)
			opr <- oprNum
		}
	}()

	// 起协程来处理opr
	go func() {
		for {
			n := <-opr
			switch n {
			// 暂停，恢复
			case 0:
				player.togglePlay()
				if player.ctrl.Paused {
					fmt.Printf("paused...\n")
				} else {
					fmt.Printf("playing...\n")
				}
			// 切为下一首
			case 1:
				fmt.Printf("切换为下一首...\n")
				player.changeSong(&currentIndex, 0)
			}

		}

	}()

	select {}

}
