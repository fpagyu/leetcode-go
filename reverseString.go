package leetcode

// https://leetcode-cn.com/problems/reverse-string/

func reverseString(s []byte) {
	p0 := 0
	p1 := len(s) - 1

	for p0 < p1 {
		s[p0], s[p1] = s[p1], s[p0]
		p0++
		p1--
	}
}
