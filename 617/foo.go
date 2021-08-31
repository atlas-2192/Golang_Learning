package main

// "time"

func maxCountWrong(m int, n int, ops [][]int) int {
	var base [400][400]int

	for _, v := range ops {
		for i := 0; i < v[1]; i++ {
			for j := 0; j < v[0]; j++ {
				base[i][j]++
			}
		}
	}

	mm := make(map[int]int)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			mm[base[i][j]] = mm[base[i][j]] + 1
			// fmt.Printf("%d ", base[i][j])
		}
		// fmt.Printf("\n")
	}

	// fmt.Printf("map:")
	// var biggest int
	// for k, v := range mm {
	// 	if v > biggest {
	// 		biggest = v
	// 	}
	// 	fmt.Printf("(%d,%d)", k, v)
	// }

	return mm[base[0][0]]
}

func maxCount(m int, n int, ops [][]int) int {
	minM, minN := m, n
	for _, op := range ops {
		minM = min(minM, op[0])
		minN = min(minN, op[1])
	}
	return minM * minN
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
