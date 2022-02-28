package main

func myPow(x float64, n int) (result float64) {
	// checking x
	if x == float64(1.0) {
		return 1.0
	} else if x == float64(0.0) {
		return 0.0
	}

	//checking n
	neg := n < 0
	if neg {
		n *= -1
	} else if n == 0 {
		return 1.0
	}

	// mod
	ls := n % 2
	lm := n / 2

	result = x
	for i := 1; i < lm; i++ {
		result *= x
	}
	if lm > 0 {
		result *= result
	}
	if ls > 0 && lm > 0 {
		result *= x
	}
	if neg {
		result = 1 / result
	}
	return result
}
