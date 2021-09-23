package foo

func reverseString(s []byte) {
	if len(s) <= 1 {
		return
	}

	s[0], s[len(s)-1] = s[len(s)-1], s[0]
	reverseString(s[1 : len(s)-1])
}

//https://leetcode.com/explore/learn/card/fun-with-arrays/526/deleting-items-from-an-array/3248/
