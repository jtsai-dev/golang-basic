package main

import "fmt"

func enums() {
	const (
		golang = iota	// 生成枚举值，默认为0
		_				// 跳过对应枚举
		java
		csharp
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(golang, java)
	fmt.Println(b, kb, mb)
}

/* golang 中没有枚举类型，使用 const 配合 iota 模拟枚举
 */
func main() {
	enums()
}