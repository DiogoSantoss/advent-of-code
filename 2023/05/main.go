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

type map_range struct {
	dest_range_start   int
	source_range_start int
	lenght             int
}

func part1(input []string) int {
	seeds := strings.Fields(strings.Split(input[0], ": ")[1])

	maps := make(map[int][]map_range)

	map_index := -1

	reading_map := false
	for _, line := range input {

		if len(line) == 0 {
			reading_map = false
		}
		if strings.Contains(line, "map") {
			reading_map = true
			map_index++
			continue
		}
		if !reading_map {
			continue
		}

		numbers := strings.Fields(line)
		n1, _ := strconv.Atoi(numbers[0])
		n2, _ := strconv.Atoi(numbers[1])
		n3, _ := strconv.Atoi(numbers[2])
		maps[map_index] = append(maps[map_index], map_range{
			dest_range_start:   n1,
			source_range_start: n2,
			lenght:             n3,
		})
	}

	min_seed_n := 100000000000000000
	// For every seed
	for _, seed := range seeds {

		seed_n, _ := strconv.Atoi(seed)

		// Go in each map
		for i := 0; i < len(maps); i++ {

			value := maps[i]
			found_mapping := false
			// search for a valid map_range
			for _, mr := range value {

				// already translated
				if found_mapping {
					break
				}

				// if within range, translate by range
				if (seed_n >= mr.source_range_start) && (seed_n <= mr.source_range_start+mr.lenght) {
					found_mapping = true
					seed_n = (seed_n - mr.source_range_start) + mr.dest_range_start
				}
			}
		}

		if seed_n < min_seed_n {
			min_seed_n = seed_n
		}
	}

	return min_seed_n
}

func part2(input []string) int {

	seeds_initial := strings.Fields(strings.Split(input[0], ": ")[1])
	seeds := make([]int, 0)

	for i := 0; i < len(seeds_initial)/2; i += 2 {
		initial_value, _ := strconv.Atoi(seeds_initial[i])
		seed_range, _ := strconv.Atoi(seeds_initial[i+1])
		for s := initial_value; s < initial_value+seed_range; s++ {
			seeds = append(seeds, s)
		}
	}

	maps := make(map[int][]map_range)

	map_index := -1

	reading_map := false
	for _, line := range input {

		if len(line) == 0 {
			reading_map = false
		}
		if strings.Contains(line, "map") {
			reading_map = true
			map_index++
			continue
		}
		if !reading_map {
			continue
		}

		numbers := strings.Fields(line)
		n1, _ := strconv.Atoi(numbers[0])
		n2, _ := strconv.Atoi(numbers[1])
		n3, _ := strconv.Atoi(numbers[2])
		maps[map_index] = append(maps[map_index], map_range{
			dest_range_start:   n1,
			source_range_start: n2,
			lenght:             n3,
		})
	}

	min_seed_n := 100000000000000000
	// For every seed
	for _, seed_n := range seeds {

		// Go in each map
		for i := 0; i < len(maps); i++ {

			value := maps[i]
			found_mapping := false
			// search for a valid map_range
			for _, mr := range value {

				// already translated
				if found_mapping {
					break
				}

				// if within range, translate by range
				if (seed_n >= mr.source_range_start) && (seed_n <= mr.source_range_start+mr.lenght) {
					found_mapping = true
					seed_n = (seed_n - mr.source_range_start) + mr.dest_range_start
				}
			}
		}

		if seed_n < min_seed_n {
			min_seed_n = seed_n
		}
	}

	return min_seed_n
}
