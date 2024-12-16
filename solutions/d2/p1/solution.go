package p1

import (
	"fmt"
	"math"
)

func Solution(arr [][]int) int {
	var save int

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i])-1; j++ {
			compL := float64(arr[i][j])
			compR := float64(arr[i][j+1])

			if math.Abs(compL-compR) < 3 && math.Abs(compL-compR) != 0 {
				fmt.Print(compL, compR, "\n")

				save++
			} else {
				break
			}
		}
	}

	return save
}

/*

The levels are either all increasing or all decreasing.
Any two adjacent levels differ by at least one and at most three.
In the example above, the reports can be found safe or unsafe by checking those rules:

7 6 4 2 1: Safe because the levels are all decreasing by 1 or 2.
1 2 7 8 9: Unsafe because 2 7 is an increase of 5.
9 7 6 2 1: Unsafe because 6 2 is a decrease of 4.
1 3 2 4 5: Unsafe because 1 3 is increasing but 3 2 is decreasing.
8 6 4 4 1: Unsafe because 4 4 is neither an increase or a decrease.
1 3 6 7 9: Safe because the levels are all increasing by 1, 2, or 3.*/
