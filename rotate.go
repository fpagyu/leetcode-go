package leetcode

func rotate(matrix [][]int) {
	n := len(matrix)

	for j := 0; j < n/2; j++ {
		for i := j; i < n-1-j; i++ {
			x, y := i, j
			src := matrix[j][i]
			cnt := 4
			for cnt > 0 {
				cnt--
				x, y = (n - 1 - y), x
				src, matrix[y][x] = matrix[y][x], src
			}
		}
	}
}
