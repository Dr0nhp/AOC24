package main

import (
	"AOC24/reader"
	p1d1 "AOC24/solutions/d1/p1"
	p1d2 "AOC24/solutions/d1/p2"
	p2d1 "AOC24/solutions/d2/p1"
	p2d2 "AOC24/solutions/d2/p2"
	"fmt"
	"log"
)

func main() {
	arr_l, arr_r, _, err := reader.ReadData("data/inp_d1.txt")
	if err != nil {
		log.Fatalf("Error while reading the file: %v", err)
	}

	solution1 := p1d1.Solution(arr_l, arr_r)
	solution2 := p1d2.Solution(arr_l, arr_r)

	fmt.Printf("Solution 1: %d\n", solution1)
	fmt.Printf("Solution 2: %d\n", solution2)
}
