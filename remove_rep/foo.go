package removerep

func removeDuplicates(nums []int) int {
	var i int

	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}

	if len(nums) == 0 {
		return 0
	}
	return i + 1
}

//https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2gy9m/
