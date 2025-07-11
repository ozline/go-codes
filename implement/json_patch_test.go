package code

/*
	Json Patch Demo
*/

import (
	"fmt"
	"testing"

	jsonpatch "github.com/evanphx/json-patch"
)

type point struct {
	x float32
	y float32
}

func TestJSONPatch(t *testing.T) {
	original := []byte(`{"k": "v"}`)
	modified := []byte(`{"k": "v"}`)

	patch, err := jsonpatch.CreateMergePatch(original, modified)
	if err != nil {
		panic(err)
	}

	fmt.Println(patch)
}
