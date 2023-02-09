package goredis

import (
	"context"
	"testing"
)

func TestNewClient(t *testing.T) {
	c := NewClient("localhost:6379", "MQgnbBmJ3tsjiiR4", 0)
	ctx := context.Background()
	type Student struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	stu := Student{
		Name: "TotemLi",
		Age:  24,
	}
	c.Set(ctx, "totemli", stu)
	var stu2 Student
	err := c.Get(ctx, "totemli", &stu2)
	t.Log("err: ", err)
	t.Log("stu2: ", stu2)
}
