package entity

import (
	"bufio"
	_const "example.com/m/entity/const"
	"io"
	"os"
)

// PlayList 歌单
type PlayList struct {
	Name      string   // 歌单名称
	SongNames []string // 歌曲列表
	index     int      // 当前歌曲在列表中的index
}

// NewPlayList 创建一个新歌单
func NewPlayList(name string, list []string) *PlayList {
	return &PlayList{
		Name:      name,
		SongNames: list,
		index:     0,
	}
}

// Next 下一首的index
func (p *PlayList) getNextSongIndex() int {
	p.index++
	if p.index >= len(p.SongNames) {
		p.index = 0
	}
	return p.index
}

// SetList 设置歌单的列表
func (p *PlayList) SetList(filePath string) bool {
	fileHandle, err := os.OpenFile(_const.PLAYLISTPATH+filePath, os.O_RDONLY, 0666)
	if err != nil {
		//log.Fatal("读取songList文件错误", err)
		return false
	}

	defer fileHandle.Close()

	reader := bufio.NewReader(fileHandle)

	var results []string
	// 按行处理txt
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		results = append(results, string(line))
	}
	p.SongNames = results
	return true
}
