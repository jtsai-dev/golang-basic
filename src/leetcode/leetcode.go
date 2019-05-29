package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	result := sortedSquares([]int{-4, -1, 0, 3, 10})

	fmt.Println(result)
}

// 771
func numJewelsInStones(J string, S string) int {
	count := 0
	for _, v := range S {
		if strings.Index(J, string(v)) > -1 {
			count++
		}
	}

	return count
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 938
func rangeSumBST(root *TreeNode, L int, R int) int {
	sum := 0
	if root == nil {
		return sum
	}

	if root.Val >= L && root.Val <= R {
		sum += root.Val
	}
	sum += rangeSumBST(root.Left, L, R)
	sum += rangeSumBST(root.Right, L, R)

	return sum
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 237
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

// 709
func toLowerCase(str string) string {
	result := ""
	for _, v := range str {
		if v >= 65 && v <= 90 {
			v = v + 32
		}
		result = result + string(v)
	}
	return result
}

// 832
func flipAndInvertImage(A [][]int) [][]int {
	var cols = len(A[0])
	var rows = len(A)
	var result [][]int
	for i := 0; i < rows; i++ {
		row := make([]int, cols)
		result = append(result, row)
	}
	fmt.Println(result)
	for i, r := range A {
		for j, c := range r {
			result[i][cols-j-1] = (1 - c)
		}
	}
	return result
}

// 977
func sortedSquares(A []int) []int {
	var result []int
	for _, v := range A {
		result = append(result, v*v)
	}
	sort.Ints(result)
	return result
}
