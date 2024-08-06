package AllTest

import (
	myUtil "example.com/m/util"
	"fmt"
	"github.com/clausecker/nfc/v2"
	"log"
	"testing"
	"time"
)

var modulations = []nfc.Modulation{
	{Type: nfc.ISO14443a, BaudRate: nfc.Nbr106},
	{Type: nfc.ISO14443b, BaudRate: nfc.Nbr106},
	{Type: nfc.Felica, BaudRate: nfc.Nbr212},
	{Type: nfc.Felica, BaudRate: nfc.Nbr424},
	{Type: nfc.Jewel, BaudRate: nfc.Nbr106},
	{Type: nfc.ISO14443biClass, BaudRate: nfc.Nbr106},
}

func TestNFC(t *testing.T) {

	//dev, err := nfc.Open("")
	//defer dev.Close()
	//if err != nil {
	//	t.Skip("Cannot open device:", err)
	//} else {
	//	t.Log("打开设备成功")
	//}
	//// 初始化为读卡器
	//err = dev.InitiatorInit()
	//if err != nil {
	//	return
	//}
	//print("开始读取NFC卡")
	//
	////读取NFC卡uid
	//buffer := make([]byte, 1024)
	//time.Sleep(time.Second * 1)
	//
	//// 读取Uid
	//tx := []byte{0x02, 0x00, 0x02, 0x35, 0x31, 0x03}
	//
	//// 读取数据
	////tx := []byte{0x02,0x00,0x04,0x35,0x33,0x00, 0x01,0x03}
	//
	//n, err := dev.InitiatorTransceiveBytes(tx, buffer, 10000)
	//
	//print(n)
	//if n > 0 {
	//	t.Log(buffer)
	//}

}

func TestGetNFCUid(t *testing.T) {
	// 打开NFC设备
	dev, err := nfc.Open("")
	if err != nil {
		log.Fatal("打开NFC设备失败", err)
	}
	defer dev.Close()

	// 循环等待NFC消息
	for {
		// Poll for 300ms
		tagCount, target, err := dev.InitiatorPollTarget(modulations, 1, 300*time.Millisecond)
		if err != nil {
			log.Println("Error polling the reader", err)
			continue
		}

		// Check if a tag was detected
		if tagCount > 0 {
			Uid, err := myUtil.GetNfcUID(target)
			if err != nil {
				fmt.Printf("获取NFC卡Uid上失败%s \n", err)
			}
			fmt.Printf("获取NFC卡Uid上%s \n ", *Uid)
		}

	}
}
