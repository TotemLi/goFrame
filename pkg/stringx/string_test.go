package stringx

import (
	"encoding/json"
	"testing"
)

func TestStrCopyToPtr(t *testing.T) {
	type Student struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	var a Student
	//var a []string
	//s := fmt.Sprintf("%v", map[string]any{"a": "123"})
	b, _ := json.Marshal(Student{
		Age:  24,
		Name: "totemli",
	})
	s := string(b)
	err := StrCopyToPtr(s, &a)
	t.Log("err: ", err)
	t.Log("a: ", a)
}
