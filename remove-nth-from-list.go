package main

func removeNthFromEnd(head *ListNode, n int) (result *ListNode) {

	//define variable needed
	temp := struct {
		found   bool
		counter int
		bh      []int
	}{}

	for i := 0; i < 30; i++ {
		// check nil head
		if head == nil {
			break
		}
		if head.Val == n {
			temp.found = true
		}

		if temp.found {
			temp.counter++
		}

		if temp.counter == n {
			head = head.Next
			break
		} else {
			temp.bh = append(temp.bh, head.Val)
		}
		head = head.Next
	}

	// regenerate list
	result = head
	for i := len(temp.bh) - 1; i >= 0; i-- {
		t := &ListNode{
			Val:  temp.bh[i],
			Next: result,
		}
		result = t

	}

	return result
}
