package day02

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

func Part2(input string) int {
	lines := strings.Split(input, "\n")

	totalRibbon := 0
	for _, l := range lines {
		dimensions := strings.Split(l, "x")
		if len(dimensions) != 3 {
			log.Fatal("invalid dimensions: ", l)
		}

		h, err := strconv.Atoi(dimensions[0])
		if err != nil {
			log.Fatal("invalid dimensions: ", l)
		}

		w, err := strconv.Atoi(dimensions[1])
		if err != nil {
			log.Fatal("invalid dimensions: ", l)
		}

		d, err := strconv.Atoi(dimensions[2])
		if err != nil {
			log.Fatal("invalid dimensions: ", l)
		}

		dims := []int{h, w, d}
		sort.Ints(dims)

		totalRibbon += 2*(dims[0]+dims[1]) + h*w*d
	}

	return totalRibbon
}
