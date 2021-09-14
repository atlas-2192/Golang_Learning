package foo

import "strconv"

func findNumbers(nums []int) int {
	count := 0
	for _, num := range nums {
		s := strconv.Itoa(num)
		l := len(s)
		if l > 2 && s[0] == '-' {
			l--
		}
		if l%2 == 0 {
			count++
		}
	}
	return count
}
