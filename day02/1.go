package day02

import (
	"log"
	"strconv"
	"strings"
)

func Part1(input string) int {
	lines := strings.Split(input, "\n")

	totalArea := 0
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

		sideA := h * w
		sideB := h * d
		sideC := d * w

		lineArea := sideA*2 + sideB*2 + sideC*2

		if sideA <= sideB {
			if sideA <= sideC {
				lineArea += sideA
			} else {
				lineArea += sideC
			}
		} else {
			if sideB <= sideC {
				lineArea += sideB
			} else {
				lineArea += sideC
			}
		}

		totalArea += lineArea
	}

	return totalArea
}
