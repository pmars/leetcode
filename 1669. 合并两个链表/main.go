package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	header := list1

	for ; a > 1; a, b = a-1, b-1 {
		list1 = list1.Next
	}
	tmp := list1.Next
	list1.Next = list2
	for list2.Next != nil {
		list2 = list2.Next
	}
	for ; b > 0; b-- {
		tmp = tmp.Next
	}
	list2.Next = tmp

	return header
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
		fmt.Printf("%v ", tmp.Val)
		tmp = tmp.Next
	}
	fmt.Println()
}

func main() {
	{
		l1 := buildList([]int{0, 1, 2, 3, 4, 5})
		l2 := buildList([]int{1000000, 1000001, 1000002})
		outputList(l1)
		outputList(l2)
		res := mergeInBetween(l1, 3, 4, l2)
		outputList(res)
	}

	{
		l1 := buildList([]int{0, 1, 2, 3, 4, 5, 6})
		l2 := buildList([]int{1000000, 1000001, 1000002, 1000003, 1000004})
		outputList(l1)
		outputList(l2)
		res := mergeInBetween(l1, 2, 5, l2)
		outputList(res)
	}
}
