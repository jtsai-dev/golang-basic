package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(heightChecker([]int{1, 1, 4, 2, 1, 3}))

	// result := sortedSquares([]int{-4, -1, 0, 3, 10})
	// fmt.Println(result)
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

// 884
func uncommonFromSentences(A string, B string) []string {
	var result []string
	maps := make(map[string]int)
	for _, k := range append(strings.Split(A, " "), strings.Split(B, " ")...) {
		if num, ok := maps[k]; ok {
			maps[k] = num + 1
		} else {
			maps[k] = 1
		}
	}

	for k, v := range maps {
		if v == 1 {
			result = append(result, k)
		}
	}
	return result
}

// 242
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	ss := strings.Split(s, "")
	tt := strings.Split(t, "")
	sort.Strings(ss)
	sort.Strings(tt)
	for i, c := range ss {
		if c != tt[i] {
			return false
		}
	}
	return true
}

// 520
func detectCapitalUse(word string) bool {
	if strings.ToUpper(word) == word ||
		strings.ToLower(word) == word ||
		strings.Title(strings.ToLower(word)) == word {
		return true
	}

	return false
}

// 1051
func heightChecker(heights []int) int {
	source := make([]int, len(heights))
	copy(source, heights)
	sort.Ints(heights)
	count := 0
	for i, v := range source {
		if v != heights[i] {
			count++
		}
	}
	return count
}
