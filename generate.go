package leetcode

// https://leetcode-cn.com/problems/pascals-triangle/
// 杨辉三角

func generate(numRows int) [][]int {
	if numRows <= 0 {
		return nil
	}

	result := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		for j := 1; j < len(row)-1; j++ {
			row[j] = result[i-1][j-1] + result[i-1][j]
		}
		row[0], row[len(row)-1] = 1, 1
		result[i] = row
	}
	return result
}
