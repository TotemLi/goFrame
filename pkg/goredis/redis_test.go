package goredis

import (
	"context"
	"encoding/json"
	"testing"
)

func TestNewClient(t *testing.T) {
	c := NewClient("localhost:6379", "MQgnbBmJ3tsjiiR4", 0)
	ctx := context.Background()
	//d := struct {
	//	A string
	//	B int
	//}{
	//	A: "a",
	//	B: 123,
	//}
	//v := map[string]any{
	//	"a": "a",
	//	"b": 123,
	//}
	//v := "abc"
	//v := []byte("abc")
	v, e := json.Marshal(12.345)

	t.Log(string(v))
	t.Log("err: ", e)
	err := c.Set(ctx, "a", v)
	t.Log("err: ", err)
}
