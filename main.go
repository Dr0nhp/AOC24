package main

import (
	"AOC24/reader"
	/*	p1d1 "AOC24/solutions/d1/p1"
		p1d2 "AOC24/solutions/d1/p2"*/
	p2d1 "AOC24/solutions/d2/p1"
	/*	p2d2 "AOC24/solutions/d2/p2"*/
	"fmt"
)

func main() {
	/*	arr_l, arr_r, _ := reader.ReadData1("data/inp_d1.txt")*/
	arr, _ := reader.ReadData2("data/inp_d2.txt")

	/*	solution1 := p1d1.Solution(arr_l, arr_r)
		solution2 := p1d2.Solution(arr_l, arr_r)*/
	solution3 := p2d1.Solution(arr)

	/*	fmt.Printf("Solution 1: %d\n", solution1)
		fmt.Printf("Solution 2: %d\n", solution2)*/
	fmt.Printf("Solution 3: %v\n", solution3)
}
