package goutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseASCII(t *testing.T) {
	str := "abcd"
	dest := ReverseASCII(str)
	assert.Equal(t, dest, "dcba")

	str = "abc"
	dest = ReverseASCII(str)
	assert.Equal(t, dest, "cba")

	str = ""
	dest = ReverseASCII(str)
	assert.Equal(t, dest, "")

	str = "a"
	dest = ReverseASCII(str)
	assert.Equal(t, dest, "a")

}
