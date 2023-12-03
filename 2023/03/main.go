package main

import (
	_ "embed"
	"flag"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		games := parseInput(input)
		result := part1(games)
		fmt.Printf("Output: %d\n", result)
	} else if part == 2 {
		games := parseInput(input)
		result := part2(games)
		fmt.Printf("Output: %d\n", result)
	}
}

func is_symbol(char rune) bool {
	return slices.Contains([]rune{'*', '+', '-', '#', '@', '/', '%', '=', '$', '&'}, char)
}

func has_adjacent_symbol(input []string, x int, y int) bool {
	for x_i := x - 1; x_i < x+2; x_i++ {
		for y_i := y - 1; y_i < y+2; y_i++ {
			if (x_i >= 0) &&
				(y_i >= 0) &&
				(y_i < len(input)) &&
				(x_i < len(input[y_i])) &&
				is_symbol(rune(input[y_i][x_i])) {
				return true
			}
		}
	}

	return false
}

func part1(input []string) int {
	total := 0

	for idx_l, line := range input {

		number_chars := []rune{}

		for idx_c, char := range line {

			// build number
			if unicode.IsNumber(char) {
				number_chars = append(number_chars, char)
			}

			// number ended with symbol or '.'
			if ((char == '.') || (is_symbol(char))) && (len(number_chars) != 0) {

				initial_index := idx_c - len(number_chars)
				final_index := idx_c - 1
				height := idx_l

				for i := initial_index; i <= final_index; i++ {
					if has_adjacent_symbol(input, i, height) {
						number, _ := strconv.Atoi(string(number_chars))
						total += number
						break
					}
				}
				// clear
				number_chars = []rune{}
			}
		}

		// line ended
		if len(number_chars) != 0 {

			initial_index := len(line) - 1 - len(number_chars)
			final_index := len(line) - 1 - 1
			height := idx_l

			for i := initial_index; i <= final_index; i++ {
				if has_adjacent_symbol(input, i, height) {
					number, _ := strconv.Atoi(string(number_chars))
					total += number
					break
				}
			}
		}
	}

	return total
}

type coord struct {
	X int
	Y int
}

func has_adjacent_symbol_2(gears map[coord][]string, input []string, x int, y int, number []rune) bool {
	for x_i := x - 1; x_i < x+2; x_i++ {
		for y_i := y - 1; y_i < y+2; y_i++ {
			if (x_i >= 0) &&
				(y_i >= 0) &&
				(y_i < len(input)) &&
				(x_i < len(input[y_i])) &&
				is_symbol(rune(input[y_i][x_i])) {
				if input[y_i][x_i] == '*' {
					gears[coord{X: x_i, Y: y_i}] = append(gears[coord{X: x_i, Y: y_i}], string(number))
					return true
				}
			}
		}
	}
	return false
}

func part2(input []string) int {
	total := 0

	gears := make(map[coord][]string)

	for idx_l, line := range input {

		number_chars := []rune{}

		for idx_c, char := range line {

			// build number
			if unicode.IsNumber(char) {
				number_chars = append(number_chars, char)

			}

			// number ended with symbol or '.'
			if ((char == '.') || (is_symbol(char))) && (len(number_chars) != 0) {

				initial_index := idx_c - len(number_chars)
				final_index := idx_c - 1
				height := idx_l

				for i := initial_index; i <= final_index; i++ {
					if has_adjacent_symbol_2(gears, input, i, height, number_chars) {
						break
					}
				}
				// clear
				number_chars = []rune{}
			}
		}

		// line ended
		if len(number_chars) != 0 {

			initial_index := len(line) - 1 - len(number_chars)
			final_index := len(line) - 1 - 1
			height := idx_l

			for i := initial_index; i <= final_index; i++ {
				if has_adjacent_symbol_2(gears, input, i, height, number_chars) {
					break
				}
			}
		}
	}

	for _, value := range gears {
		if len(value) == 2 {
			a, _ := strconv.Atoi(value[0])
			b, _ := strconv.Atoi(value[1])
			total += a * b
		}
	}

	return total
}
