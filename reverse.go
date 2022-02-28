package main

import (
	"strconv"
	"strings"
)

func reverse(x int) int {
	s := strconv.Itoa(int(x))
	s1 := 2147483647

	var isNegative bool
	if x < 0 {
		isNegative = true
		s = s[1:]
	}

	ar := strings.Split(s, "")
	for i, j2 := 0, len(ar)-1; i < j2; i, j2 = i+1, j2-1 {
		ar[i], ar[j2] = ar[j2], ar[i]
	}

	r, _ := strconv.Atoi(strings.Join(ar, ""))
	if r > s1 {
		return 0
	}

	if isNegative {
		if r > (s1 + 1) {
			return 0
		}
		r *= -1
	}
	return r
}
