package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExamples1(t *testing.T) {
	a := assert.New(t)

	a.Equal(1, Part1("ugknbfddgicrmopn"))
	a.Equal(1, Part1("aaa"))
	a.Equal(0, Part1("jchzalrnumimnmhp"))
	a.Equal(0, Part1("haegwjzuvuyypxyu"))
	a.Equal(0, Part1("dvszwmarrgswjxmb"))
}

func TestExamples2(t *testing.T) {
	a := assert.New(t)

	a.Equal(1, Part2("qjhvhtzxzqqjkmpb"))
	a.Equal(1, Part2("xxyxx"))
	a.Equal(0, Part2("uurcxstgmygtbstg"))
	a.Equal(0, Part2("ieodomkazucvgmuy"))
}
