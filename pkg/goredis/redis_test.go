package goredis

import (
	"context"
	"github.com/stretchr/testify/require"
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
	var err error
	err = c.Set(ctx, "totemli", stu)
	require.NoError(t, err)
	var stu2 Student
	err = c.Get(ctx, "totemli", &stu2)
	require.NoError(t, err)
	require.EqualValues(t, stu, stu2)
}
