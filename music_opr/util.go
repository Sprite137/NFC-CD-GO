package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Rsp struct {
	// Code int    `json:"code"` // 比如这里仅需要rsp字段（视具体情况而定），那Rsp结构体不解析code字段即可~
	Rsp string `json:"rsp"`
}

// GetReq 测试get请求
func GetReq() string {
	url := "http://127.0.0.1:8080/getReq"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("GetReq http.NewRequest err:", err)
		return ""
	}

	client := &http.Client{Timeout: 5 * time.Second} // 设置请求超时时长5s
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("GetReq http.DefaultClient.Do() err: ", err)
		return ""
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("GetReq ioutil.ReadAll() err: ", err)
		return ""
	}
	// fmt.Println("respBody: ", string(respBody))

	var rsp Rsp
	err = json.Unmarshal(respBody, &rsp)
	if err != nil {
		fmt.Println("GetReq json.Unmarshal() err: ", err)
		return ""
	}
	return rsp.Rsp

	//// 最后经过字段筛选后，再序列化成json格式即可（比如这里仅需要rsp字段（视具体情况而定），那Rsp结构体不解析code字段即可~）
	//result, err := json.Marshal(rsp)
	//if err != nil {
	//	fmt.Println("GetReq json.Marshal() err2: ", err)
	//	return ""
	//}

}
