package leetcode

// https://leetcode-cn.com/problems/first-unique-character-in-a-string/

// 方法1
func firstUniqChar(s string) int {
	h := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		h[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		if h[s[i]] == 1 {
			return i
		}
	}
	return -1
}

// 方法2
// func firstUniqChar(s string) int {
// 	h := make(map[byte]int)
// 	for i := 0; i < len(s); i++ {
// 		h[s[i]] = i
// 	}
// 	for i := 0; i < len(s); i++ {
// 		if h[s[i]] == i {
// 			return i
// 		} else {
// 			delete(h, s[i])
// 		}
// 	}
// 	return -1
// }
