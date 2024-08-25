package util

import (
	"bufio"
	_const "example.com/m/entity/const"
	"io"
	"log"
	"os"
	"strings"
)

func GetUid2SongListMap() map[string]string {
	// 读取txt文件
	Uid2SongListMap := make(map[string]string)

	fileHandle, err := os.OpenFile(_const.PLAYLISTPATH+"Uid2SongList.txt", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal("读取songList文件错误", err)
		return Uid2SongListMap
	}

	defer fileHandle.Close()

	reader := bufio.NewReader(fileHandle)

	// 按行处理txt
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		result := string(line)
		Uid2SongListMap[strings.Split(result, ":")[0]] = strings.Split(result, ":")[1] + ":" + strings.Split(result, ":")[2]

	}
	return Uid2SongListMap

}
