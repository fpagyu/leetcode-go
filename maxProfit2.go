package leetcode

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/

func maxProfit2(prices []int) int {
	result := 0
	min := prices[0]

	for i := 1; i < len(prices); i++ {
		if v := prices[i] - min; v > 0 {
			result += v
		}
		min = prices[i]
	}

	return result
}
