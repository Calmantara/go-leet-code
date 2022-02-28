package main

import "sort"

func sortList(head *ListNode) (result *ListNode) {
	var ta []int
	for i := 0; i < 50000; i++ {
		ta = append(ta, head.Val)
		if head.Next == nil {
			break
		}
		head = head.Next
	}
	head = nil
	sort.Ints(ta)

	(result) = &ListNode{
		Val: ta[len(ta)-1],
	}
	for i := len(ta) - 2; i >= 0; i-- {
		t := &ListNode{
			Val:  ta[i],
			Next: result,
		}
		result = t
	}
	return result
}
