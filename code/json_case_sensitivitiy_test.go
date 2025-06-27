package code

/*
	面试遇到过的，go JSON tag 大小写敏感问题
*/

import (
	"encoding/json"
	"fmt"
	"testing"
)

type TestJSON struct {
	UserID string `json:"UserID"`
}

func TestJSONSensitive(t *testing.T) {
	var a = `{"UserID": "111234", "UserId": "222222", "UserID": "33333", "userid": "44444"}`

	// 主要是大小写不敏感问题，会匹配到最后一个，也就是 userid 这个东西

	// 解决方案：
	// 1. 使用第三方库，e.g. github.com/json-iterator/go
	// 2. 预处理 JSON 字符串，避免这个问题产生
	// 3. 使用自定义的 json.Unmarshal 函数，对于这个实现，实际上就是将变量设置为 map[string]interface{}，之后按 map 解析

	var temp TestJSON

	if err := json.Unmarshal([]byte(a), &temp); err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", temp.UserID)
}
