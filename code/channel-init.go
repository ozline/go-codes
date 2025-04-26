package code

// 在已初始化、未初始化、已关闭的 Channel 中进行读写的测试

import (
	"fmt"
)

var (
	InitChannel    = make(chan int, 1) // 初始化了的 channel
	NotInitChannel chan int            // 未初始化的 channel
)

func ChannelInit() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	// 事实上没有观测到 Panic，我不知道是不是 go 版本的问题，至少在 1.19的时候没有观测到

	Test10()

	// 测试已初始化的 Channel
	WriteToChannel(InitChannel) // 正常
	// ReadFromChannel(InitChannel)    // 如果没有WriteToChannel，会阻塞，然后报 fatal error: all goroutines are asleep - deadlock!

	// 测试未初始化的 Channel
	// ReadFromChannel(NotInitChannel) // fatal error: all goroutines are asleep - deadlock!
	// WriteToChannel(NotInitChannel) // fatal error: all goroutines are asleep - deadlock!

	// 测试已关闭的 Channel
	ClosedChannelTest()
}

func WriteToChannel(channel chan int) {
	// 向 channel 写入数据
	channel <- 1
}

func ReadFromChannel(channel chan int) {
	// 从 channel 读取数据
	<-channel
}

func ClosedChannelTest() {
	closedChannel := make(chan int, 2)
	closedChannel <- 10
	closedChannel <- 20
	close(closedChannel)

	// 从已关闭的 channel 读取数据
	fmt.Println("Reading from closed channel:")
	fmt.Println(<-closedChannel) // 输出 10
	fmt.Println(<-closedChannel) // 输出 20
	fmt.Println(<-closedChannel) // 输出 0（channel 类型的零值）

	// 向已关闭的 channel 写入数据（会 panic）
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from writing to closed channel:", r)
		}
	}()
	fmt.Println("Writing to closed channel:")
	closedChannel <- 30 // 这里会 panic
}

func Test10() {
	const (
		x = iota
		_
		y
		z = "pi"
		k
		p = iota
		q
	)
	fmt.Println(x, y, z, k, p, q)
}
