package leetcode

// https://leetcode-cn.com/problems/can-place-flowers/

func canPlaceFlowers(flowerbed []int, n int) bool {
	// if length := len(flowerbed) / 2; length > 0 &&  length < n {
	// 	return false
	// }

	for i := range flowerbed {
		if n == 0 {
			break
		}

		if flowerbed[i] == 0 {
			if (i == 0 || flowerbed[i-1] == 0) && (i+1 == len(flowerbed) || flowerbed[i+1] == 0) {
				flowerbed[i] = 1
				n--
			}
		}
	}

	return n == 0
}
