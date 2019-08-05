package day03

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	a := assert.New(t)
	a.Equal(2, Part1(">"))
}

func TestExample2(t *testing.T) {
	a := assert.New(t)
	a.Equal(4, Part1("^>v<"))
}

func TestExample3(t *testing.T) {
	a := assert.New(t)
	a.Equal(2, Part1("^v^v^v^v^v"))
}

func TestPart1(t *testing.T) {
	inputBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		assert.FailNow(t, err.Error())
	}

	t.Log(Part1(string(inputBytes)))
}
