package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func server(address string, exitChan chan int) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println(err.Error())
		exitChan <- 1
	}
	defer listener.Close()

	fmt.Println("listening address:", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		go func() {
			reader := bufio.NewReader(conn)
			for {
				str, err := reader.ReadString('\n')
				if err != nil {
					fmt.Println(err.Error())
					conn.Close()
					break
				} else {
					fmt.Println(str)
				}
			}
		}()
	}
}

func main() {
	// 创建一个程序结束码的通道
	exitChan := make(chan int)
	// 将服务器并发运行
	go server("127.0.0.1:7001", exitChan)
	// 通道阻塞, 等待接收返回值
	code := <-exitChan
	// 标记程序返回值并退出
	os.Exit(code)
}
