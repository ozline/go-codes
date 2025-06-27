package code

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

// 定义一个复数类型
type cnum struct {
	r, i float64
}

var (
	// 定义三个全局变量
	x        cnum
	y        cnum
	z        cnum
	mutex_y1 sync.Mutex
	mutex_y2 sync.Mutex
	mutex_z  sync.Mutex

	passed atomic.Int64 = atomic.Int64{} // 用于记录已经执行完的goroutine数量
	// 这里想用一个channel来实现解锁 foo3 和 foo4，应该如何实现呢？
	done     chan struct{} = make(chan struct{}, 1) // 同步foo1、2 和 foo3
	doneFoo3 chan struct{} = make(chan struct{}, 1) // 同步foo3 和 foo4
)

// 设计中，执行顺序应该是：foo1 和 foo2 同时执行，foo3 等待 foo1 和 foo2 执行完毕，foo4 等待 foo3 执行完毕
// 有难点的实际上是 mutex_y1 和 mutex_y2

func TestMutex(t *testing.T) {
	// 运行 10 次
	var allRight bool = true
	for i := 0; i < 10; i++ {
		x = cnum{2, 3}
		y = cnum{1, 1}
		z = cnum{0, 0}

		fmt.Printf("Run %d:\n", i+1)
		results := make([]string, 5)

		// results[0] = foo1()
		// results[1] = foo2()
		// results[2], results[3] = foo3()
		// results[4] = foo4()

		var wg sync.WaitGroup
		wg.Add(4)
		go func(index int) {
			defer wg.Done()
			results[index] = foo1()
		}(0)
		go func(index int) {
			defer wg.Done()
			results[index] = foo2()
		}(1)
		go func(index int) {
			defer wg.Done()
			results[index], results[index+1] = foo3()
		}(2)
		go func(index int) {
			defer wg.Done()
			results[index] = foo4()
			// results[index] = "foo4 x = 3.000000 + 3.000000i"
		}(4)
		wg.Wait()
		fmt.Println()

		// 断言输出
		expected := []string{
			"foo1 w = 3.000000 + 4.000000i",
			"foo2 w = 1.000000 + 1.000000i",
			"foo3 z = 1.000000 + 1.000000i",
			"foo3 y = 2.000000 + 2.000000i",
			"foo4 x = 3.000000 + 3.000000i",
		}

		// 验证结果
		for j, result := range results {
			if result != expected[j] {
				fmt.Printf("Assertion failed: expected '%s', got '%s'\n", expected[j], result)
				break
			}
		}

		fmt.Println()
	}

	if !allRight {
		fmt.Println("Some assertions failed")
	} else {
		fmt.Println("All assertions passed")
	}
}

// 添加操作
func add(p cnum, q cnum) cnum {
	var s cnum
	s.r = p.r + q.r
	s.i = p.i + q.i
	return s
}

// foo1 负责计算 x 和 y 的和，并存储在局部变量 w 中
func foo1() string {
	mutex_y1.Lock()
	w := add(x, y)
	mutex_y1.Unlock()
	result := fmt.Sprintf("foo1 w = %f + %fi", w.r, w.i)
	fmt.Println(result) // 输出结果
	if passed.Add(1)%2 == 0 {
		done <- struct{}{}
	}
	return result
}

// foo2 负责计算 y 和 z 的和，并存储在局部变量 w 中
func foo2() string {
	mutex_y2.Lock()
	w := add(y, z)
	mutex_y2.Unlock()
	result := fmt.Sprintf("foo2 w = %f + %fi", w.r, w.i)
	fmt.Println(result) // 输出结果
	if passed.Add(1)%2 == 0 {
		done <- struct{}{}
	}
	return result
}

// foo3 负责计算 z 和 w 的和，并存储在全局变量 y 中
func foo3() (string, string) {
	<-done // 等待 foo1 和 foo2 执行完毕

	var w cnum = cnum{1, 1}
	mutex_z.Lock()
	z = add(z, w)
	mutex_z.Unlock()
	mutex_y1.Lock()
	mutex_y2.Lock()
	y = add(y, w)
	mutex_y2.Unlock()
	mutex_y1.Unlock()
	result1 := fmt.Sprintf("foo3 z = %f + %fi", z.r, z.i)
	result2 := fmt.Sprintf("foo3 y = %f + %fi", y.r, y.i)
	fmt.Println(result1) // 输出结果
	fmt.Println(result2) // 输出结果

	doneFoo3 <- struct{}{}

	return result1, result2
}

// foo4 负责计算 y 和 z 的和，并存储在全局变量 x 中
func foo4() string {
	<-doneFoo3 // 等待 foo3 执行完毕

	mutex_y1.Lock()
	mutex_y2.Lock()
	mutex_z.Lock()
	x = add(y, z)
	mutex_z.Unlock()
	mutex_y2.Unlock()
	mutex_y1.Unlock()
	result := fmt.Sprintf("foo4 x = %f + %fi", x.r, x.i)
	fmt.Println(result) // 输出结果
	return result
}
