package code

import (
	"encoding/json"
	"fmt"
	"testing"
)

// 测试对 map[string]any 的 JSON marshal 行为
func TestJSONMarshalMap(t *testing.T) {
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
	resp, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("Failed to marshal map[string]any: %v", err)
	}
	fmt.Printf("Marshaled JSON: %s\n", string(resp))
	// 这里可以看到，map[string]any 的 JSON marshal 行为是正确的
}

func TestJSONMarshal(t *testing.T) {
	var data map[string]any
	err := json.Unmarshal([]byte(rawString), &data)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	type testStructStrongType struct {
		Type string         `json:"type"`
		Data map[string]any `json:"data"`
	}

	res := testStructStrongType{
		Type: "whocares",
		Data: data,
	}

	var resp []byte
	resp, err = json.Marshal(res)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}

	fmt.Printf("Marshaled JSON with strong type: %v\n\n\n\n\n", string(resp))

	type testStructString struct {
		Type string `json:"type"`
		Data string `json:"data"`
	}

	remarshalledData, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("Failed to marshal data: %v", err)
	}

	resString := testStructString{
		Type: "whocares",
		Data: string(remarshalledData),
	}

	respString, err := json.Marshal(resString)
	if err != nil {
		t.Fatalf("Failed to marshal string data: %v", err)
	}
	fmt.Printf("Marshaled JSON with strong type: %v", string(respString))
}

// 可以注意到如果对 Struct 中的数据进行强类型定义（即使是 map[string]any），解析时是不会出现斜杠的
// 但如果将数据作为 string 进行解析，则会出现斜杠

// 从某个文档偷的原始大长串 string
const rawString = `{"type":"card","children":[{"type":"relative","children":[{"type":"view","background":{"background_color":"#99601E#99601E"},"width":-2,"height":-2},{"type":"image","width":-2,"height":135,"image_url":"https://{sensititive_solve}/heybox/2023/12/17/c529d4967ac785f876c2145bcea6d966.jpg"},{"type":"view","margin":{"top":65},"background":{"background_gradient":{"background_color_start":"#0099601E#0099601E","background_color_end":"#99601E#99601E","background_color_orientation":2}},"width":-2,"height":70},{"type":"image","width":-2,"height":-2,"image_url":"https://{sensititive_solve}/oa/2024/12/03/9bbf221f34b86de91e01ad4ffe271501.png"},{"type":"text","text":"\" CS2 \"","text_color":"#FFE64C#FFE64C","text_size":14,"margin":{"top":245},"layout_gravity":1},{"type":"image","margin":{"bottom":88,"left":20,"right":20},"width":-2,"height":1,"image_url":"https://{sensititive_solve}/oa/2024/12/03/cb6ece1c1477593644a60255fd10a124.png","view_id":"line","layout_gravity":7},{"type":"image","margin":{"bottom":12},"width":82,"height":18,"image_url":"https://{sensititive_solve}/oa/2024/10/09/d0336112afc5861efc75e7db40b00387.png","view_id":"bottom_logo","layout_gravity":7},{"type":"linear","children":[{"type":"linear","orientation":1,"children":[{"type":"image","width":12,"height":12,"image_url":"https://{sensititive_solve}/oa/2024/12/03/c90077f7a1b97e8b387db0cf667fbb07.png"},{"type":"linear","orientation":0,"children":[{"type":"text","text":"0","text_color":"#FFE64C#FFE64C","text_size":13,"font_name":"Helvetica-Bold","margin":{"right":2}},{"type":"text","text":"/","text_color":"#4DFFFFFF#4DFFFFFF","text_size":13,"font_name":"Helvetica"},{"type":"text","text":"0","text_color":"#4DFFFFFF#4DFFFFFF","text_size":13,"font_name":"Helvetica"}],"margin":{"top":4},"gravity":4}],"gravity":4,"weight":1},{"type":"linear","orientation":1,"children":[{"type":"image","width":12,"height":12,"image_url":"https://{sensititive_solve}/oa/2024/12/03/428c596f06ee86760c944c707a1bbc86.png"},{"type":"linear","orientation":0,"children":[{"type":"text","text":"1","text_color":"#E1EDED#E1EDED","text_size":13,"font_name":"Helvetica-Bold","margin":{"right":2}},{"type":"text","text":"/","text_color":"#4DFFFFFF#4DFFFFFF","text_size":13,"font_name":"Helvetica"},{"type":"text","text":"1","text_color":"#4DFFFFFF#4DFFFFFF","text_size":13,"font_name":"Helvetica"}],"margin":{"top":4},"gravity":4}],"gravity":4,"weight":1},{"type":"linear","orientation":1,"children":[{"type":"image","width":12,"height":12,"image_url":"https://{sensititive_solve}/oa/2024/12/03/538ac4c3eb408c9b7579adbfc1d3e814.png"},{"type":"text","text":"1530.7h","text_color":"#FFFFFF#FFFFFF","text_size":13,"font_name":"Helvetica-Bold","margin":{"top":4}}],"gravity":4,"weight":1}],"margin":{"bottom":48,"left":12,"right":12},"gravity":7,"width":-2,"height":34,"view_id":"vg_data","layout_gravity":7},{"type":"relative","children":[{"type":"linear","children":[{"type":"image","width":15,"height":15,"image_url":"https://{sensititive_solve}/heybox/avatar/steamcommunity/public/images/2022/12/28/ebb8b69ad6762b3f98a166df4c3ff72e.jpg","image_radius":1},{"type":"text","text":"OZLIINEX","text_color":"#B2FFFFFF#B2FFFFFF","text_size":12,"margin":{"left":4}},{"type":"text","text":"2023.10.21","text_color":"#B3FFE64C#B3FFE64C","text_size":12,"margin":{"left":4}}],"gravity":4,"width":-1,"height":15,"layout_gravity":4}],"margin":{"bottom":8,"left":12,"right":12},"width":-2,"height":15,"view_id":"vg_user","top_of":"line"},{"type":"relative","children":[{"type":"linear","children":[{"type":"image","width":14,"height":28,"image_url":"https://{sensititive_solve}/oa/2024/12/05/f839a35c893c79de9e44900e091a1f8a.png"},{"type":"image","margin":{"left":4,"right":4},"width":110,"height":25,"image_url":"https://{sensititive_solve}/oa/2024/12/05/d82749d20373a2dee938ad4e9ffd5d4a.png"},{"type":"image","width":14,"height":28,"image_url":"https://{sensititive_solve}/oa/2024/12/05/1a30edf4cb919a12e64147fbd1c75b0f.png"}],"gravity":4,"width":-1,"layout_gravity":4}],"margin":{"bottom":4},"width":-2,"height":34,"top_of":"vg_user"},{"type":"image","width":-2,"height":-2,"image_url":"https://{sensititive_solve}/oa/2025/05/21/6dec79a9c14de9d2095b9e1692e69282.png"}],"width":-2,"height":-2}],"width":300,"height":420,"view_id":"parent","corner_radius":0}`
