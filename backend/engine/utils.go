package engine

import (
	"strconv"
	"unicode"
)

/*
 * Check if a certain number is in an array
 */
func isInArray(nums []int, n int) bool {
	for _, x := range nums {
		if x == n {
			return true
		}
	}

	return false
}

/*
 * Get sum of the numbers that are represented as string
 */
func sum(numStrings []string) int {
	s := 0

	for _, x := range numStrings {
		num, _ := strconv.Atoi(x)
		s += num
	}

	return s
}

/*
 * Trim non letter characters on the left
*/
func ltrimNonLetter(str string) string {
	pos := 0

	for _, r := range str {
		if (unicode.IsLetter(r)) {
			break
		}

		pos ++
	}

	return str[pos:]
}