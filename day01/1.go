package day01

func Part1(input string) int {
	count := 0
	for _, c := range input {
		if c == '(' {
			count++
		} else if c == ')' {
			count--
		}
	}
	return count
}
