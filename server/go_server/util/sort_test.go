package util

import (
	"math/rand"
	"testing"
	"time"
)

func TestQuickSort(t *testing.T) {
	n := 1000000
	s := makeSlice(n)
	QuickSort(s)
	assertSort(t, s)
}

func assertSort(t *testing.T, s []int) {
	l := len(s)
	for i, v := range s {
		if i < l-1 && v > s[i+1] {
			t.Fatalf("bad sort, %v", s)
		}
	}
}

func BenchmarkQuickSort(b *testing.B) {
	b.StopTimer()

	n := 10000
	s := makeSlice(n)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var s1 = make([]int, n)
		copy(s1, s)
		QuickSort(s1)
	}
}

func TestBubbleSort(t *testing.T) {
	n := 10
	s := makeSlice(n)
	BubbleSort(s)
	assertSort(t, s)
}

func BenchmarkBubbleSort(b *testing.B) {
	b.StopTimer()

	n := 10000
	s := makeSlice(n)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var s1 = make([]int, n)
		copy(s1, s)
		BubbleSort(s1)
	}
}

func makeSlice(n int) (s []int) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		v := rand.Intn(n) + 1
		s = append(s, v)
	}

	return s
}
