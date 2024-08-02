package util

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetAllSongList() {
	const songPath = "resources/"
	err := filepath.WalkDir(songPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}

		if d.IsDir() {
			fmt.Printf("目录: %s\n", path)
		} else {
			fmt.Printf("文件: %s\n", path)
		}

		return nil
	})

	if err != nil {
		fmt.Println("遍历目录时发生错误:", err)
	}
}

func main() {
	GetAllSongList()
}
