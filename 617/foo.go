package main

// "time"

func maxCount(m int, n int, ops [][]int) int {
	var base [40000][40000]int

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
