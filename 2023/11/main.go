package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
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
		result := part2(games, 1000000)
		fmt.Printf("Output: %d\n", result)
	}
}

type galaxy struct {
	id    int
	row   int
	col   int
	dists map[int]int
}

func part1(input []string) int {
	total := 0

	// create new grid
	new_grid := make([]string, 0)
	for _, line := range input {
		new_grid = append(new_grid, line)
		if !strings.Contains(line, "#") {
			new_grid = append(new_grid, line)
		}
	}
	added_cols := 0
	for i := 0; i < len(input[0]); i++ {
		empty_col := true
		for j := 0; j < len(input); j++ {
			if input[j][i] == '#' {
				empty_col = false
				break
			}
		}
		if empty_col {
			for k := 0; k < len(new_grid); k++ {
				new_grid[k] = new_grid[k][:i+added_cols] + "." + new_grid[k][i+added_cols:]
			}
			added_cols++
		}
	}

	galaxy_count := 0
	galaxies := make([]galaxy, 0)
	for r, line := range new_grid {
		for c, char := range line {
			if char == '#' {
				galaxies = append(galaxies, galaxy{
					id:    galaxy_count,
					row:   r,
					col:   c,
					dists: make(map[int]int),
				})
				galaxy_count++
			}
		}
	}

	for _, galaxy := range galaxies {
		for _, other_galaxy := range galaxies {
			if galaxy.id != other_galaxy.id {
				total += int(math.Abs(float64(galaxy.row)-float64(other_galaxy.row)) + math.Abs(float64(galaxy.col)-float64(other_galaxy.col)))
			}
		}
	}

	return total / 2
}

func part2(input []string, distance_b int) int {
	total := 0

	galaxy_count := 0
	galaxies := make([]galaxy, 0)
	for r, line := range input {
		for c, char := range line {
			if char == '#' {
				galaxies = append(galaxies, galaxy{
					id:    galaxy_count,
					row:   r,
					col:   c,
					dists: make(map[int]int),
				})
				galaxy_count++
			}
		}
	}

	empty_rows := make([]int, 0)
	for r, line := range input {
		if !strings.Contains(line, "#") {
			empty_rows = append(empty_rows, r)
		}
	}

	empty_cols := make([]int, 0)
	for i := 0; i < len(input[0]); i++ {
		empty_col := true
		for j := 0; j < len(input); j++ {
			if input[j][i] == '#' {
				empty_col = false
				break
			}
		}
		if empty_col {
			empty_cols = append(empty_cols, i)
		}
	}

	for _, galaxy := range galaxies {
		for _, other_galaxy := range galaxies {
			if galaxy.id != other_galaxy.id && galaxy.dists[other_galaxy.id] == 0 {

				galaxy_col := galaxy.col
				galaxy_row := galaxy.row

				new_galaxy_col := galaxy_col
				new_galaxy_row := galaxy_row

				other_galaxy_col := other_galaxy.col
				other_galaxy_row := other_galaxy.row

				new_other_galaxy_col := other_galaxy_col
				new_other_galaxy_row := other_galaxy.row

				for _, col := range empty_cols {
					if (galaxy_col < col && other_galaxy_col > col) {
						new_other_galaxy_col = new_other_galaxy_col - 1 + distance_b
					} else if other_galaxy_col < col && galaxy_col > col {
						new_galaxy_col = new_galaxy_col - 1 + distance_b
					}
				}

				for _, row := range empty_rows {
					if (galaxy_row < row && other_galaxy_row > row) {
						new_other_galaxy_row = new_other_galaxy_row - 1 + distance_b
					} else if (other_galaxy_row < row && galaxy_row > row) {
						new_galaxy_row = new_galaxy_row -1 + distance_b
					}
				}

				dist := int(math.Abs(float64(new_galaxy_row)-float64(new_other_galaxy_row)) + math.Abs(float64(new_galaxy_col)-float64(new_other_galaxy_col)))
				total += dist
				galaxy.dists[other_galaxy.id] = dist
				other_galaxy.dists[galaxy.id] = dist
			}
		}
	}

	return total
}
