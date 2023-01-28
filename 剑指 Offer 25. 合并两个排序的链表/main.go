package main

import "fmt"

// https://leetcode.cn/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof/submissions/

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归方法
func mergeTwoListsDigui(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}

// 循环方法
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	header := &ListNode{}
	tmp := header
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tmp.Next = l1
			l1 = l1.Next
		} else {
			tmp.Next = l2
			l2 = l2.Next
		}
		tmp = tmp.Next
	}
	for l1 != nil {
		tmp.Next = l1
		tmp = tmp.Next
		l1 = l1.Next
	}
	for l2 != nil {
		tmp.Next = l2
		tmp = tmp.Next
		l2 = l2.Next
	}

	return header.Next
}

func buildList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	header := &ListNode{}
	header.Val = nums[0]
	header.Next = buildList(nums[1:])

	return header
}

func outputList(header *ListNode) {
	tmp := header
	for tmp != nil {
		fmt.Printf("%v\t", tmp.Val)
		tmp = tmp.Next
	}
	fmt.Println()
}

func main() {
	l1 := buildList([]int{1, 2, 3, 4, 5, 6})
	l2 := buildList([]int{1, 3, 5, 7})
	outputList(l1)
	outputList(l2)
	res := mergeTwoLists(l1, l2)
	outputList(res)
}
