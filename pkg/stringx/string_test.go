package stringx

import "testing"

func TestStrCopyToPtr(t *testing.T) {
	var a = make(map[string]any)
	err := StrCopyToPtr("abc", &a)
	t.Log("err: ", err)
}
