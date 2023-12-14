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

func tilt_north(R int, C int, input []string) {
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if input[i][j] == 'O' {

				if i == 0 {
					continue
				}

				var f int
				found_wall := false
				for f = i - 1; f > 0; f-- {
					if input[f][j] == '#' || input[f][j] == 'O' {
						found_wall = true
						break
					}
				}

				// either hit a wall or hit the (ocupied) end of platform
				if (f == 0 && (input[f][j] == '#' || input[f][j] == 'O')) || found_wall {
					f++
				}

				if input[f][j] == '.' {
					input[i] = input[i][:j] + "." + input[i][j+1:]
					input[f] = input[f][:j] + "O" + input[f][j+1:]
				}
			}
		}
	}
}

func tilt_south(R int, C int, input []string) {
	for i := R - 1; i >= 0; i-- {
		for j := 0; j < C; j++ {
			if input[i][j] == 'O' {

				if i == R-1 {
					continue
				}

				var f int
				found_wall := false
				for f = i + 1; f < R-1; f++ {
					if input[f][j] == '#' || input[f][j] == 'O' {
						found_wall = true
						break
					}
				}

				// either hit a wall or hit the (ocupied) end of platform
				if (f == R-1 && (input[f][j] == '#' || input[f][j] == 'O')) || found_wall {
					f--
				}
				if input[f][j] == '.' {
					input[i] = input[i][:j] + "." + input[i][j+1:]
					input[f] = input[f][:j] + "O" + input[f][j+1:]
				}
			}
		}
	}
}

func tilt_west(R int, C int, input []string) {
	for j := 0; j < C; j++ {
		for i := 0; i < R; i++ {
			if input[i][j] == 'O' {

				if j == 0 {
					continue
				}

				var f int
				found_wall := false
				for f = j - 1; f > 0; f-- {
					if input[i][f] == '#' || input[i][f] == 'O' {
						found_wall = true
						break
					}
				}
				// either hit a wall or hit the (ocupied) end of platform
				if (f == 0 && (input[i][f] == '#' || input[i][f] == 'O')) || found_wall {
					f++
				}
				if input[i][f] == '.' {
					input[i] = input[i][:j] + "." + input[i][j+1:]
					input[i] = input[i][:f] + "O" + input[i][f+1:]
				}
			}
		}
	}
}

func tilt_east(R int, C int, input []string) {
	for j := C - 1; j >= 0; j-- {
		for i := 0; i < R; i++ {
			if input[i][j] == 'O' {

				if j == C-1 {
					continue
				}

				var f int
				found_wall := false
				for f = j + 1; f < C-1; f++ {
					if input[i][f] == '#' || input[i][f] == 'O' {
						found_wall = true
						break
					}
				}
				// either hit a wall or hit the (ocupied) end of platform
				if (f == C-1 && (input[i][f] == '#' || input[i][f] == 'O')) || found_wall {
					f--
				}
				if input[i][f] == '.' {
					input[i] = input[i][:j] + "." + input[i][j+1:]
					input[i] = input[i][:f] + "O" + input[i][f+1:]
				}
			}
		}
	}
}

type cache_value struct {
	res []string
	idx int
}

func build_key(a []string) string {
	return strings.Join(a, "")
}

func compute_value(b []string) int {
	total := 0

	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[0]); j++ {
			if b[i][j] == 'O' {
				total += len(b) - i
			}
		}
	}

	return total
}

func part1(input []string) int {
	R := len(input)
	C := len(input[0])

	tilt_north(R, C, input)

	return compute_value(input)
}

func part2(input []string) int {
	R := len(input)
	C := len(input[0])

	og_input := slices.Clone(input)

	cache := make(map[string]cache_value)
	for i := 0; i < 1000000000; i++ {

		key := build_key(input)

		if value, ok := cache[key]; ok {

			// first time cycle appeared
			cycle_start := value.idx
			// how long the cycle is
			cycle_len := i - cycle_start
			// e.g. i = 3 was saved to cache and then i = 7 hit cache
			// means cycle starts at 3 -> 4 -> 5 -> 6, then 7 is equal to 3
			// cycle len is 7-3 = 4
			// since loop might not start right away must discount the tail
			// effective_cnt tells you when in the loop do the itterations end
			// e.g. might end after 2 steps in the loop
			effective_cnt := (1000000000 - cycle_start) % cycle_len
			
			// from the original input do the tail plus partial loop
			// e.g.  (1->2->)(initial)(3->4->5)(partial loop)
			input = og_input
			for i := 0; i < cycle_start+effective_cnt; i++ {
				tilt_north(R, C, input)
				tilt_west(R, C, input)
				tilt_south(R, C, input)
				tilt_east(R, C, input)
			}

			return compute_value(input)
		}

		tilt_north(R, C, input)
		tilt_west(R, C, input)
		tilt_south(R, C, input)
		tilt_east(R, C, input)

		cache[key] = cache_value{
			res: input,
			idx: i,
		}
	}
	panic("shouldn't be here")
}
