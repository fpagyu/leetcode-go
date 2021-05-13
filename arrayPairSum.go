package leetcode

import "sort"

// https://leetcode-cn.com/problems/array-partition-i/
// 重刷
// 理解题目

/**
 * 一个朴素（贪心）的理解：
 *
 * 假设排完序的结果为a1<=b1<=a2<=b2<=...<=an<=bn
 * 那么a1应该跟谁一组呢？
 *
 * a1作为全局最小值，无论跟谁一组a1都会被累加进答案，
 * 相反，a1的搭档会被永久排除。
 * 既然如此，莫不如排除一个较小的数，即给a1找一个“最小的搭档”b1。
 *
 * 当a1、b1被处理之后，a2同理分析。
 * 所以，最终选择a1,a2,...,an会得到最好的结果。
 */
func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	result := 0
	for i := 0; i < len(nums); i += 2 {
		result += nums[i]
	}
	return result
}
