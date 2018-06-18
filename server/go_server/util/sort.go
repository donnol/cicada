package util

// QuickSort 快速排序
func QuickSort(s []int) {
	l, r := 0, len(s)-1
	quickSort(s, l, r)
}

func quickSort(s []int, l, r int) {
	if l < r {
		p := partition(s, l, r)
		quickSort(s, l, p-1)
		quickSort(s, p+1, r)
	}
}

func partition(s []int, l, r int) int {
	pivot := s[r] // 选取r对应的元素作为基准
	i := l - 1
	for j := l; j <= r-1; j++ {
		if s[j] < pivot { // 改变符号，可以改变排序的方向
			i++
			s[j], s[i] = s[i], s[j]
		}
	}
	s[i+1], s[r] = s[r], s[i+1]
	return i + 1
}
