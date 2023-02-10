package stringx

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStrCopyToPtr(t *testing.T) {
	type Student struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	var a Student
	b, _ := json.Marshal(Student{
		Age:  24,
		Name: "totemli",
	})
	s := string(b)
	err := StrCopyToPtr(s, &a)
	require.NoError(t, err)
}
