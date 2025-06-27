package code

import (
	"strings"
	"testing"
)

/* 这份代码主要用来测试 StringReplacer 的表现 */
func TestStringReplacer(t *testing.T) {
	tempString := "测试一下 StringReplacer"
	replacer := strings.NewReplacer("String", "Replacer", "Replacer", "String")
	result := replacer.Replace(tempString)
	t.Logf("Original: %s, Replaced: %s", tempString, result)

	// 我们倒过来替换试一下
	replacer = strings.NewReplacer("Replacer", "NewReplacer", "String", "Replacer")
	result = replacer.Replace(tempString)
	t.Logf("Original: %s, Reversed Replaced: %s", tempString, result)

	// 判断是否可以多次替换
	multiString := "测试一下 StringReplacer StringReplacer"
	replacer = strings.NewReplacer("String", "Replacer", "Replacer", "String")
	multiResult := replacer.Replace(multiString)
	t.Logf("Original: %s, Multi Replaced: %s", multiString, multiResult)
}

// 结果可以注意到，它是使用了一对替换字符串，scan 到第一个字符串的时候就会替代，对顺序不敏感
// 比如倒过来替换的测试，如果对顺序敏感，那么会在 Replacer 替换为 NewReplacer 后无法再对 String 进行替换，但事实是可以
// Original: 测试一下 StringReplacer, Reversed Replaced: 测试一下 ReplacerNewReplacer
