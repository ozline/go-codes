package code

import (
	"math/rand"
	"testing"
	"time"
)

// 测试 Select 相关的代码，说起来一直到准备以 go 为主力语言应聘的时候，我才知道还有 select 这个东西，实在惭愧
// 在 macOS 下，偏离度大约在 5% 左右，但看到网上的资料基本都是 0.5% 左右，怎么差了一个数量级？

// 不过复习一下，在一个 channel 的时候退化为阻塞操作，多个 channel 则实现了随机轮询，近似 5% 左右（起码在我机器上是 5%）

func TestSelect(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// 模拟一个函数持续向两个通道发送数据
	go func() {
		for {
			ch1 <- rand.Intn(100)
		}
	}()
	go func() {
		for {
			ch2 <- rand.Intn(100)
		}
	}()

	count1, count2 := 0, 0
	total := 10000000 // 100 万次测试
	rand.New(rand.NewSource(time.Now().Unix()))

	for range total {
		select {
		case <-ch1:
			count1++
		case <-ch2:
			count2++
		}
	}

	ratio1 := float64(count1) / float64(total) * 100
	ratio2 := float64(count2) / float64(total) * 100

	t.Logf("\nch1被选中: %.3f%%(%d 次)\nch2被选中: %.3f%%(%d 次)\n偏离度: %.2f%%", ratio1, count1, ratio2, count2, abs(ratio1-ratio2))
}

func abs(f float64) float64 {
	if f < 0 {
		return -f
	}
	return f
}
