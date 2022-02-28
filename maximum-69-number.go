package main

import (
	"strconv"
	"strings"
)

func maximum69Number(num int) int {
	s := strings.Split(strconv.Itoa(num), "")
	for i, _ := range s {
		if string(s[i]) == "6" {
			s[i] = "9"
			break
		}
	}
	i, _ := strconv.Atoi(strings.Join(s, ""))
	return i
}
