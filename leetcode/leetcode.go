package main

import (
	"fmt"
)

func main() {

}

func testmergeKLists() {
	print(mergeKLists([]*ListNode{
		&ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
		&ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
		&ListNode{Val: 2, Next: &ListNode{Val: 6}},
	}))
	fmt.Println()

	print(mergeKLists([]*ListNode{
		&ListNode{Val: 2},
		nil,
		&ListNode{Val: -1},
	}))
	fmt.Println()

	print(mergeKLists([]*ListNode{
		nil,
	}))
	fmt.Println()

	print(mergeKLists([]*ListNode{
		&ListNode{Val: -2, Next: &ListNode{Val: -1, Next: &ListNode{Val: -1, Next: &ListNode{Val: -1}}}},
		nil,
	}))
	fmt.Println()

	print(mergeKLists([]*ListNode{
		nil,
		&ListNode{Val: 1},
	}))
}
func print(l *ListNode) {
	if l != nil {
		fmt.Println(l.Val)
		print(l.Next)
	}
}

// 23 --
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) < 1 {
		return nil
	}
	res := &ListNode{Val: ^int(^uint(0) >> 1)}
	cur := res
	t := 0
	for len(lists) > 0 {
		cur.Next = &ListNode{Val: ^int(^uint(0) >> 1)}
		cur = cur.Next
		t = 0

		for i := 0; i < len(lists); i++ {
			if lists[i] == nil {
				lists = append(lists[:i], lists[i+1:]...)
				continue
			}
			if lists[t].Val > lists[i].Val {
				t = i
			}
		}
		if t < len(lists) {
			if lists[t] != nil {
				cur.Val = lists[t].Val
				lists[t] = lists[t].Next
			}
			if lists[t] == nil {
				lists = append(lists[:t], lists[t+1:]...)
			}
		} else {
			res.Next = nil
		}
	}
	return res.Next
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
