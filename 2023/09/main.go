package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

// input splitted on \n
var parsed_input []string

func init() {
	// do this in init (not main) so test file has same input
	parsed_input = strings.Split(strings.TrimRight(input, "\n"), "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		games := parsed_input
		result := part1(games)
		fmt.Printf("Output: %d\n", result)
	} else if part == 2 {
		games := parsed_input
		result := part2(games)
		fmt.Printf("Output: %d\n", result)
	}
}

func part(input []string, part int) int {
	total_1 := 0
	total_2 := 0

	/*
		  idx -> [[initial_array], difference_arrays]
			0 -> [[1 2 3 4], ...]
	*/
	b := make(map[int][][]int, len(input))

	for idx, line := range input {
		ns := strings.Fields(line)
		b[idx] = make([][]int, 0)
		arr := []int{}
		for _, c := range ns {
			n, _ := strconv.Atoi(string(c))
			arr = append(arr, n)
		}
		b[idx] = append(b[idx], arr)
	}

	// At each prediction
	for i := 0; i < len(b); i++ {

		// Compute new array from the difference between elements
		for idx := 0; idx < len(b[i]); idx++ {

			arr := []int{}

			all_zeros := true
			for j := 0; j < len(b[i][idx])-1; j++ {
				// Compute difference and store
				arr = append(arr, b[i][idx][j+1]-b[i][idx][j])

				// want to reach and all zeros array
				if arr[j] != 0 {
					all_zeros = false
				}
			}

			// Append new array containing the differences
			b[i] = append(b[i], arr)
			if all_zeros {
				break
			}
		}

		// Append zero to last array
		b[i][len(b[i])-1] = append(b[i][len(b[i])-1], 0)
		// Preppend zero to last array
		b[i][len(b[i])-1] = append([]int{0}, b[i][len(b[i])-1]...)

		// Bottom up approach
		for j := len(b[i]) - 1 - 1; j >= 0; j-- {

			last_element := b[i][j][len(b[i][j])-1]
			last_element_prev_arr := b[i][j+1][len(b[i][j+1])-1]

			//append
			b[i][j] = append(b[i][j], last_element+last_element_prev_arr)

			first_element := b[i][j][0]
			first_element_below_arr := b[i][j+1][0]

			// prepend
			b[i][j] = append([]int{first_element - first_element_below_arr}, b[i][j]...)
		}

		total_1 += b[i][0][len(b[i][0])-1]
		total_2 += b[i][0][0]
	}

	if part == 1 {
		return total_1
	} else {
		return total_2
	}
}

func part1(input []string) int {
	return part(input, 1)
}

func part2(input []string) int {
	return part(input, 2)
}