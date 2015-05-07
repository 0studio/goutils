package goutils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTea(t *testing.T) {
	var in string = "hasdfasdfasdasdfasdfsaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaello"
	var key string = "12key1111100000000" // must len(key)%16==0
	out := TeaEncrypt(in, key)

	fmt.Println(out, len(out))
	fmt.Println("out", out)
	outout := TeaDecrypt(out, key)
	fmt.Println("decode", outout)
	assert.True(t, outout == in)

}
