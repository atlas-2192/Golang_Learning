package foo

func checkIfExist(arr []int) bool {

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == 2*arr[j] || 2*arr[i] == arr[j] {
				return true
			}
		}
	}
	return false
}

//https://leetcode.com/explore/learn/card/fun-with-arrays/527/searching-for-items-in-an-array/3250/
