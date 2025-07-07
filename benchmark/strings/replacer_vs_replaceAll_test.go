package code

import (
	"strings"
	"testing"
)

// 主要测试 replacer 的 cache 功能，相较于直接使用 strings.ReplaceAll，strings.Replacer 支持缓存，按理说可以提升显著性能

// 测试结果表明起码提升 1 倍以上的性能

/*
	strings.ReplaceAll: 遍历整个子串，逐个查找目标子串并替换，每次调用时都会从头开始执行
	strings.Replacer：预先创建一个 Replacer 对象，内部维护了一个缓存，可以在多次替换时复用这个对象，从而减少重复的查找和替换操作

	replacer 的内部数据结构（1.24）
*/

func BenchmarkReplaceAll(b *testing.B) {
	// 测试数据
	input := "hello world, hello gopher, hello universe"
	old := "hello"
	new := "hi"

	// 基准测试 strings.ReplaceAll
	for b.Loop() {
		_ = strings.ReplaceAll(input, old, new)
	}
}

func BenchmarkReplacer(b *testing.B) {
	// 测试数据
	input := "hello world, hello gopher, hello universe"
	old := "hello"
	new := "hi"

	// 创建 Replacer 对象
	replacer := strings.NewReplacer(old, new)

	// 基准测试 strings.Replacer
	for b.Loop() {
		_ = replacer.Replace(input)
	}
}

/*
	goos: darwin
	goarch: arm64
	pkg: github.com/ozline/go-codes/language
	cpu: Apple M1 Pro
	=== RUN   BenchmarkReplaceAll
	BenchmarkReplaceAll
	BenchmarkReplaceAll-10           9238455               124.1 ns/op            32 B/op          1 allocs/op
	PASS
	ok      github.com/ozline/go-codes/language     1.463s


	goos: darwin
	goarch: arm64
	pkg: github.com/ozline/go-codes/language
	cpu: Apple M1 Pro
	=== RUN   BenchmarkReplacer
	BenchmarkReplacer
	BenchmarkReplacer-10            14662659                78.77 ns/op           40 B/op          2 allocs/op
	PASS
	ok      github.com/ozline/go-codes/language     1.427s
*/
