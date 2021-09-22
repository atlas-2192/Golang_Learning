package foo

func findDisappearedNumbers(nums []int) []int {
	max := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := max[nums[i]]; !ok {
			max[nums[i]] = nums[i]
		}
	}

	ret := []int{}

	for i := 1; i <= len(nums); i++ {
		if _, ok := max[i]; !ok {
			ret = append(ret, i)
		}
	}
	return ret
}

//https://leetcode.com/explore/learn/card/fun-with-arrays/526/deleting-items-from-an-array/3248/
