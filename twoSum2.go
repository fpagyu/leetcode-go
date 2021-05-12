package leetcode

// https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/

// 暴力破解
// func twoSum2(numbers []int, target int) []int {
// 	p := 0
// 	for i := p + 1; i < len(numbers); i++ {
// 		v := numbers[i] + numbers[p]
// 		switch {
// 		case v == target:
// 			return []int{p + 1, i + 1}
// 		case v > target:
// 			p++
// 			i = p
// 			if numbers[p] > target {
// 				return nil
// 			}
// 		case v < target:
// 			p++
// 			i = p
// 		}
// 	}
// 	return nil
// }

func twoSum2(numbers []int, target int) []int {
	p0 := 0
	p1 := 1

	for p0 < p1 && p1 < len(numbers) {
		v := numbers[p0] + numbers[p1]
		if v == target {
			return []int{p0 + 1, p1 + 1}
		}
		if v > target {
			p0++
			p1 = p0 + 1
			continue
		}
		if p1 == len(numbers)-1 {
			p0++
			p1 = p0 + 1
		} else {
			p1++
		}
	}

	return nil
}
