package main

import (
	"fmt"
)

func selects() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	// // 使用遍历的方式接收数据
	// for {
	// 	if _, ok := <-ch1; ok {

	// 	}
	// 	if _, ok := <-ch2; ok {
	// 	}
	// }
	select {
	case d := <-ch1:
		fmt.Println(d)
	case d := <-ch2:
		fmt.Println(d)
	default:
		break
	}
}

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

// 带缓冲通道的 channel
//  当缓冲通道被填满时，尝试再次发送数据发生阻塞
func bufferChannel() {
	channel := make(chan int, 2) // 初始化一个数据类型为 int，带有2个元素缓冲区的通道
	go func() {
		for d := range channel {
			fmt.Printf("len: %d\n", len(channel))
			fmt.Println(d)
		}
	}()
	channel <- 1
	channel <- 2
	channel <- 3
	// 往 channel 中写入数据到缓冲通道，此时关闭通道（close(channel)），缓冲通道中
	// 的数据不会被释放，通道也未消失，仍可从改关闭的 channel 中获取数据
}

// 单向通道
func unilateralDemo() {
	// var sender chan<- int   // 只发送
	// var receiver <-chan int // 只接收
	// sender = make(chan<- int)
	// receiver = make(<-chan int)
}

func loopReceve() {
	c := make(chan int)
	go func() {
		for d := range c { // 循环接收数据
			fmt.Println(d)
		}
	}()

	for i := 0; i < 5; i++ {
		c <- i
	}
}

func channelDemo() {
	c := make(chan int)
	go func() {
		for {
			// d := <-c              // 阻塞接收数据：执行时阻塞，直到接收到数据并赋值给
			if d, ok := <-c; ok { // 非阻塞接收数据
				fmt.Println(d)
			}
		}
	}()

	for i := 0; i < 5; i++ {
		c <- i
	}
	// 最后的输入在 channelDemo 中可能未来得及输出，可尝试：
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
	// workerDemo()
	bufferChannel()
	// loopReceve()
	// channelDemo()
	// deadlock()
}
