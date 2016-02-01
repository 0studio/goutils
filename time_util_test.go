package goutils

import "testing"
import "github.com/stretchr/testify/assert"

func TestIsLearYear(t *testing.T) {
	assert.True(t, IsLeapYear(2000))
	assert.True(t, IsLeapYear(1996))
	assert.False(t, IsLeapYear(1000))

}
