package AllTest

import (
	"fmt"
	"github.com/clausecker/nfc/v2"
	"testing"
)

func TestNFC(t *testing.T) {
	//fmt.Println(nfc.Version())
	//devices, err := nfc.ListDevices()
	//if err != nil {
	//	fmt.Print(err)
	//	return
	//}
	//for _, d := range devices {
	//	fmt.Printf("found device %s", d)
	//}
	//// 检测 NFC 标签
	//if nfcDev. {
	//	// 读取 NFC 标签
	//	tag, err := nfcDev.Read()
	//	if err != nil {
	//		fmt.Printf("读取 NFC 标签失败: %v\n", err)
	//		return
	//	}
	//
	//	// 显示标签信息
	//	fmt.Printf("检测到 NFC 标签: %+v\n", tag)
	//} else {
	//	fmt.Println("未检测到 NFC 标签")
	//}
	//fmt.Println(nfc.Version())
	devices, err := nfc.ListDevices()
	if err != nil {
		fmt.Print(err)
		return
	}
	for _, d := range devices {
		fmt.Printf("found device %s", d)
	}
}
