package main

import (
	"unicode/utf8"
	"fmt"
)

func main() {
	str := "eng中文"
	
	for _, c := range str {
		fmt.Printf("%X ", c)
	}

	fmt.Println("--str--")
	// 9：len获取字节长度，中文字符占3个字节
	fmt.Println(len(str))
	// utf8.RuneCountInString：获取字符长度
	fmt.Println(utf8.RuneCountInString(str))

	fmt.Println("--str1--")
	str1 := []rune(str)
	for _, c := range str1 {
		fmt.Println(c)
	}
	fmt.Println(len(str1))
}