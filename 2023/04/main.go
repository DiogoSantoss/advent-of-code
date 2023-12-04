package main

import (
	_ "embed"
	"flag"
	"fmt"
	"slices"
	"strings"
)

//go:embed input.txt
var input string
// input splitted on \n
var parsed_input []string

func init() {
	// do this in init (not main) so test file has same input
	parsed_input = strings.Split(strings.TrimRight(input, "\n"),"\n")
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

	for _, card := range input {

		card_total := 0
		numbers := strings.Split(card, ": ")[1]
		numbers_split := strings.Split(numbers, "|")

		winning_numbers := strings.Fields(numbers_split[0])
		have_numbers := strings.Fields(numbers_split[1])

		for _, number := range have_numbers {

			if slices.Contains(winning_numbers, number) {

				if card_total == 0 {
					card_total = 1
				} else {
					card_total *= 2
				}
			}
		}
		total += card_total
	}

	return total
}

func part2(input []string) int {
	total := 0

	// card id -> copies
	cards_copies := make(map[int]int, len(input))
	for i := 1; i <= len(input); i++ {
		cards_copies[i] = 1
	}

	for i := 1; i <= len(input); i++ {
		card := input[i-1]

		card_copy_idx := i

		numbers := strings.Split(card, ": ")[1]
		numbers_split := strings.Split(numbers, "|")

		winning_numbers := strings.Fields(numbers_split[0])
		have_numbers := strings.Fields(numbers_split[1])

		for _, number := range have_numbers {
			if slices.Contains(winning_numbers, number) {
				card_copy_idx++
				// for every match will have one more copy of the next card times the number 
				// of copies of the current card, therefore 1*copis_current = cards_copies[i]
				cards_copies[card_copy_idx] += cards_copies[i]
			}
		}
	}

	for _, card_amount := range cards_copies {
		total += card_amount
	}

	return total
}
