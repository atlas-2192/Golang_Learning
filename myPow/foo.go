package foo

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n > 0 {
		return myPow(x, n-1) * x
	}
	return 1 / (myPow(x, (0-n)-1) * x)
}

//https://leetcode.com/explore/learn/card/fun-with-arrays/527/searching-for-items-in-an-array/3250/
