package main

import "strings"

func findSubstring(s string, words []string) (result []int) {
	for i, _ := range words {
		l := words[i:]
		l = append(l, words[:i]...)

		t := strings.Join(l, "")
		idx := strings.Index(s, t)
		if idx >= 0 {
			result = append(result, idx)
		}
	}
	return result
}
