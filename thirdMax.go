package leetcode

// https://leetcode-cn.com/problems/third-maximum-number/

// 使用大顶堆
// func thirdMax(nums []int) int {
// 	h := new(IntHeap)
// 	heap.Init(h)

// 	for _, n := range nums {
// 		heap.Push(h, n)
// 	}

// 	if h.Len() < 3 {
// 		return heap.Pop(h).(int)
// 	} else {
// 		heap.Pop(h)
// 		heap.Pop(h)
// 		return heap.Pop(h).(int)
// 	}
// }

// type IntHeap []int

// func (h IntHeap) Len() int {
// 	return len(h)
// }

// func (h IntHeap) Less(i, j int) bool {
// 	return h[i] > h[j]
// }

// func (h IntHeap) Swap(i, j int) {
// 	h[i], h[j] = h[j], h[i]
// }

// func (h *IntHeap) Push(x interface{}) {
// 	*h = append(*h, x.(int))
// }

// func (h *IntHeap) Pop() interface{} {
// 	old := *h
// 	n := len(old)
// 	x := old[n-1]
// 	*h = old[:n-1]
// 	return x
// }

func thirdMax(nums []int) int {
	a := -(1 << 32)
	b := -(1 << 32)
	c := -(1 << 32)

	for _, n := range nums {
		switch {
		case n == a:
		case n > a:
			a, b, c = n, a, b
		case n == b:
		case n > b:
			b, c = n, b
		case n > c:
			c = n
		}
	}

	if c == -(1 << 32) {
		return a
	}

	return c
}
