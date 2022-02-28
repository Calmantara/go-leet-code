package main

func mergeTwoLists(list1 *ListNode, list2 *ListNode) (result *ListNode) {
	n := []int{}
	for i := 0; i < 50; i++ {
		if list1.Next == nil {
			n = append(n, list2.Val)
			list2 = list2.Next
			continue
		} else if list2.Next == nil {
			n = append(n, list1.Val)
			list1 = list1.Next
			continue
		} else if list2.Next == nil && list1.Next == nil {
			n = append(n, list1.Val, list2.Val)
			break
		}

		if list1.Val > list2.Val {
			n = append(n, list2.Val)
			list2 = list2.Next
		} else if list2.Val > list1.Val {
			n = append(n, list1.Val)
			list1 = list1.Next
		} else {
			n = append(n, list1.Val, list2.Val)
			list1 = list1.Next
			list2 = list2.Next
		}
	}

	result = &ListNode{
		Val: n[len(n)-1],
	}
	for i := len(n) - 2; i >= 0; i-- {
		a := &ListNode{
			Val:  n[i],
			Next: result,
		}

		result = a
	}

	return result
}
