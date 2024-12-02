package p2

func Solution(arr_l, arr_r []int) int {
	var solution2 int
	for _, el := range arr_l {
		var count int
		for _, innerEl := range arr_r {
			if innerEl == el {
				count++
			}
		}
		solution2 += count * el
	}
	return solution2
}
