package foo

func removeElement(nums []int, val int) int {
	out := nums[:len(nums)]
	count := 0
	for _, v := range nums {
		if v != val {
			out[count] = v
			count++
		}
	}
	return count
}

//https://leetcode.com/explore/learn/card/fun-with-arrays/526/deleting-items-from-an-array/3247/
