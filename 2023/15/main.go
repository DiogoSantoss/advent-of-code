package main

import (
	_ "embed"
	"flag"
	"fmt"
	"slices"
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

func hash(word string) int {
	total := 0
	for _, char := range word {
		total += int(char)
		total *= 17
		total = total % 256
	}
	return total
}

func part1(input []string) int {

	total := 0
	for _, step := range strings.Split(input[0], ",") {
		total += hash(step)
	}

	return total
}

type lens struct {
	label string
	fl    int
}

func part2(input []string) int {

	boxes := make(map[int][]lens, 256)

	for _, step := range strings.Split(input[0], ",") {
		step_split := strings.FieldsFunc(step, func(r rune) bool {
			return r == '-' || r == '='
		})

		label := step_split[0]
		box_id := hash(label)
		op := step[len(step_split[0])]

		if op == '-' {
			boxes[box_id] = slices.DeleteFunc(boxes[box_id], func(l lens) bool {
				return l.label == label
			})
		} else {
			fl, _ := strconv.Atoi(step_split[1])
			new_lens := lens{
				label: label,
				fl:    fl,
			}
			idx := slices.IndexFunc(boxes[box_id], func(l lens) bool {
				return l.label == label
			})

			if idx == -1 {
				boxes[box_id] = append(boxes[box_id], new_lens)
			} else {
				boxes[box_id][idx] = new_lens
			}
		}
	}

	total := 0
	for i := 0; i < 256; i++ {
		for j, lens := range boxes[i] {
			total += (j + 1) * (i + 1) * lens.fl
		}
	}

	return total
}
