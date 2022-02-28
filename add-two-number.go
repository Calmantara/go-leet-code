package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) (result *ListNode) {

	buf := 0
	var tr []*ListNode
	for i := 0; i < 100; i++ {

		sum := (l1.Val + l2.Val) + buf
		buf = sum / 10
		tr = append(tr, &ListNode{sum % 10, nil})

		// conditional for next linked
		if l2.Next == nil && l1.Next == nil {
			if buf > 0 {
				tr = append(tr, &ListNode{buf, nil})
			}
			break
		}

		if l1.Next != nil {
			l1 = l1.Next
		} else {
			l1 = &ListNode{0, nil}
		}

		if l2.Next != nil {
			l2 = l2.Next
		} else {
			l2 = &ListNode{0, nil}
		}
	}

	var j *ListNode
	result = tr[len(tr)-1]
	for i := len(tr) - 2; i >= 0; i-- {
		j = tr[i]
		j.Next = result
		result = j
	}

	return result
}
