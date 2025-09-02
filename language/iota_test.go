package code

import (
	"testing"
)

func TestIota(t *testing.T) {
	const (
		_ = 1 << (10 * iota)
		KB
		MB
		GB
		TB
		PB
		EB
		ZB // 从这开始 fmt.Println(ZB) 编译器就会提示 overflows, 但是实际上并不会溢出
		YB
		NB
		OB // OB 是 1 << (10 * 10)

		B1
		B2
		B3
		B4
		B5
		B6
		B7
		B8
		B9
		B10
		B11
	)

	t.Logf("Type of EB: %T", EB)
	// t.Logf("Value of B11: %v", B11) // 编译器会报错 overflows
}
