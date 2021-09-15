package foo

func merge(nums1 []int, m int, nums2 []int, n int) {
	var out []int
	j, k := 0, 0

	for i := 0; i < m+n; i++ {
		if j >= m {
			out = append(out, nums2[k])
			k++
		} else if k >= n {
			out = append(out, nums1[j])
			j++
		} else if nums1[j] <= nums2[k] {
			out = append(out, nums1[j])
			j++
		} else {
			out = append(out, nums2[k])
			k++
		}
	}
	nums1 = nums1[:m]
	nums2 = nums2[:n]
	nums1 = append(nums1, nums2...)
	copy(nums1, out)
}

//https://leetcode.com/explore/learn/card/fun-with-arrays/525/inserting-items-into-an-array/3253/
