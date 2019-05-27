package main

import "fmt"

func main() {
	/* 
	 * 数组是值类型
	 * [3]int 和 [4]int 是不同的数组类型
	 * slice 相当于数组的view 
	 * 添加元素如果超过cap，系统会重新分配更大的底层数组，
	   会增长1倍当前的长度，即 cap *= 2
	 */

	arr := create()
	fmt.Println(arr)

	arr1 := arr[2:]
	updateSlice(arr1)
	fmt.Println(arr, arr1)

	fmt.Println(sub(arr, 1, 2))

	arr = remove(arr, 2)
	fmt.Println(arr)

	fmt.Println(pop(arr))
	fmt.Println(arr)

	fmt.Println(tail(arr))
	fmt.Println(arr)

	arr2 := []int{ 0, 1, 2, 3, 4, 5, 6}
	arr3 := arr2[2:3]
	info(arr3)	// [2]
	arr4 := arr3[0:2]
	info(arr4)	// [2 3]：cap(arr3)=5，此时arr[0:2]未超过范围
	// fmt.Println(arr3[2])	// out of range: index > len(arr3)

	arr5 := arr2[2:6]
	arr6 := append(arr5, 4)// arr5对应的arr2[6]被append为4
	arr7 := append(arr6, 5)// arr7 不再是对arr2的view
	fmt.Println(arr2, arr5, arr6, arr7)
}

func create() []int {
	arr1 := []int{ 0, 1, 2, 3, 4, 5, 6, 7 }
	// arr2 := make([]int, 8)
	// arr3 := arr1[0: 2]

	return arr1
}

func updateSlice(slice []int){
	slice[0] = 100
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