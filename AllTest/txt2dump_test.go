package AllTest

import (
	myutil "example.com/m/util"
	"testing"
)

func TestDump2Txt(t *testing.T) {
	dumpFilePath := "output_dump.dump" // M1卡的dump文件路径
	txtFilePath := "output_txt.txt"    // 输出的txt文件路径
	myutil.Dump2txt(txtFilePath, dumpFilePath)
}

func TestTxt2Dump(t *testing.T) {
	dumpFilePath := "output_dump.dump" // M1卡的dump文件路径
	txtFilePath := "songs.txt"         // 输出的txt文件路径
	myutil.Txt2dump(txtFilePath, dumpFilePath)
}
