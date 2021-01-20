package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CommaClean(t *testing.T) {
	assert.Equal(t, "A,B,C,D", CommaClean("A,B,C,D"))
	assert.Equal(t, "A,B,C,D", CommaClean("A, B, C, D"))
	assert.Equal(t, "A,B,C,D", CommaClean("A, \t B, \t C, \t D"))
	assert.Equal(t, "A,B,C,D", CommaClean("A \t , \t B \t , \t C \t , \t D"))
}

func Test_CheckSignleBuildFlag(t *testing.T) {
	assert.True(t, CheckSignleBuildFlag("!DDE_TEST", "// +build !DDE_TEST"))
	assert.True(t, CheckSignleBuildFlag("!DDE_TEST", "    //    +build    !DDE_TEST    "))
	assert.True(t, CheckSignleBuildFlag("!DDE_TEST", " \t // \t +build \t !DDE_TEST \t "))
	assert.True(t, CheckSignleBuildFlag("!DDE_TEST", " \t // \t +build \t !DDE_TEST \t \n"))

	assert.False(t, CheckSignleBuildFlag("!DDE_TEST", "// +build DDE_TEST"))
	assert.False(t, CheckSignleBuildFlag("!DDE_TEST", "// +build !DDE_TESTn"))
	assert.False(t, CheckSignleBuildFlag("!DDE_TEST", "// +build FLAG1 !DDE_TEST"))
	assert.False(t, CheckSignleBuildFlag("!DDE_TEST", "// +build !DDE_TEST FLAG2"))
	assert.False(t, CheckSignleBuildFlag("!DDE_TEST", "// +build FLAG1 !DDE_TEST FLAG2"))
	assert.False(t, CheckSignleBuildFlag("!DDE_TEST", "// +build FLAG1,!DDE_TEST,FLAG2"))
}

func Test_CheckMultiBuildFlag(t *testing.T) {
	assert.True(t, CheckMultiBuildFlag("!DDE_TEST", "// +build !DDE_TEST"))
	assert.True(t, CheckMultiBuildFlag("!DDE_TEST", "    //    +build    !DDE_TEST    "))
	assert.True(t, CheckMultiBuildFlag("!DDE_TEST", " \t // \t +build \t !DDE_TEST \t "))
	assert.True(t, CheckMultiBuildFlag("!DDE_TEST", " \t // \t +build \t !DDE_TEST \t \n"))

	assert.False(t, CheckSignleBuildFlag("!DDE_TEST", "// +build DDE_TEST"))
	assert.False(t, CheckSignleBuildFlag("!DDE_TEST", "// +build !DDE_TESTn"))
	assert.False(t, CheckSignleBuildFlag("!DDE_TEST", "// +build FLAG1 !DDE_TEST"))
	assert.False(t, CheckSignleBuildFlag("!DDE_TEST", "// +build !DDE_TEST FLAG2"))
	assert.False(t, CheckSignleBuildFlag("!DDE_TEST", "// +build FLAG1 !DDE_TEST FLAG2"))
	assert.False(t, CheckSignleBuildFlag("!DDE_TEST", "// +build FLAG1,!DDE_TEST,FLAG2"))

	assert.True(t, CheckMultiBuildFlag("!DDE_TEST", "// +build FLAG1, !DDE_TEST"))
	assert.True(t, CheckMultiBuildFlag("!DDE_TEST", "// +build FLAG1, !DDE_TEST, FLAG2"))
	assert.True(t, CheckMultiBuildFlag("!DDE_TEST", "// +build !DDE_TEST, FLAG2"))
	assert.True(t, CheckMultiBuildFlag("!DDE_TEST", "    // +build FLAG1, !DDE_TEST    "))
	assert.True(t, CheckMultiBuildFlag("!DDE_TEST", "    // +build FLAG1, !DDE_TEST, FLAG2    "))
	assert.True(t, CheckMultiBuildFlag("!DDE_TEST", "    // +build !DDE_TEST, FLAG2    "))

	assert.False(t, CheckMultiBuildFlag("!DDE_TEST", "    // +build FLAG1 !DDE_TEST"))
	assert.False(t, CheckMultiBuildFlag("!DDE_TEST", "    // +build FLAG1 !DDE_TEST FLAG2"))
	assert.False(t, CheckMultiBuildFlag("!DDE_TEST", "    // +build !DDE_TEST FLAG2"))
}

func Test_RemoveFlag(t *testing.T) {
	assert.Equal(t, "", RemoveFlag("!DDE_TEST", "// +build !DDE_TEST"))
	assert.Equal(t, "// +build FLAG1, FLAG2", RemoveFlag("!DDE_TEST", "// +build FLAG1, !DDE_TEST, FLAG2"))
	assert.Equal(t, "// +build FLAG1,FLAG2", RemoveFlag("!DDE_TEST", "// +build FLAG1,!DDE_TEST,FLAG2"))
	assert.Equal(t, "// +build FLAG1 !DDE_TEST FLAG2", RemoveFlag("!DDE_TEST", "// +build FLAG1 !DDE_TEST FLAG2"))
}
