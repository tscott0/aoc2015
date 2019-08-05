package day03

func Part2(input string) int {
	type coord struct {
		x, y int
	}
	grid := map[coord]int{
		coord{x: 0, y: 0}: 2,
	}

	x, y := 0, 0
	xr, yr := 0, 0
	for i, d := range input {
		var currentLoc coord

		if i%2 == 0 {
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
			currentLoc = coord{x: x, y: y}
		} else {
			switch d {
			case '<':
				xr--
			case '>':
				xr++
			case '^':
				yr++
			case 'v':
				yr--
			}
			currentLoc = coord{x: xr, y: yr}
		}

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
