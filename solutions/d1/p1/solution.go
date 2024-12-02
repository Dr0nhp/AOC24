package p1

import "sort"

func Solution(arr_l, arr_r []int) int {
	sort.Ints(arr_l)
	sort.Ints(arr_r)

	var solution1 int
	for i := range arr_l {
		if arr_l[i] < arr_r[i] {
			solution1 += (arr_r[i] - arr_l[i])
		} else {
			solution1 += (arr_l[i] - arr_r[i])
		}
	}
	return solution1
}
