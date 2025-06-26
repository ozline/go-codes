package code

/*
	合并 map 的几种用法
*/

import (
	"fmt"
	"maps"
)

func MergeMaps() {
	src := map[string]int{
		"one": 1,
		"two": 2,
	}
	dst := map[string]int{
		"two":   42,
		"three": 3,
	}

	maps.Copy(dst, src)

	fmt.Println("src:", src)
	fmt.Println("dst:", dst)

	// src: map[one:1 two:2]
	// dst: map[one:1 three:3 two:2]
	// 使用 maps.Copy() 会对同样的 key 进行覆盖操作

	// 在 go 1.18 - go1.20 中，也可以使用 golang.org/x/exp/maps.Copy 函数，但是这个 Copy 函数要求 map 键类型必须是具体类型而不是接口类型
	// 1.21 后直接用 maps.Copy() 是支持使用接口类型的

	// 也可以使用下面的 mergeMaps 函数手动合并
}

// nolint
func mergeMaps(dst, src map[string]int) {
	for key, value := range src {
		dst[key] = value
	}
}
