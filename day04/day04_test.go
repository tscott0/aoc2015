package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	a := assert.New(t)

	a.Equal(609043, Run1("abcdef", "00000"))
}

func TestExample2(t *testing.T) {
	a := assert.New(t)

	a.Equal(1048970, Run1("pqrstuv", "00000"))
}
