package main

func removeElement(nums []int, val int) (result int) {
	for _, v := range nums {
		if v != val {
			nums[result] = v
			result++
		}
	}
	return result
}
