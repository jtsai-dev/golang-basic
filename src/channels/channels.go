package main

import (
	"fmt"
)

func worker(id int, c chan int) {
	for {
		v := <-c
		fmt.Printf("worker:%d, value:%c\n", id, v)
	}
}
func workerDemo() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go worker(i, channels[i])
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
}

func channelDemo() {
	c := make(chan int)
	go func() {
		for {
			fmt.Println(<-c)
		}
	}()

	c <- 1
	c <- 2
	c <- 3
	c <- 4
	c <- 5
	// 最后的输入在 channelDemo 中未来得及输出，可尝试：
	// time.Sleep(time.Millisecond)

	// output:
	// 1
	// 2
	// 3
	// 4
}

func deadlock() {
	// it will cause deadlock when input value into chan but no handling
	c := make(chan int)
	c <- 1
	c <- 2
	fmt.Println(<-c)
}

func main() {
	workerDemo()
	// channelDemo()
	// deadlock()
}
