package slice

import (
	"strconv"
	"testing"
)

// 使用 Slice 来加速 String，提升 2 倍性能
// 原理在于切片 b，在每次循环中复用缓冲区，避免了性能问题
// 基本可以分析知道性能瓶颈主要会在开缓冲区上

// 来源： https://blog.thinkeridea.com/201901/go/slice_de_yi_xie_shi_yong_ji_qiao.html

/*
	goos: darwin
	goarch: arm64
	pkg: github.com/ozline/go-codes/benchmark/slice
	cpu: Apple M1 Pro
	BenchmarkSliceInt2String4-10    	 5500680	       215.4 ns/op	      72 B/op	       7 allocs/op
	BenchmarkSliceInt2String5-10    	11950638	        99.61 ns/op	      64 B/op	       4 allocs/op
*/

func BenchmarkSliceInt2String4(t *testing.B) {
	for i := 0; i < t.N; i++ {
		SliceInt2String4([][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		})
	}
}

func BenchmarkSliceInt2String5(t *testing.B) {
	for i := 0; i < t.N; i++ {
		SliceInt2String5([][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		})
	}
}

func SliceInt2String4(s [][]int) []string {
	res := make([]string, len(s))
	for i, v := range s {
		if len(v) < 1 {
			res[i] = ""
			continue
		}

		res[i] += strconv.Itoa(v[0])
		for j := 1; j < len(v); j++ {
			res[i] += "," + strconv.Itoa(v[j])
		}
	}

	return res
}

// 使用了 slice 来提速
func SliceInt2String5(s [][]int) []string {
	res := make([]string, len(s))
	b := make([]byte, 0, 256)
	for i, v := range s {
		if len(v) < 1 {
			res[i] = ""
			continue
		}

		b = b[:0]
		b = append(b, strconv.Itoa(v[0])...)
		for j := 1; j < len(v); j++ {
			b = append(b, ',')
			b = append(b, strconv.Itoa(v[j])...)
		}

		res[i] = string(b)
	}

	return res
}
