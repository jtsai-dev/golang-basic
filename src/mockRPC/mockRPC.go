package main

import (
	"errors"
	"fmt"
	"time"
)

func client(ch chan string, request string) (string, error) {
	// 向 server 发送请求
	ch <- request

	// 对返回数据做处理
	select {
	case ack := <-ch:
		return ack, nil
	case <-time.After(time.Second):
		return "", errors.New("timeout")
	}
}

func server(ch chan string) {
	for {
		// 接收 client 请求的数据
		data := <-ch
		fmt.Println("server received:", data)

		// 模拟超时
		time.Sleep(time.Second * 2)

		// 回复确认信息
		ch <- "from server: copy that"
	}
}

func main() {
	ch := make(chan string)

	go server(ch)

	if rec, err := client(ch, "hello world"); err != nil {
		fmt.Printf("error: %s\n", err)
	} else {
		fmt.Println(rec)
	}
}
