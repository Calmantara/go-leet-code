package main

import "strings"

func lengthOfLongestSubstring(s string) int {
	a := func(j *int, st string) {
		if len(st) > *j {
			*j = len(st)
		}
	}

	sub := ""
	i := 0
	for _, v := range s {
		if strings.Contains(sub, string(v)) {
			a(&i, sub)
			if strings.Index(sub, string(v)) != len(sub)-1 {
				sub = sub[strings.Index(sub, string(v))+1:]
				sub += string(v)
			} else {
				sub = string(v)
			}
			continue
		}
		sub += string(v)
	}
	a(&i, sub)
	return i
}
