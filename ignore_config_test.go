package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IgnoreConfig(t *testing.T) {
	config := NewIgnoreConfig()
	config.AddIgnore("dir1/dir2/dir3/file")
	assert.True(t, config.ShouldIgnore("dir1/dir2/dir3/file"))
	assert.True(t, config.ShouldIgnore("./dir1/dir2/dir3/file"))
	assert.True(t, config.ShouldIgnore("./dir1/dir2/dir3/../../dir2/dir3/file"))

	config.DelIgnore("./dir1/dir2/dir3/../../dir2/dir3/file")
	assert.False(t, config.ShouldIgnore("dir1/dir2/dir3/file"))
	assert.False(t, config.ShouldIgnore("./dir1/dir2/dir3/file"))
	assert.False(t, config.ShouldIgnore("./dir1/dir2/dir3/../../dir2/dir3/file"))
}
