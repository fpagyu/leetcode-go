package leetcode

// https://leetcode-cn.com/problems/ransom-note/

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	h := make(map[byte]int)
	for i := 0; i < len(magazine); i++ {
		h[magazine[i]]++
	}

	for i := 0; i < len(ransomNote); i++ {
		if h[ransomNote[i]] == 0 {
			return false
		}
		h[ransomNote[i]]--
	}

	return true
}
