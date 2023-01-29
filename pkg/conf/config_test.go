package conf

import "testing"

func TestLoadConfig(t *testing.T) {
	var val struct {
		A string `json:"a"`
		B int    `json:"b"`
		C []struct {
			A string `json:"a"`
			B int    `json:"b"`
		} `json:"c"`
	}
	err := LoadConfig("./temp.yaml", &val)
	t.Log("err: ", err)
	t.Log("val: ", val)
}
