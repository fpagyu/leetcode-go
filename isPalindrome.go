package leetcode

func isPalindrome(s string) bool {
	p0 := 0
	p1 := len(s) - 1

	for p0 < p1 {
		a := toLower(s[p0])
		if a == 0 {
			p0++
			continue
		}

		b := toLower(s[p1])
		if b == 0 {
			p1--
			continue
		}

		if a != b {
			return false
		}

		p0++
		p1--
	}

	return true
}

func toLower(ch byte) byte {
	if ch >= 48 && ch <= 57 {
		return ch
	}

	if ch >= 65 && ch <= 90 {
		return ch
	}

	if ch >= 97 && ch <= 122 {
		return ch - 32
	}

	return 0
}
