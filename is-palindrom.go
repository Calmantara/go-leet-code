package main

import (
	"strconv"
	"strings"
)

func isPalindrom(x int) bool {
	j := strconv.Itoa(x)

	ar := strings.Split(j, "")
	for i, j2 := 0, len(ar)-1; i < j2; i, j2 = i+1, j2-1 {
		ar[i], ar[j2] = ar[j2], ar[i]
	}

	return j == strings.Join(ar, "")
}
