package code

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestBytesBuffer(t *testing.T) {
	var compactedJSON bytes.Buffer
	// 测试直接就是空的 bytes.Buffer 会输出的文本
	t.Logf("compactedJSON: [%s]", compactedJSON.String())

	// 如果是空的 json，这里会报错 json.Compact failed: unexpected end of JSON input
	err := json.Compact(&compactedJSON, []byte(""))
	if err != nil {
		// 就当无事发生
		// t.Errorf("json.Compact failed: %v", err)
	} else {
		t.Logf("compactedJSON after compact: [%s]", compactedJSON.String())
	}
}
