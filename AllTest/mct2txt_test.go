package AllTest

import (
	_ "io/ioutil"
	"testing"
	_ "unicode/utf16"
)

func TestMct2txt(t *testing.T) {
	mctFilePath := "/Users/xuzhi/Documents/work_project/NFC-CD-GO/AllTest/UID_13B6C75F_2024-08-04_13-21-32.mct" // MCT文件路径
	print(mctFilePath)
}
