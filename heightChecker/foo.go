package foo

import "sort"

func heightChecker(nums []int) int {
	ori := make([]int, len(nums))
	copy(ori, nums)
	sort.Ints(nums)
	count := 0
	for i := 0; i < len(nums); i++ {
		if ori[i] != nums[i] {
			count++
		}
	}
	return count
}

//https://leetcode.com/explore/learn/card/fun-with-arrays/526/deleting-items-from-an-array/3248/
