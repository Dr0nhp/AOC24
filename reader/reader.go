package reader

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadData(filename string) ([]int, []int, []string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var arr_l []int
	var arr_r []int
	var lines []string

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		lines = append(lines, line)

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
	return arr_l, arr_r, lines, err
}
