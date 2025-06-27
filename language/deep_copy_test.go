package code

import (
	"reflect"
	"testing"
)

// 测试深拷贝函数，这里提供了一个特别万能的 map[string]any 来作为例子

// 给未来的自己发散一下思维，如果要实现一个map，这个 map 的 k-v 对的 value 属性只有 2 种可能，一种是 string，另一种是这个 map 类型本身
// 那么能否实现这样的一个强类型，而不使用 map[string]any 呢？

func TestDeepCopyMapAny(t *testing.T) {
	var data map[string]any = map[string]any{
		"key1": "value1",
		"key2": 123,
		"key3": true,
		"key4": []string{"item1", "item2"},
		"key5": nil,
		"key6": map[string]any{
			"nestedKey1": "nestedValue1",
			"nestedKey2": 456,
			"nestedKey3": false,
		},
		"key7": map[string]any{
			"nestedKey4": "nestedValue2",
			"nestedKey5": map[string]any{
				"deepNestedKey1": "deepNestedValue1",
				"deepNestedKey2": map[string]any{
					"deepDeepNestedKey1": "deepDeepNestedValue1",
				},
			},
		},
	}

	// 调用深拷贝函数
	copiedData := DeepCopyMapAny(data)

	// 验证深拷贝结果
	if !reflect.DeepEqual(data, copiedData) {
		t.Errorf("Deep copy failed: copied data does not match original")
	}

	// 验证修改原始数据不会影响拷贝后的数据
	data["key1"] = "modifiedValue"
	if copiedData["key1"] == "modifiedValue" {
		t.Errorf("Deep copy failed: copied data was affected by changes to original")
	}

	// 验证嵌套 map 的深拷贝
	nestedMap := data["key6"].(map[string]any)
	nestedMap["nestedKey1"] = "modifiedNestedValue"
	if copiedData["key6"].(map[string]any)["nestedKey1"] == "modifiedNestedValue" {
		t.Errorf("Deep copy failed: copied nested map was affected by changes to original")
	}

	// 验证嵌套的深层结构
	deepNestedMap := data["key7"].(map[string]any)["nestedKey5"].(map[string]any)["deepNestedKey2"].(map[string]any)
	deepNestedMap["deepDeepNestedKey1"] = "modifiedDeepDeepNestedValue"
	if copiedData["key7"].(map[string]any)["nestedKey5"].(map[string]any)["deepNestedKey2"].(map[string]any)["deepDeepNestedKey1"] == "modifiedDeepDeepNestedValue" {
		t.Errorf("Deep copy failed: copied deep nested map was affected by changes to original")
	}

}

// DeepCopyMapAny 对 map[string]any 进行任意拷贝
func DeepCopyMapAny(original map[string]any) map[string]any {
	// 创建一个新的 map
	copied := make(map[string]any)

	for key, value := range original {
		switch v := value.(type) {
		case map[string]any:
			// 如果值是 map[string]any 类型，递归拷贝
			copied[key] = DeepCopyMapAny(v)
		case []any:
			// 如果值是 []any 类型，递归处理每个元素
			copiedSlice := make([]any, len(v))
			for i, item := range v {
				copiedSlice[i] = deepCopyValue(item)
			}
			copied[key] = copiedSlice
		case []string:
			// 如果值是 []string 类型，直接拷贝（字符串是不可变的）
			copied[key] = append([]string{}, v...)
		default:
			// 对于其他值，直接赋值（基本类型或不可变类型）
			copied[key] = deepCopyValue(v)
		}
	}

	return copied
}

// 深拷贝辅助函数，处理单个值
func deepCopyValue(value any) any {
	switch v := value.(type) {
	case map[string]any:
		return DeepCopyMapAny(v)
	case []any:
		copiedSlice := make([]any, len(v))
		for i, item := range v {
			copiedSlice[i] = deepCopyValue(item)
		}
		return copiedSlice
	case []string:
		return append([]string{}, v...)
	default:
		// 对于基本类型或不可变类型，直接返回
		return v
	}
}
