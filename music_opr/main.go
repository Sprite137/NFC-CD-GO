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
			time.Sleep(1 * time.Second)
			fmt.Print(player.currentPosition(), "\n")
			// 当前音乐播放完，切换下一首
			if player.isDone() {
				fmt.Println("正在切换下一首...")
				player.changeSong(&currentIndex)
				currentIndex++
			}
		}
	}()

	select {}

}
