package foo

func moveZeroes(nums []int) {
	if len(nums) <= 1 {
		return
	}
	i, j := 0, 0
	for j < len(nums) {
		if nums[j] == 0 {
			j++
		} else {
			nums[i], nums[j] = nums[j], nums[i]
			j++
			i++
		}
	}
	return
}

//https://leetcode.com/explore/learn/card/fun-with-arrays/526/deleting-items-from-an-array/3248/
