package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) (result [][]int) {
	l := len(nums)
	mp := make(map[string]bool)
	for i, _ := range nums {
		for k := i + 1; k < l; k++ {
			for j := k + 1; j < l; j++ {
				if (nums[i] + nums[k] + nums[j]) == 0 {
					it := []int{nums[i], nums[k], nums[j]}
					sort.Ints(it)
					s := fmt.Sprintf(`%v`, it)
					if !mp[s] {
						result = append(result, []int{nums[i], nums[k], nums[j]})
						mp[s] = true
					}
				}
			}
		}
	}
	return result
}
