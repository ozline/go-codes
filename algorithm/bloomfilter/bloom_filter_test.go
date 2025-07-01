package bloomfilter

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestBloomFilterBasic(t *testing.T) {
	bf := NewWithFalsePositiveRate(1000, 0.01) // 1%误判率
	// 测试添加和查询
	testItems := [][]byte{
		[]byte("1"),
		[]byte("2"),
		[]byte("3"),
		[]byte("4"),
	}
	for _, item := range testItems {
		bf.Add(item)
		if !bf.Contains(item) {
			t.Errorf("Expected item %q to be in bloom filter", item)
		}
	}
	// 测试不存在的项
	nonExistentItems := [][]byte{
		[]byte("5"),
		[]byte("6"),
		[]byte("7"),
	}
	for _, item := range nonExistentItems {
		if bf.Contains(item) {
			t.Errorf("Item %q should not be in bloom filter", item)
		}
	}
}

func TestFalsePositiveRate(t *testing.T) {
	expectedItems := uint64(1000000)
	falsePositiveRate := 0.01 // 1%
	bf := NewWithFalsePositiveRate(expectedItems, falsePositiveRate)

	// 添加10000个随机字符串
	addedItems := make([][]byte, expectedItems)
	for i := uint64(0); i < expectedItems; i++ {
		item := randomString(30)
		addedItems[i] = item
		bf.Add(item)
	}

	// 检查所有添加的项是否都在过滤器中
	for _, item := range addedItems {
		if !bf.Contains(item) {
			t.Errorf("Added item not found in bloom filter")
		}
	}
	start := time.Now()
	// 检查1000000个未添加的随机字符串的误判率
	falsePositives := 0
	totalTests := 1000000
	for i := 0; i < totalTests; i++ {
		item := randomString(30)
		if bf.Contains(item) {
			falsePositives++
		}
	}
	consumingTime := time.Since(start)
	actualRate := float64(falsePositives) / float64(totalTests)
	t.Logf("Expected false positive rate: %.2f%%", falsePositiveRate*100)
	t.Logf("Actual false positive rate: %.2f%%, consuming time: %v", actualRate*100, consumingTime)
	if actualRate > falsePositiveRate*2 { // 允许实际误判率是预期的2倍以内
		t.Errorf("False positive rate too high: expected %.2f%%, got %.2f%%",
			falsePositiveRate*100, actualRate*100)
	}
}

func TestConcurrency(t *testing.T) {
	bf := NewWithFalsePositiveRate(10000, 0.01)
	var wg sync.WaitGroup

	// 并发添加
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			item := []byte(randomString(20))
			bf.Add(item)
			if !bf.Contains(item) {
				t.Errorf("Item not found after concurrent add")
			}
		}(i)
	}

	wg.Wait()
}

func randomString(length int) []byte {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return b
}
