package main

import "fmt"

func main() {
	l1 := &ListNode{
		Val:  5,
		Next: nil,
	}
	l2 := &ListNode{
		Val:  4,
		Next: l1,
	}
	l3 := &ListNode{
		Val:  3,
		Next: l2,
	}
	l4 := &ListNode{
		Val:  2,
		Next: l3,
	}
	l5 := &ListNode{
		Val:  1,
		Next: l4,
	}

	r := removeNthFromEnd(l5, 1)
	for i := 0; i < 30; i++ {
		if r == nil {
			break
		}
		fmt.Println(r.Val)
		r = r.Next
	}

}
