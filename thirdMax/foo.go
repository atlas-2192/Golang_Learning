package foo

import "sort"

func thirdMax(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	max := make(map[int]int)
	keys := []int{}
	for i := 0; i < len(nums); i++ {
		if _, ok := max[nums[i]]; !ok {
			max[nums[i]] = nums[i]
			keys = append(keys, nums[i])
		}
	}
	sort.Ints(keys)

	if len(max) < 3 {
		return keys[len(keys)-1]
	}
	keys = keys[len(keys)-3 : len(keys)-1]
	return keys[0]
}

//https://leetcode.com/explore/learn/card/fun-with-arrays/526/deleting-items-from-an-array/3248/
