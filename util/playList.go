package util

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func GetAllSongList() {
	separator := "/"
	if runtime.GOOS == "windows" {
		separator = "\\"
	}

	var songPath = filepath.Join("resources", "music")
	workingDir, _ := os.Getwd()
	var filePath = filepath.Join("resources", "playList", "allSongList.txt")
	var allSongList []string
	err := filepath.WalkDir(filepath.Join(workingDir, songPath), func(path string, d os.DirEntry, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}

		if !d.IsDir() && strings.Contains(path, ".mp3") {
			allSongList = append(allSongList, strings.Split(path, separator)[len(strings.Split(path, separator))-1])
			fmt.Println(strings.Split(path, separator)[len(strings.Split(path, separator))-1])
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
