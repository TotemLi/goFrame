package conf

import (
	"github.com/stretchr/testify/require"
	"testing"
)

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
	require.NoError(t, err)
	t.Log("err: ", err)
	t.Log("val: ", val)
}
