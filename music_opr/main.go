package main

import (
	_const "example.com/m/entity/const"
	myUtil "example.com/m/util"
	"fmt"
	"github.com/clausecker/nfc/v2"
	"github.com/schollz/progressbar/v3"
	"log"
	"os"
	"strings"
	"time"
)

var isListening = true

var genAllSongTxt = true

var bar *progressbar.ProgressBar

var originSong = "jazz-logo.mp3"

var Modulations = []nfc.Modulation{
	{Type: nfc.ISO14443a, BaudRate: nfc.Nbr106},
	{Type: nfc.ISO14443b, BaudRate: nfc.Nbr106},
	{Type: nfc.Felica, BaudRate: nfc.Nbr212},
	{Type: nfc.Felica, BaudRate: nfc.Nbr424},
	{Type: nfc.Jewel, BaudRate: nfc.Nbr106},
	{Type: nfc.ISO14443biClass, BaudRate: nfc.Nbr106},
}

func main() {

	if genAllSongTxt {
		myUtil.GetAllSongList()
	}

	// 打开音乐文件
	file, err := os.Open(_const.SONGPATH + originSong)
	if err != nil {
		log.Fatal("读取file错误", err)
	}
	defer file.Close()

	// new一个播放器
	player := NewPlayer()

	// 设置播放的音乐
	player.PlayMp3(file)

	// 开始播放
	player.Open()
	fmt.Printf("playing %v \n", strings.Split(file.Name(), "/")[2])
	//player.togglePlay()

	// 打印歌曲进度，切换下一首
	currentIndex := 0

	// 获取当前歌曲在歌单的index
	for i, songName := range player.currentPlayList.SongNames {
		if songName == originSong {
			currentIndex = i
			break
		}
	}
	length := targetFormat.SampleRate.D(player.currentStream.Len()) / time.Second
	bar = myUtil.GetProgressBar(int(length), strings.Split(file.Name(), "/")[1])
	// 打印歌曲进度，播放切换下一首
	go func() {
		for {
			if player.ctrl.Paused != true {
				bar.Describe(fmt.Sprintf("playing    当前进度：%v ", player.currentPosition()))
				bar.Add(1)
				time.Sleep(1 * time.Second)
			}
			// 当前音乐播放完，切换下一首
			if player.isDone() {
				fmt.Println("正在切换下一首...")
				// 3表明是播放完切换下一首，便于LOOP时切换下一首
				player.changeSong(&currentIndex, 3)
			}
		}
	}()

	// 起协程获取opr
	opr := make(chan interface{})

	oprNFCChan := make(chan string)

	// 起协程监听网址的变化
	if isListening {
		go func() {
			// 打开NFC设备
			dev, err := nfc.Open("")
			if err != nil {
				log.Println("nfc设备打开失败")
			} else {
				fmt.Print("请将NFC卡靠近NFC读卡器...")
				// 循环等待NFC消息
				for {
					// Poll for 300ms
					tagCount, target, _ := dev.InitiatorPollTarget(Modulations, 1, 300*time.Millisecond)
					if err != nil {
						continue
					}

					// Check if a tag was detected
					if tagCount > 0 {
						Uid, err := myUtil.GetNfcUID(target)
						if err != nil {
							fmt.Printf("获取NFC卡Uid上失败%s \n", err)
						}
						oprNFCChan <- *Uid
						fmt.Printf("获取NFC卡Uid上%s \n ", *Uid)
					}

					// 休眠1s
					time.Sleep(1 * time.Second)

				}
			}
			defer dev.Close()
		}()
	}

	isInput := make(chan int, 1)

	go func() {
		oprNum := -1
		//fmt.Println("请输入操作：")
		//fmt.Println("0. 暂停/播放")
		//fmt.Println("1. 下一首")
		//fmt.Println("2. 上一首")
		//fmt.Println("3. 切换歌单")
		//fmt.Println("4. 切换播放逻辑 顺序-随机-循环")
		//fmt.Println("输入")
		for {
			if len(isInput) == 0 {
				fmt.Scanln(&oprNum)
				opr <- oprNum
				if oprNum == 3 {
					isInput <- 1
				}
			}

		}
	}()

	// 起协程来处理opr
	go func() {
		for {
			n := <-opr
			switch n {

			// 暂停，恢复
			case 0:
				player.playOrPause()

			// 切为下一首
			case 1:
				fmt.Printf("切换为下一首...\n")
				player.changeSong(&currentIndex, 0)

			// 切为上一首
			case 2:
				fmt.Printf("切换为上一首...\n")
				player.changeSong(&currentIndex, 1)

			// 切换歌单
			case 3:
				fmt.Print("请输入歌单txt: \n ")
				songListPath := ""
				fmt.Scanln(&songListPath)
				for {
					if songListPath == "" {
						songListPath = "自定义-许嵩x.txt"
					}
					if !player.currentPlayList.SetList(songListPath) {
						fmt.Printf("输入错误，请重新输入:")
						songListPath = ""
						fmt.Scanln(&songListPath)
						continue
					}
					if player.currentPlayList.SongNames != nil {
						currentIndex = -1
						player.changeSong(&currentIndex, 0)
						break
					}
				}
				fmt.Printf("已切换为歌单:%s \n", songListPath)
				<-isInput
			// 切换播放逻辑
			case 4:
				player.changePlayLogic()
			}

		}

	}()

	Uid2SOngListMap := myUtil.GetUid2SongListMap()
	go func() {
		for {
			oprWeb := <-oprNFCChan
			switch {
			case strings.Contains(Uid2SOngListMap[oprWeb], "更换专辑"):
				fmt.Printf("准备切换歌单为%v\n", strings.Split(Uid2SOngListMap[oprWeb], ":")[1])
				if !player.currentPlayList.SetList(strings.Split(Uid2SOngListMap[oprWeb], ":")[1]) {
					fmt.Printf("切换歌单失败 \n")
				} else {
					fmt.Printf("成功切换歌单为%v \n", strings.Split(Uid2SOngListMap[oprWeb], ":")[1])
				}
				if player.currentPlayList.SongNames != nil {
					currentIndex = -1
					player.changeSong(&currentIndex, 0)
				}
			case strings.Contains(Uid2SOngListMap[oprWeb], "暂停/播放"):
				player.playOrPause()
			}
		}
	}()

	select {}

}
