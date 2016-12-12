package main

import (
	"testing"
	"github.com/alecthomas/assert"
)

func TestDevelopmentTeam(t *testing.T) {


	p := people()
	assert.Equal(t, 9, len(p))
	//Find Test for this
	//
}