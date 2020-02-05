package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func Part1(input string) int {
	niceCount := 0

	lines := strings.Split(input, "\n")

	for _, l := range lines {
		if nice1(l) {
			niceCount++
		}
	}

	return niceCount
}

func nice1(input string) bool {
	hasVowels := false
	hasDouble := false
	hasBadString := false

	vowelCount := 0
	p := ' '
	for _, c := range input {
		if c == 'a' ||
			c == 'e' ||
			c == 'i' ||
			c == 'o' ||
			c == 'u' {
			vowelCount++
		}

		if vowelCount >= 3 {
			hasVowels = true
		}

		if c == p {
			hasDouble = true
		}

		if (p == 'a' && c == 'b') ||
			(p == 'c' && c == 'd') ||
			(p == 'p' && c == 'q') ||
			(p == 'x' && c == 'y') {
			hasBadString = true
		}

		p = c
	}

	return hasVowels && hasDouble && !hasBadString
}

func Part2(input string) int {
	niceCount := 0

	lines := strings.Split(input, "\n")

	for _, l := range lines {
		if nice2(l) {
			niceCount++
		}
	}

	return niceCount
}

func nice2(input string) bool {
	return false
}

func main() {

	dat, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(Part1(string(dat)))
}
