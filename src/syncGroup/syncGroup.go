package main

import (
	"fmt"
)

func main() {
	// var wg sync.WaitGroup
	n := 5
	// wg.Add(n) // 等待 n 个任务
	var channels [5]chan int
	for i := 0; i < n; i++ {
		channels[i] = make(chan int)
	}

	go func() {
		for i, c := range channels {
			c <- i
		}
	}()

	go func() {
		for _, c := range channels {
			// d := <-c              // 阻塞接收数据：执行时阻塞，直到接收到数据并赋值给
			if d, ok := <-c; ok { // 非阻塞接收数据
				fmt.Println(d)
			}
		}
	}()

	// wg.Wait()
}
