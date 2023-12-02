package main

import (
	_ "embed"
	"flag"
	"fmt"
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
		calibration_values := parseInput(input)
		result := part1(calibration_values)
		fmt.Printf("Output: %d\n", result)
	} else if part == 2 {
		calibration_values := parseInput(input)
		result := part2(calibration_values)
		fmt.Printf("Output: %d\n", result)
	}
}

func part1(input []string) int {

	total := 0
	for _, value := range input {
		l, r := 0, len(value)-1
		first_digit, last_digit, final_digits := ' ', ' ', ""

		for {

			left_char, right_char := rune(value[l]), rune(value[r])

			if unicode.IsNumber(left_char) {
				first_digit = left_char
			} else {
				l++
			}

			if unicode.IsNumber(right_char) {
				last_digit = right_char
			} else {
				r--
			}

			// Pointers met
			if l >= r {
				if first_digit != ' ' {
					final_digits = string(first_digit) + string(first_digit)
				} else if last_digit != ' ' {
					final_digits = string(last_digit) + string(last_digit)
				} else {
					final_digits = string(left_char) + string(right_char)
				}

				break
			}

			// Found first and last digit
			if (first_digit != ' ') && (last_digit != ' ') {
				final_digits = string(first_digit) + string(last_digit)
				break
			}
		}

		number, _ := strconv.Atoi(final_digits)
		total += number
	}

	return total
}

func part2(input []string) int {

	numberToValue := map[string]rune{
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}

	spelled_numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	total := 0

	for _, word := range input {

		numbers := make([]rune, 0)

		for idx, c := range word {

			if unicode.IsNumber(c) {
				numbers = append(numbers, c)
				continue
			}

			for _, spelled_number := range spelled_numbers {
				if strings.HasPrefix(word[idx:], spelled_number) {
					numbers = append(numbers, numberToValue[spelled_number])
				}
			}
		}

		final_number, _ := strconv.Atoi(string(numbers[0]) + string(numbers[len(numbers)-1]))
		total += final_number
	}

	return total
}
