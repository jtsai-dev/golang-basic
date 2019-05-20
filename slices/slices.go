package main

import "fmt"

func main() {
	/* MARK
	 * 数组是值类型
	 * [3]int 和 [4]int 是不同的类型
	 * slice 相当于数组的view 
	 * cap 在被使用完后，会增长1倍当前的长度，即 cap *= 2
	 */

	arr := make([]int, 0)
	for i := 0; i < 10; i++ {
		arr = append(arr, i)
	}

	fmt.Println(arr)

	fmt.Println(sub(arr, 1, 2))

	arr = remove(arr, 2)
	fmt.Println(arr)

	fmt.Println(pop(arr))
	fmt.Println(arr)

	fmt.Println(tail(arr))
	fmt.Println(arr)

	arr1 := make([]int, 0)
	for i := 0; i < 10; i++{
		arr1 = append(arr1, i)
		info(arr1)
	}

	arr2 := []int{ 0, 1, 2, 3, 4, 5, 6}
	arr3 := arr2[2:3]
	info(arr3)	// [2]
	arr4 := arr3[0:2]
	info(arr4)	// [2 3]：cap(arr3)=5，此时arr[0:2]未超过范围
	fmt.Println(arr3[2])	// out of range: index > len(arr3)
}

func init() {
	// arr1 := [...]int{ 0, 1, 2, 3, 4, 5, 6, 7 }
	// arr2 := make([]int, 8)
	// arr3 := arr1[0: 2]
}

func find(arr []int, target int) int{
	index := -1
	for i,v := range arr {
		if v == target{ 
			index = i
			break
		}
	}
	return index
}

func info(arr []int){
	fmt.Printf("len: %d, cap: %d ", len(arr), cap(arr))
	fmt.Println(arr)
}

func sub(arr []int, start int, length int) []int{
	end := start + length
	return arr[start:end]
}

func remove(arr []int, index int) []int {
	arr = append(arr[:index], arr[index + 1:]...)
	return arr
}

func pop(arr []int) int {
	front := arr[0]
	arr = arr[1:]
	return front
}

func tail(arr []int) int {
	tail := arr[len(arr) - 1]
	arr = arr[:len(arr) - 1]
	return tail
}