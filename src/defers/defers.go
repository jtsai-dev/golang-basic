package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

/* defer
 *  方法结束前执行(即使panic也执行)
 *  先进后出
 */

func main() {
	writeFile("hello.txt")
	// defers()	 // 33 2 1 panic
	deferInFor() // 10 9 8 7 6 5 4 3 2 1 0 panic
}

func deferInFor() {
	for i := 0; i < 20; i++ {
		defer fmt.Println(i)
		if i == 10 {
			panic("largest then 10")
		}
	}
}

func defers() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	// panic("error")
	fmt.Println(4)
}

func writeFile(fileName string) {
	// file, err := os.Create(fileName)
	file, err := os.OpenFile(fileName, os.O_EXCL|os.O_CREATE, 0666)

	// 自定义的错误
	err = errors.New("custom err")

	if err != nil {
		// 对err的具体处理
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf(pathError.Op, pathError.Path, pathError.Err)
			return
		}
	}
	// 方法结束前关闭文件
	defer file.Close()

	writer := bufio.NewWriter(file)
	// 将数据 Flush 到文件中
	defer writer.Flush()

	for i := 0; i < 10; i++ {
		fmt.Fprintln(writer, i)
	}
}
