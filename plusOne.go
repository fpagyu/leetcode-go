package leetcode

// https://leetcode-cn.com/problems/plus-one/

func plusOne(digits []int) []int {
	var flag bool
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
			flag = i == 0
		} else {
			digits[i] += 1
			break
		}
	}

	result := digits
	if flag {
		result = []int{1}
		result = append(result, digits...)
	}

	return result
}
