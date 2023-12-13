package main

import (
	_ "embed"
	"flag"
	"fmt"
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

func part1(input []string) int {

	total := 0
	patterns := make([][]string, 0)

	counter := 0
	for i := 0; i < len(input); i++ {
		if len(input[i]) == 0 {
			counter++
			continue
		}

		if counter >= len(patterns) {
			patterns = append(patterns, []string{})
		}

		patterns[counter] = append(patterns[counter], input[i])
	}

	for _, p := range patterns {

		rev_left_side := []string{}
		for _, line := range p {
			rev_left_side = append(rev_left_side, string(line[0]))
		}

		// for each possible vertical line
		v := 0
		for i := 1; i < len(p[0]); i++ {
			// check if horizontal lines on the right are reflection of left
			is_reflection := true
			for j := 0; j < len(p); j++ {

				// handle reflection on the left being bigger than on the right
				if len(rev_left_side[j]) <= len(p[j][i:]) {
					if !strings.HasPrefix(p[j][i:], rev_left_side[j]) {
						is_reflection = false
					}
				} else {
					if !strings.HasPrefix(rev_left_side[j], p[j][i:]) {
						is_reflection = false
					}
				}
				// add in reverse to check as prefix
				rev_left_side[j] = string(p[j][i]) + rev_left_side[j]
			}
			if is_reflection {
				v = i
				break
			}
		}

		if v != 0 {
			total += v
			continue
		}

		rev_up_side := []string{}
		for _, c := range p[0] {
			rev_up_side = append(rev_up_side, string(c))
		}

		// for each horizontal line
		h := 0
		for i := 1; i < len(p); i++ {

			// check if vertical lines on bottom are reflection of top
			is_reflection := true
			for j := 0; j < len(p[0]); j++ {

				// build line
				p2_horizontal_line := ""
				for k := i; k < len(p); k++ {
					p2_horizontal_line += string(p[k][j])
				}

				// handle reflection on the up being bigger than on the bottom
				if len(p2_horizontal_line) >= len(rev_up_side[j]) {
					if !strings.HasPrefix(p2_horizontal_line, rev_up_side[j]) {
						is_reflection = false
					}
				} else {
					if !strings.HasPrefix(rev_up_side[j], p2_horizontal_line) {
						is_reflection = false
					}
				}

				// add in reverse to check as prefix
				rev_up_side[j] = string(p2_horizontal_line[0]) + rev_up_side[j]
			}
			if is_reflection {
				h = i
				break
			}
		}

		if h == 0 {
			panic("should have reflection")
		}

		total += 100*h
	}

	return total
}

func part2(input []string) int {

	total := 0
	patterns := make([][]string, 0)

	counter := 0
	for i := 0; i < len(input); i++ {
		if len(input[i]) == 0 {
			counter++
			continue
		}

		if counter >= len(patterns) {
			patterns = append(patterns, []string{})
		}

		patterns[counter] = append(patterns[counter], input[i])
	}

	for _, p := range patterns {

		rev_left_side := []string{}
		for _, line := range p {
			rev_left_side = append(rev_left_side, string(line[0]))
		}

		// for each possible vertical line
		v := 0
		for i := 1; i < len(p[0]); i++ {
			// check if horizontal lines on the right are reflection of left
			is_reflection := true
			differences := 0
			for j := 0; j < len(p); j++ {

				// handle reflection on the left being bigger than on the right
				if len(rev_left_side[j]) <= len(p[j][i:]) {
					if !strings.HasPrefix(p[j][i:], rev_left_side[j]) {
						is_reflection = false
						for z := 0;z<len(rev_left_side[j]);z++ {
							if rev_left_side[j][z] != p[j][i:][z] {
								differences++
							}
						}
					}
				} else {
					if !strings.HasPrefix(rev_left_side[j], p[j][i:]) {
						is_reflection = false
						for z := 0;z<len(p[j][i:]);z++ {
							if rev_left_side[j][z] != p[j][i:][z] {
								differences++
							}
						}
					}
				}
				// add in reverse to check as prefix
				rev_left_side[j] = string(p[j][i]) + rev_left_side[j]
			}
			if !is_reflection && differences == 1 {
				v = i
				break
			}
		}

		if v != 0 {
			total += v
			continue
		}

		rev_up_side := []string{}
		for _, c := range p[0] {
			rev_up_side = append(rev_up_side, string(c))
		}

		// for each horizontal line
		h := 0
		for i := 1; i < len(p); i++ {

			// check if vertical lines on bottom are reflection of top
			is_reflection := true
			differences := 0
			for j := 0; j < len(p[0]); j++ {

				p2_horizontal_line := ""
				for k := i; k < len(p); k++ {
					p2_horizontal_line += string(p[k][j])
				}

				// handle reflection on the up being bigger than on the bottom
				if len(p2_horizontal_line) > len(rev_up_side[j]) {
					if !strings.HasPrefix(p2_horizontal_line, rev_up_side[j]) {
						is_reflection = false
						for z := 0;z<len(rev_up_side[j]);z++ {
							if rev_up_side[j][z] != p2_horizontal_line[z] {
								differences++
							}
						}
					}
				} else {
					if !strings.HasPrefix(rev_up_side[j], p2_horizontal_line) {
						is_reflection = false
						for z := 0;z<len(p2_horizontal_line);z++ {
							if rev_up_side[j][:len(p2_horizontal_line)][z] != p2_horizontal_line[z] {
								differences++
							}
						}
					}
				}

				// add in reverse to check as prefix
				rev_up_side[j] = string(p2_horizontal_line[0]) + rev_up_side[j]
			}
			if !is_reflection && differences == 1 {
				h = i
				break
			}
		}

		if h == 0 {
			panic("should have reflection")
		}

		total += 100*h
	}

	return total
}
