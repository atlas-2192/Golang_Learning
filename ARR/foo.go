package foo

func duplicateZeros(arr []int) {
	var ret []int
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			ret = append(ret, 0)
			ret = append(ret, 0)
		} else {
			ret = append(ret, arr[i])
		}
		if len(ret) >= len(arr) {
			break
		}
	}
	copy(arr, ret)
}

//https://leetcode.com/explore/learn/card/fun-with-arrays/525/inserting-items-into-an-array/3245/
