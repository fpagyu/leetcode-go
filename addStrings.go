package leetcode

// https://leetcode-cn.com/problems/add-strings/

func addStrings(num1 string, num2 string) string {
	length := len(num1)
	if len(num2) > length {
		length = len(num2)
	}

	ans := make([]byte, length+1)

	p1 := len(num1) - 1
	p2 := len(num2) - 1
	var a, b, c byte
	for i := len(ans) - 1; i != 0; i-- {
		if p1 >= 0 {
			a = num1[p1] - '0'
		} else {
			a = 0
		}

		if p2 >= 0 {
			b = num2[p2] - '0'
		} else {
			b = 0
		}

		s := (a + b + c)

		ans[i] = (s % 10) + '0'
		c = s / 10
		p1--
		p2--
	}

	if c != 0 {
		ans[0] = c + '0'
		return string(ans)
	}

	return string(ans[1:])
}
