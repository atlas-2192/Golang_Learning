package foo

func duplicateZeros(arr []int) {
	var a []int
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 && i < len(arr)-1 && len(a) == 0 {
			a = append(a, arr[i+1])
			arr[i+1] = 0
			i++
		} else if len(a) != 0 {
			a = append(a, arr[i])
			if a[0] == 0 {
				a = append(a, arr[i])
				a = append(a, arr[i+1])
				arr[i+1] = 0
			}
			arr[i], a = a[0], a[1:]
		}
	}
}

//https://leetcode.com/explore/learn/card/fun-with-arrays/525/inserting-items-into-an-array/3245/
