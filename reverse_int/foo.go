package reverseint

import (
	"strconv"
)

func reverse(x int) int {
	t := strconv.Itoa(x)
	var o []byte
	var negtive bool
	if t[0] == '-' {
		negtive = true
		t = t[1:]
	}
	// fmt.Println(string(t))
	for i := 0; i < len(t); i++ {
		o = append(o, t[len(t)-i-1])
	}
	if negtive {
		// p := o
		o = append([]byte{'-'}, o...)
		// fmt.Println(string(t))
	}

	// fmt.Println(string(o))
	ret, _ := strconv.Atoi(string(o))
	if ret > 2147483647 || ret < -2147483648 {
		return 0
	}
	return ret
}
