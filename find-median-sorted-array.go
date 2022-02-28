package main

import (
	"math"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) (result float64) {
	n := make(chan int)
	defer close(n)
	go func(n1, n2 []int) {
		n <- len(n1) + len(n2)
	}(nums1, nums2)

	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)

	k := <-n
	if k%2 == 1 {
		mid := float64(k) / 2.0
		result = float64(nums1[int(math.Ceil(mid)-1)])
	} else {
		mid := float64(k) / 2.0
		result = float64(
			(float64(nums1[int(math.Ceil(mid)-1)]) + float64(nums1[int(math.Ceil(mid))])) / 2.0)
	}

	return result
}
