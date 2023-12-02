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

func part1(input []string) int {
	total := 0
	total_red_cubes, total_green_cubes, total_blue_cubes := 12, 13, 14

	for _, game := range input {

		game_sets := strings.Split(game, ": ")

		game_id, _ := strconv.Atoi(strings.Split(game_sets[0], " ")[1])
		sets := strings.Split(game_sets[1], "; ")

		correct_set := false

		for _, set := range sets {

			total_red_set, total_green_set, total_blue_set := 0, 0, 0
			cubes := strings.Split(set, ", ")

			for _, cube := range cubes {

				qnt_color := strings.Split(cube, " ")
				qnt, _ := strconv.Atoi(qnt_color[0])

				switch qnt_color[1] {
				case "red":
					total_red_set += qnt
					break
				case "green":
					total_green_set += qnt
					break
				case "blue":
					total_blue_set += qnt
					break
				}
			}

			// change to correct once, if ever fails, break
			if !((total_red_set > total_red_cubes) || (total_green_set > total_green_cubes) || (total_blue_set > total_blue_cubes)) {
				correct_set = true
			} else {
				correct_set = false
				break
			}
		}

		if correct_set {
			total += game_id
		}
	}

	return total
}

func part2(input []string) int {
	total := 0

	for _, game := range input {

		game_sets := strings.Split(game, ": ")
		sets := strings.Split(game_sets[1], "; ")

		max_red, max_green, max_blue := 0, 0, 0

		for _, set := range sets {

			total_red_set, total_green_set, total_blue_set := 0, 0, 0
			cubes := strings.Split(set, ", ")

			for _, cube := range cubes {

				qnt_color := strings.Split(cube, " ")
				qnt, _ := strconv.Atoi(qnt_color[0])

				switch qnt_color[1] {
				case "red":
					total_red_set += qnt
					break
				case "green":
					total_green_set += qnt
					break
				case "blue":
					total_blue_set += qnt
					break
				}
			}

			if total_red_set > max_red {
				max_red = total_red_set
			}
			if total_green_set > max_green {
				max_green = total_green_set
			}
			if total_blue_set > max_blue {
				max_blue = total_blue_set
			}

		}

		total += (max_red * max_green * max_blue)

	}

	return total
}
