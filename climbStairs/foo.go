package foo

var cache [46]int

func climbStairs(n int) int {
	if n == 0 {
		cache[0] = 1
		return 1
	} else if n < 3 {
		cache[n] = n
		return n
	}

	if cache[n] != 0 {
		return cache[n]
	}

	cache[n] = climbStairs(n-2) + climbStairs(n-1)
	return cache[n]
}

//https://leetcode.com/explore/featured/card/recursion-i/255/recursion-memoization/1662/
