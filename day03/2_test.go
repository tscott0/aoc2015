package day03

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample4(t *testing.T) {
	a := assert.New(t)
	a.Equal(3, Part2("^v"))
}

func TestExample5(t *testing.T) {
	a := assert.New(t)
	a.Equal(3, Part2("^>v<"))
}

func TestExample6(t *testing.T) {
	a := assert.New(t)
	a.Equal(11, Part2("^v^v^v^v^v"))
}
func TestPart2(t *testing.T) {
	inputBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		assert.FailNow(t, err.Error())
	}

	t.Log(Part2(string(inputBytes)))
}
