package leetcode

// https://leetcode-cn.com/problems/pascals-triangle-ii/

func getRow(rowIndex int) []int {
	result := make([]int, rowIndex+1)

	result[0] = 1
	for i := 1; i <= rowIndex; i++ {
		for j := i - 1; j != 0; j-- {
			result[j] += result[j-1]
		}
	}

	return result
}
