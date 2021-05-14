package leetcode

// https://leetcode-cn.com/problems/roman-to-integer/

// func romanToInt(s string) int {
// 	var prev byte
// 	ans := 0
// 	for i := len(s) - 1; i >= 0; i-- {
// 		switch s[i] {
// 		case 'I':
// 			if prev == 'V' || prev == 'X' {
// 				ans -= 1
// 			} else {
// 				ans += 1
// 			}
// 		case 'V':
// 			ans += 5
// 		case 'X':
// 			if prev == 'L' || prev == 'C' {
// 				ans -= 10
// 			} else {
// 				ans += 10
// 			}
// 		case 'L':
// 			ans += 50
// 		case 'C':
// 			if prev == 'D' || prev == 'M' {
// 				ans -= 100
// 			} else {
// 				ans += 100
// 			}
// 		case 'D':
// 			ans += 500
// 		case 'M':
// 			ans += 1000
// 		}
// 		prev = s[i]
// 	}
// 	return ans
// }

func romanToInt(s string) int {

	ans := 0
	prev := 0
	kv := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	for i := len(s) - 1; i >= 0; i-- {
		v := kv[s[i]]
		if v < prev {
			ans -= v
		} else {
			ans += v
		}
		prev = v
	}
	return ans
}
