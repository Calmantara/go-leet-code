package main

func divide(dividend int, divisor int) (result int) {
	var n bool
	if divisor < 0 {
		n = !n
		divisor *= -1
	}

	if dividend < 0 {
		n = !n
		dividend *= -1
	}

	for i := divisor; i <= dividend; i += divisor {
		result++
	}

	if n {
		result *= -1
	}
	if result > 2147483647 {
		result = 2147483647
	} else if result < -2147483648 {
		result = -2147483648
	}
	return result
}
