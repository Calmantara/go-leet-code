package main

import "math"

func divideString(s string, k int, fill byte) (result []string) {
	l := int(math.Ceil(float64(len(s) / k)))
	r := len(s) % k
	rl := r > 0
	rs := make(chan string)
	defer close(rs)
	if rl {
		l++
		go func() {
			m := ""
			for i := r; i < k; i++ {
				m += string(fill)
			}
			rs <- m
		}()
	}

	for i := 0; i < l; i++ {
		if i == l-1 {
			h := s[i*k:]
			if rl {
				h += <-rs
			}
			result = append(result, h)
			continue
		}
		result = append(result, s[i*k:(k*i)+k])
	}
	return result
}
