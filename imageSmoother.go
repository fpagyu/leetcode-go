package leetcode

// https://leetcode-cn.com/problems/image-smoother/

func imageSmoother(img [][]int) [][]int {
	rows := len(img)
	cols := len(img[0])

	ans := make([][]int, rows)
	for i := range ans {
		ans[i] = make([]int, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			ans[i][j] += img[i][j]
			cnt := 1

			if col := j - 1; col >= 0 {
				ans[i][j] += img[i][col]
				cnt++
			}

			if col := j + 1; col < cols {
				ans[i][j] += img[i][col]
				cnt++
			}

			if row := i - 1; row >= 0 {
				ans[i][j] += img[row][j]
				cnt++

				if col := j - 1; col >= 0 {
					ans[i][j] += img[row][j-1]
					cnt++
				}

				if col := j + 1; col < cols {
					ans[i][j] += img[row][j+1]
					cnt++
				}
			}

			if row := i + 1; row < rows {
				ans[i][j] += img[row][j]
				cnt++

				if col := j - 1; col >= 0 {
					ans[i][j] += img[row][j-1]
					cnt++
				}

				if col := j + 1; col < cols {
					ans[i][j] += img[row][j+1]
					cnt++
				}
			}

			ans[i][j] /= cnt
		}
	}

	return ans
}
