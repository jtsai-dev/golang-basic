package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0))
}

func searchInsert(nums []int, target int) int {
	begin, end := 0, len(nums)-1
	for begin <= end {
		// mid := begin + (end-begin)/2
		mid := (begin + end) >> 1
		if target == nums[mid] {
			return mid
		}
		if target < nums[mid] {
			end = mid - 1
		} else {
			begin = mid + 1
		}
	}
	return begin
}

func longestValidParentheses1() {
	fmt.Println(longestValidParentheses("(()"))            // () 2
	fmt.Println(longestValidParentheses(")()())"))         // ()() 4
	fmt.Println(longestValidParentheses(")(())))(())())")) // (())() 6
	fmt.Println(longestValidParentheses(")()())"))         // ()() 4
	fmt.Println(longestValidParentheses("()(())"))         // ()(()) 6
}

// 32--
func longestValidParentheses(s string) int {
	s = strings.TrimRight(strings.TrimLeft(s, ")"), "(")
	dp := make([]int, len(s))
	res := 0
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' && i > 1 {
				dp[i] = dp[i-2] + 2
			} else {
				dp[i] = 2
			}
			res = max(res, dp[i])
		} else {

		}
	}
	return res

	// maxans := 0
	// s = strings.TrimRight(strings.TrimLeft(s, ")"), "(")
	// dp := make([]int, len(s))
	// for i := 1; i < len(s); i++ {
	// 	if s[i] == ')' {
	// 		if s[i-1] == '(' {
	// 			if i >= 2 {
	// 				dp[i] = dp[i-2] + 2
	// 			} else {
	// 				dp[i] = 2
	// 			}
	// 		} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
	// 			dp[i] = dp[i-1] + 2
	// 			if (i - dp[i-1]) >= 2 {
	// 				dp[i] = dp[i] + dp[i-dp[i-1]-2]
	// 			}
	// 		}
	// 		if dp[i] > maxans {
	// 			maxans = dp[i]
	// 		}
	// 	}
	// }
	// return maxans

	// 非连续
	// s = strings.TrimLeft(s, ")")
	// s = strings.TrimRight(s, "(")
	// left, right := 0, 0
	// for i := 0; i < len(s); i++ {
	// 	if s[i] == '(' {
	// 		left++
	// 	} else {
	// 		right++
	// 	}
	// }
	// if left < right {
	// 	return left * 2
	// }

	// return right * 2
}

const (
	INT_MAX = int(^uint(0) >> 1)
	INT_MIN = ^INT_MAX
)

func absDiff(a, b int) int {
	diff := a - b
	if diff < 0 {
		diff *= -1
	}
	return diff
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}
