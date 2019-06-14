package main

import (
	"unicode/utf8"
	"fmt"
)

func main() {
	str := "eng中文"
	
	for i, c := range str {
		fmt.Printf("(%d %X %c) ", i, c, c)
	}
	fmt.Println()

	fmt.Println("--str--")
	// len获取字节长度，中文字符占3个字节；utf8.RuneCountInString，获取字符长度
	fmt.Println("len: ", len(str), "RuneCount: ", utf8.RuneCountInString(str))

	fmt.Println("--str1--")
	str1 := []rune(str)
	for _, c := range str1 {
		fmt.Println(c)
	}

	bytes := []byte(str)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
}