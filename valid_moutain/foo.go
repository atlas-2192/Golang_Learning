package foo

func validMountainArray(arr []int) bool {
	up, down := false, false
	i, j := 0, 1
	for j < len(arr) {
		if arr[i] < arr[j] && down == false {
			up = true
		} else if arr[i] > arr[j] && up == true {
			down = true
		} else {
			return false
		}
		i++
		j++
	}
	return up && down
}

//https://leetcode.com/explore/learn/card/fun-with-arrays/527/searching-for-items-in-an-array/3250/
