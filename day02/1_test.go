package day02

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	a := assert.New(t)
	a.Equal(58, Part1("2x3x4"))
}

func TestExample2(t *testing.T) {
	a := assert.New(t)
	a.Equal(43, Part1("1x1x10"))
}

func TestPart1(t *testing.T) {
	inputBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		assert.FailNow(t, err.Error())
	}

	t.Log(Part1(string(inputBytes)))
}
