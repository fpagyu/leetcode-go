package leetcode

import (
	"strings"
)

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	table := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	if len(digits) == 1 {
		return strings.Split(table[digits], "")
	}

	a := strings.Split(table[string(digits[0])], "")
	b := letterCombinations(digits[1:])

	r := make([]string, 0, len(a)*len(b))

	for i := range a {
		for j := range b {
			r = append(r, a[i]+b[j])
		}
	}

	return r
}
