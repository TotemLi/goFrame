package lang

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepr(t *testing.T) {
	st := struct {
		Name string
		Age  int
	}{
		Name: "TotemLi",
		Age:  23,
	}

	assert.Equal(t, "23.4", Repr(23.4))
	assert.Equal(t, "{TotemLi 23}", Repr(st))
}
