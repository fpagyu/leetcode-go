package leetcode

// https://leetcode-cn.com/problems/reverse-vowels-of-a-string/

func reverseVowels(s string) string {
	arr := make([]byte, len(s))

	p0 := 0
	p1 := len(s) - 1

	h := map[byte]struct{}{
		'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {},
		'A': {}, 'E': {}, 'I': {}, 'O': {}, 'U': {},
	}
	for p0 <= p1 {
		if _, has := h[s[p0]]; !has {
			arr[p0] = s[p0]
			p0++
			continue
		}

		if _, has := h[s[p1]]; !has {
			arr[p1] = s[p1]
			p1--
			continue
		}

		arr[p0], arr[p1] = s[p1], s[p0]
		p0++
		p1--
	}

	return string(arr)
}
