package encoding

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestYamlToJson(t *testing.T) {
	got, err := YamlToJson([]byte("a: foo\nb: 1\nc: ${FOO}\nd: abcd!@#$112"))
	assert.NoError(t, err)
	assert.Equal(
		t,
		"{\"a\":\"foo\",\"b\":1,\"c\":\"${FOO}\",\"d\":\"abcd!@#$112\"}\n",
		string(got),
	)
}
