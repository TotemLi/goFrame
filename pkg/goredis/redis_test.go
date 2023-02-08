package goredis

import (
	"context"
	"fmt"
	"testing"
)

func TestNewClient(t *testing.T) {
	c := NewClient("localhost:6379", "MQgnbBmJ3tsjiiR4", 0)
	ctx := context.Background()
	var s bool
	err := c.Get(ctx, "totemli", &s)
	if err != nil {
		if err == ErrNotFind {
			fmt.Println("not find")
			return
		}
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("s: ", s)
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
	//v, e := json.Marshal(12.345)
	//
	//t.Log(string(v))
	//t.Log("err: ", e)
	//err := c.Set(ctx, "a", v)
	//t.Log("err: ", err)
}
