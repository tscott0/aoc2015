package day03

func Part1(input string) int {
	type coord struct {
		x, y int
	}
	grid := map[coord]int{
		coord{x: 0, y: 0}: 1,
	}

	x, y := 0, 0
	for _, d := range input {
		switch d {
		case '<':
			x--
		case '>':
			x++
		case '^':
			y++
		case 'v':
			y--
		}

		currentLoc := coord{x: x, y: y}

		if c, ok := grid[currentLoc]; ok {
			grid[currentLoc] = c + 1
		} else {
			grid[currentLoc] = 1
		}
	}

	duplicates := 0

	for _, c := range grid {
		if c > 0 {
			duplicates++
		}
	}

	return duplicates
}
