package day1

func Part2(input string) int {
	count := 0
	for i, c := range input {
		if c == '(' {
			count++
		} else if c == ')' {
			count--
		}
		if count < 0 {
			return i + 1
		}
	}
	return -1
}
