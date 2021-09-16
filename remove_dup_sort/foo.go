package foo

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	i, j := 0, 1
	for j < len(nums) {
		if nums[i] == nums[j] {
			j++
		} else {
			nums[i+1] = nums[j]
			j++
			i++
		}
	}
	return i + 1
}

//https://leetcode.com/explore/learn/card/fun-with-arrays/526/deleting-items-from-an-array/3248/
