package entity

// PlayList 歌单
type PlayList struct {
	Name      string   // 歌单名称
	SongsName []string // 歌曲列表
	index     int      // 当前歌曲在列表中的index
}

// NewPlayList 创建一个新歌单
func NewPlayList(name string, list []string) *PlayList {
	return &PlayList{
		Name:      name,
		SongsName: list,
		index:     0,
	}
}

// Next 下一首的index
func (p *PlayList) getNextSongIndex() int {
	p.index++
	if p.index >= len(p.SongsName) {
		p.index = 0
	}
	return p.index
}

// SetList 设置歌单的列表
func (p *PlayList) SetList(list []string) {
	p.SongsName = list
}
