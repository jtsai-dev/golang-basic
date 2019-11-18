package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
	"strconv"
)

func reflects() {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	v := reflect.ValueOf(x)
	fmt.Println("value:", v)
	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind())
	fmt.Println("value:", v.Float())
	fmt.Printf("value is(v.Interface()) %5.2e\n", v.Interface())
	y := v.Interface().(float64)
	fmt.Println(y)
}

func main() {
	a, b := 1, 2
	fmt.Println(a, b)
	a, b = swap(a, b)
	fmt.Println("after swap:", a, b)
	fmt.Println(eval(a, b, "+"))
	fmt.Println(eval(a, b, "%"))

	fmt.Println("pow(3, 4) is:", apply(
		func(a int, b int) int {
			return int(math.Pow(
				float64(a), float64(b)))
		}, 3, 4))

	fmt.Println("convertToBin", convertToBin(2))
}

func sum(a, b int) int {
	return a + b
}

// 返回多个值
func div(a, b int) (int, int) {
	return a / b, a % b
}
func div1(a, b int) (q, r int) {
	return a / b, a % b
}

func swap(a, b int) (int, int) {
	return b, a
}

func eval(a, b int, ops string) (int, error) {
	switch ops {
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		// panic("unsupported operation: " + ops)		// 对外暴露异常并中断程序执行
		return 0, fmt.Errorf("unsupported operation: %s", ops)
	}
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name() // 获取传入的方法名
	fmt.Printf("Calling function: %s with args "+
		"(%d, %d)\n", opName, a, b)

	return op(a, b)
}

func loopForever() {
	for {
		fmt.Println("loop forever")
	}
}

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}
