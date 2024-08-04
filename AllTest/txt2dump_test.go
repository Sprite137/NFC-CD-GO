package AllTest

import (
	myutil "example.com/m/util"
	"testing"
)

func TestTxt2dump(t *testing.T) {
	dumpFilePath := "22.dump"   // M1卡的dump文件路径
	txtFilePath := "output.txt" // 输出的txt文件路径
	myutil.Dump2txt(txtFilePath, dumpFilePath)
}
