package main

import (
	"strconv"
	"strings"
)

func removeKdigits(num string, k int) string {

	if len(num) == k {
		return "0"
	}

	min := []string{}
	for j, v := range num {

		limit := len(min)
		if len(min) > 0 {
			limit--
		}

		for i := limit; i >= 0; i-- {
			// exit if stack nil
			if len(min) <= 0 {
				break
			}
			if k > 0 && string(v) < min[i] {
				k--
				min = min[:len(min)-1]
			}
		}

		min = append(min, string(v))

		if k <= 0 {
			min = append(min, strings.Split(num[j+1:], "")...)
			break
		}
	}

	if len(min) > 0 {
		for i := k; i >= 1; i-- {
			if len(min)-1 < 0 {
				break
			}
			min = min[:len(min)-1]
		}
	}
	if len(min) <= 0 {
		return "0"
	}
	b, _ := strconv.Atoi(strings.Join(min, ""))
	return strconv.Itoa(b)
}
