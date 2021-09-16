package foo

func replaceElements(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		max := 0
		found := false
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > max {
				max = arr[j]
				found = true
			}
		}
		if found {
			arr[i] = max
		} else {
			arr[i] = -1
		}
	}
	return arr
}

//https://leetcode.com/explore/learn/card/fun-with-arrays/511/in-place-operations/3259/
