package main

import (
	"strconv"
	"strings"
)

func myAtoi(s string) int {
	s = strings.ReplaceAll(s, " ", "")
	sn := ""
	st := false
	for _, k := range s {
		if !st && (k >= 48 && k <= 57) {
			st = true
		}
		if st {
			if k >= 48 && k <= 57 {
				sn += string(k)
			} else {
				break
			}
		}
	}
	if len(sn) <= 0 {
		return 0
	}
	if strings.Index(s, sn) > 0 {
		if s[strings.Index(s, sn)-1] == 45 || s[strings.Index(s, sn)-1] == 46 {
			sn = string(s[strings.Index(s, sn)-1]) + sn
		}
	}

	h, _ := strconv.Atoi(sn)

	if h > 2147483647 {
		h = 2147483647
	} else if h < -2147483648 {
		h = -2147483648
	}
	return h
}
