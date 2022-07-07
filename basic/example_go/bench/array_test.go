package bench

import (
	"math/rand"
	"testing"
)

const (
	size = 10000
)

func TestT(t *testing.T) {
	a := []int{1, 2, 3, 4}
	if len(a) != 4 {
		t.Fail()
	}
}

func preAllocByVar(src []int) {
	dst := make([]int, len(src))
	for i, v := range src {
		dst[i] = v
	}
}

func preAllocByRef(src []int) {
	dst := make([]int, len(src))
	for i := range src {
		dst[i] = src[i]
	}
}

func preAllocAppendByVar(src []int) {
	dst := make([]int, 0, len(src))
	for _, v := range src {
		dst = append(dst, v)
	}
}

func preAllocAppendByRef(src []int) {
	dst := make([]int, 0, len(src))
	for i := range src {
		dst = append(dst, src[i])
	}
}

func appendByVar(src []int) {
	var dst []int
	for _, v := range src {
		dst = append(dst, v)
	}
}

func appendByRef(src []int) {
	var dst []int
	for i := range src {
		dst = append(dst, src[i])
	}
}

func generateArray(size int) []int {
	x := make([]int, size)
	for i := range x {
		x[i] = rand.Int()
	}
	return x
}

func BenchmarkPreAllocByVar(b *testing.B) {
	a := generateArray(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		preAllocByVar(a)
	}
}

func BenchmarkPreAllocByRef(b *testing.B) {
	a := generateArray(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		preAllocByRef(a)
	}
}

func BenchmarkPreAllocAppendByVar(b *testing.B) {
	a := generateArray(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		preAllocAppendByVar(a)
	}
}

func BenchmarkPreAllocAppendByRef(b *testing.B) {
	a := generateArray(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		preAllocAppendByRef(a)
	}
}

func BenchmarkAppendByVar(b *testing.B) {
	a := generateArray(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		appendByVar(a)
	}
}

func BenchmarkAppendByRef(b *testing.B) {
	a := generateArray(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		appendByRef(a)
	}
}
