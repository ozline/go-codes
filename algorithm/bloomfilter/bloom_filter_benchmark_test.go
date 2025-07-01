package bloomfilter

import (
	"testing"
)

func BenchmarkAdd(b *testing.B) {
	bf := NewWithFalsePositiveRate(uint64(b.N), 0.01)
	items := make([][]byte, b.N)
	for i := 0; i < b.N; i++ {
		items[i] = randomString(20)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bf.Add(items[i])
	}
}

func BenchmarkContains(b *testing.B) {
	bf := NewWithFalsePositiveRate(uint64(b.N), 0.01)
	items := make([][]byte, b.N)
	for i := 0; i < b.N; i++ {
		items[i] = randomString(20)
		bf.Add(items[i])
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = bf.Contains(items[i])
	}
}

func BenchmarkAddParallel(b *testing.B) {
	bf := NewWithFalsePositiveRate(uint64(b.N), 0.01)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			item := randomString(20)
			bf.Add(item)
		}
	})
}

func BenchmarkContainsParallel(b *testing.B) {
	bf := NewWithFalsePositiveRate(uint64(b.N), 0.01)
	// 预先生成测试数据并添加到布隆过滤器
	for i := 0; i < b.N; i++ {
		bf.Add(randomString(20))
	}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			bf.Contains(randomString(20))
		}
	})
}
