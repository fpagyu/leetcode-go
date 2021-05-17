package leetcode

// https://leetcode-cn.com/problems/degree-of-an-array/
// todo: 重刷: 可看题解
func findShortestSubArray(nums []int) int {
	d := 1                     // 数组的度
	h := make(map[int]*[3]int) // [3]int数组分别存储元素第一次，最后一次出现的位置,以及出现频率

	for i, num := range nums {
		if v, has := h[num]; has {
			(*v)[1] = i
			(*v)[2] += 1
			if (*v)[2] > d {
				d = (*v)[2]
			}
		} else {
			h[num] = &([3]int{i, i, 1})
		}
	}

	min := -1
	for _, v := range h {
		if (*v)[2] == d {
			if tmp := (*v)[1] - (*v)[0] + 1; tmp < min || min == -1 {
				min = tmp
			}
		}
	}

	return min
}
