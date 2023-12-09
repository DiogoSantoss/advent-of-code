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

type instruction struct {
	source string
	left   string
	right  string
}

func part1(input []string) int {
	moves := input[0]
	input = input[2:]

	instrucions := make(map[string]instruction, 0)
	for _, line := range input {

		split := strings.Split(line, " = ")
		source := split[0]

		left_right := strings.Split(split[1], ", ")

		instrucions[source] = instruction{
			source: source,
			left:   left_right[0][1:],
			right:  left_right[1][:len(left_right[1])-1],
		}
	}

	steps := 0
	move_idx := 0
	current_node := "AAA"
	for current_node != "ZZZ" {
		steps++
		move := moves[move_idx]
		move_idx = (move_idx + 1) % len(moves)
		fmt.Printf("%s (%c) ->",current_node, move)
		if move == 'R' {
			current_node = instrucions[current_node].right
		} else {
			current_node = instrucions[current_node].left
		}
	}

	return steps
}

func part2(input []string) int {
	moves := input[0]
	input = input[2:]

	nodes := make([]string, 0)
	instrucions := make(map[string]instruction, 0)
	for _, line := range input {

		split := strings.Split(line, " = ")
		source := split[0]

		if source[2] == 'A' {
			nodes = append(nodes, source)
		}

		left_right := strings.Split(split[1], ", ")

		instrucions[source] = instruction{
			source: source,
			left:   left_right[0][1:],
			right:  left_right[1][:len(left_right[1])-1],
		}
	}

	nodes_pattern := make([]int, 0)
	for _, node := range nodes {
		steps := 0
		move_idx := 0
		current_node := node
		for current_node[2] != 'Z' {
			steps++
			move := moves[move_idx]
			move_idx = (move_idx + 1) % len(moves)
			if move == 'R' {
				current_node = instrucions[current_node].right
			} else {
				current_node = instrucions[current_node].left
			}
		}
		nodes_pattern = append(nodes_pattern, steps)
	}

	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	// least common multiplier
	// simple case: 4 and 3 -> 4x3
	// however: 10 and 5 -> 5
	// therefore need gcd for this case (10*5/5)
	lcm := func(v []int) int {
		total := 1
		for _, s := range v {
			total = (s * total) / gcd(s, total)
		}
		return total
	}

	return lcm(nodes_pattern)
}

	// AAA -> 4
	// AAA -> AAB -> AAC -> AAZ
	// meaning it takes 3 step to go from AAA to AAZ

	// AAB -> 3
	// AAB -> AAC -> AAZ
	// meaning it takes 3 step to go from AAB to AAZ

	// AAC -> 2
	// AAC -> AAZ
	// meaning it takes 2 step to go from AAC to AAZ

	// AAK -> 7

	// when do they meet ?
	//  if its divider they meet c=max(a,b) d=min(a,b)
	//  else they meet after a*b steps

