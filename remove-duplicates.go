package main

func removeDuplicates(nums []int) (result int) {
	m := make(map[int]int)
	for _, v := range nums {
		if m[v] == 0 {
			nums[result] = v
			result++
		}
		m[v]++
	}
	return result
}
