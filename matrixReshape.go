package leetcode

// https://leetcode-cn.com/problems/reshape-the-matrix/

func matrixReshape(mat [][]int, r int, c int) [][]int {
	var rows, cols int
	if rows = len(mat); rows == 0 {
		return mat
	}
	cols = len(mat[0])

	if r*c != rows*cols {
		return mat
	}

	result := make([][]int, r)
	for i := range result {
		result[i] = make([]int, c)
	}

	p := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result[p/c][p%c] = mat[i][j]
			p++
		}
	}

	return result
}
