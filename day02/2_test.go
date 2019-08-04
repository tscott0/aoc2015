package day02

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample3(t *testing.T) {
	a := assert.New(t)
	a.Equal(34, Part2("2x3x4"))
}

func TestExample4(t *testing.T) {
	a := assert.New(t)
	a.Equal(14, Part2("1x1x10"))
}

func TestPart2(t *testing.T) {
	inputBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		assert.FailNow(t, err.Error())
	}

	t.Log(Part2(string(inputBytes)))
}
