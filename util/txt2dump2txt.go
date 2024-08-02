package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"unicode/utf8"
)

const (
	sectors   = 16 // M1卡有16个扇区
	blocks    = 4  // 每个扇区有4个块
	blockSize = 16 // 每个块的大小为16字节
)

// 模拟M1卡的内存结构
type M1CardMemory struct {
	Sectors [sectors][blocks][blockSize]byte
}

func txt2dump() {
	txtFilePath := "resources/songList.txt" // 你的txt文件路径
	dumpFilePath := "m1_card.dump"          // 输出的M1卡dump文件路径

	// 读取txt文件内容
	txtContent, err := ioutil.ReadFile(txtFilePath)
	if err != nil {
		log.Fatalf("读取txt文件失败: %v", err)
	}

	// 将中文字符转换为字节序列
	byteData := []byte(string(txtContent))

	// 检查数据长度是否符合M1卡的内存大小
	if len(byteData) > sectors*blocks*blockSize {
		log.Fatalf("数据长度超出M1卡的内存大小")
	}

	// 创建M1卡内存结构实例
	memory := M1CardMemory{}

	// 填充数据到M1卡内存结构中
	idx := 0
	for i := range memory.Sectors {
		for j := range memory.Sectors[i] {
			for k := 0; k < blockSize && idx < len(byteData); k++ {
				memory.Sectors[i][j][k] = byteData[idx]
				idx++
			}
		}
	}

	// 序列化M1卡内存结构到二进制文件
	serializedData, err := serializeM1CardMemory(memory)
	if err != nil {
		log.Fatalf("序列化M1卡内存结构失败: %v", err)
	}

	// 写入到.dump文件
	if err := ioutil.WriteFile(dumpFilePath, serializedData, 0644); err != nil {
		log.Fatalf("写入dump文件失败: %v", err)
	}

	fmt.Printf("M1卡的dump文件已保存为: %s\n", dumpFilePath)
}

// serializeM1CardMemory 序列化M1卡内存结构为二进制数据
func serializeM1CardMemory(memory M1CardMemory) ([]byte, error) {
	var serializedData []byte
	for _, sector := range memory.Sectors {
		for _, block := range sector {
			serializedData = append(serializedData, block[:]...)
		}
	}
	return serializedData, nil
}

func dump2txt() {
	dumpFilePath := "m1_card.dump" // M1卡的dump文件路径
	txtFilePath := "output.txt"    // 输出的txt文件路径

	// 读取dump文件内容
	dumpData, err := ioutil.ReadFile(dumpFilePath)
	if err != nil {
		log.Fatalf("读取dump文件失败: %v", err)
	}

	// 将二进制数据解码成十六进制字符串
	hexStr := hex.EncodeToString(dumpData)

	// 将十六进制字符串转换回字节序列
	byteData, err := hex.DecodeString(hexStr)
	if err != nil {
		log.Fatalf("十六进制解码失败: %v", err)
	}

	// 构建最终的字符串，包括中文字符和换行符，排除无效的UTF-8字符
	var result strings.Builder
	for len(byteData) > 0 {
		r, size := utf8.DecodeRune(byteData)
		if isValidUTF8(r) {
			result.WriteRune(r)
		}
		byteData = byteData[size:]
	}

	// 清理字符串，移除末尾的非预期字符
	cleanResult := strings.Trim(result.String(), "\x00")

	// 写入到txt文件
	if err := ioutil.WriteFile(txtFilePath, []byte(cleanResult), 0644); err != nil {
		log.Fatalf("写入txt文件失败: %v", err)
	}

	fmt.Printf("转换完成，txt文件已保存为: %s\n", txtFilePath)

	stringArray := strings.Split(result.String(), "\n")
	fmt.Println(stringArray)

}

func isValidUTF8(r rune) bool {
	// 排除非预期的UTF-8字符，只保留有效的UTF-8字符
	return r != '\uFFFD' && r != '\u0000'
}

func main() {
	txt2dump()
	dump2txt()
}
