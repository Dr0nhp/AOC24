package reader

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadData1(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var arr_l []int
	var arr_r []int

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		replacer := strings.NewReplacer(",", "")
		line = replacer.Replace(line)
		numbers := strings.Fields(line)

		for i, numStr := range numbers {
			number, err := strconv.Atoi(numStr)
			if err != nil {
				continue
			}

			if i%2 == 0 {
				arr_l = append(arr_l, number)
			} else {
				arr_r = append(arr_r, number)
			}
		}
	}
	return arr_l, arr_r, err
}

func ReadData2(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var result [][]int

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		numbers := strings.Fields(line)

		var intNumbers []int
		for _, numStr := range numbers {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, err
			}
			intNumbers = append(intNumbers, num)
		}
		result = append(result, intNumbers)
	}

	return result, nil
}
