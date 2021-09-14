package foo

func findMaxConsecutiveOnes(nums []int) int {
	length := len(nums)
	max := 0
	for i := 0; i < length; i++ {
		if nums[i] == 1 {
			count := 1
			for j := i + 1; j < length; j++ {
				if nums[i] == 1 && nums[i] == nums[j] {
					count = count + 1
				} else {
					break
				}
			}
			if count > max {
				max = count
			}
		}
	}

	if length == 0 {
		return 0
	}

	if length == 1 {
		return nums[0]
	}
	return max
}
