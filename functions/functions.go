package main

import "fmt"

type T struct {
	int
}

func (t T) test() {
	fmt.Println("类型 T 方法集包含全部 receiver T 方法。")
}
func (t *T) testP() {
	fmt.Println("类型 *T 方法集包含全部 receiver *T 方法。")
}

type S struct {
	T
}

func main() {
	t := T{}
	t.test()
	t.testP()
	(&t).testP()
	s := S{}
	s.test()
	s.testP()
	(&s).test()
	(&s).testP()
}
