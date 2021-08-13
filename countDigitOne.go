package leetcode

func countDigitOne(n int) int {
	if n == 0 {
		return 0
	}

	if n < 10 {
		return 1
	}

	v := 10
	for n/v > 9 {
		v *= 10
	}

	l := n % v
	return (l + 2) + countDigitOne(l)
}
