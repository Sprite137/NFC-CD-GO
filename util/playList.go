package util

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetAllSongList() {
	const songPath = "resources/music/"
	// 定义文件路径
	const filePath = "resources/playList/" + "allSongList.txt"

	var allSongList []string
	err := filepath.WalkDir(songPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}

		if !d.IsDir() && strings.Split(path, ".")[1] == "mp3" {
			allSongList = append(allSongList, strings.Split(path, "\\")[2])
			fmt.Println(strings.Split(path, "\\")[2])
		}

		return nil
	})

	if err != nil {
		fmt.Println("genAllSongList err:", err)
	}

	// 写入txt
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open filePath fail: %v\n", err)
		return
	}
	defer file.Close() // 确保在函数结束时关闭文件

	// 写入切片中的每个元素到文件
	for _, song := range allSongList {
		_, err := file.WriteString(song + "\n")
		if err != nil {
			fmt.Printf("写入allSongList文件失败: %v\n", err)
			return
		}
	}

}