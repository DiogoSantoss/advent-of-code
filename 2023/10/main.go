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
		_, _, result := part1(games)
		fmt.Printf("Output: %d\n", result)
	} else if part == 2 {
		games := parsed_input
		result := part2(games)
		fmt.Printf("Output: %d\n", result)
	}
}

func get_next_pos(pos_row, pos_col, prev_row, prev_col int, input []string) (int, int) {
	switch input[pos_row][pos_col] {
	case '|':
		if prev_row+1 == pos_row {
			return pos_row + 1, pos_col
		} else {
			return pos_row - 1, pos_col
		}
	case '-':
		if prev_col+1 == pos_col {
			return pos_row, pos_col + 1
		} else {
			return pos_row, pos_col - 1
		}
	case 'L':
		if prev_row+1 == pos_row {
			return pos_row, pos_col + 1
		} else {
			return pos_row - 1, pos_col
		}
	case 'J':
		if prev_row+1 == pos_row {
			return pos_row, pos_col - 1
		} else {
			return pos_row - 1, pos_col
		}
	case '7':
		if prev_row-1 == pos_row {
			return pos_row, pos_col - 1
		} else {
			return pos_row + 1, pos_col
		}
	case 'F':
		if prev_row-1 == pos_row {
			return pos_row, pos_col + 1
		} else {
			return pos_row + 1, pos_col
		}
	default:
		panic("unknown char")
	}
}

func part1(input []string) (map[int][]int, []position, int) {

	s_row, s_col := 0, 0
	for row, line := range input {
		for col, c := range line {
			if c == 'S' {
				s_row = row
				s_col = col
			}
		}
	}

	loop := make(map[int][]int, 0)
	path := make([]position, 0)

	steps := 0
	pos_row, pos_col := s_row, s_col
	prev_row, prev_col := -1, -1
	for true {

		// handle last step
		if (pos_row == s_row) && (pos_col == s_col) && (steps > 0) {
			break
		}

		// handle first step
		if (pos_row == s_row) && (pos_col == s_col) && (steps == 0) {
			if (pos_row+1 < len(input)) && (input[pos_row+1][pos_col] == '|') || (input[pos_row+1][pos_col] == 'L') || (input[pos_row+1][pos_col] == 'J') {
				prev_row, prev_col = pos_row, pos_col
				pos_row = pos_row + 1
			} else if (pos_row-1 >= 0) && (input[pos_row-1][pos_col] == '|') || (input[pos_row-1][pos_col] == '7') || (input[pos_row-1][pos_col] == 'F') {
				prev_row, prev_col = pos_row, pos_col
				pos_row = pos_row - 1
			} else if (pos_col+1 < len(input[0])) && (input[pos_row][pos_col+1] == '-') || (input[pos_row][pos_col+1] == '7') || (input[pos_row][pos_col+1] == 'J') {
				prev_row, prev_col = pos_row, pos_col
				pos_col = pos_col + 1
			} else {
				prev_row, prev_col = pos_row, pos_col
				pos_col = pos_col - 1
			}

			steps += 1

			path = append(path, position{
				y: prev_col,
				x: prev_row,
			})
			loop[prev_row] = append(loop[prev_row], prev_col)
			continue
		}

		// handle all other steps
		tmp_row, tmp_col := get_next_pos(pos_row, pos_col, prev_row, prev_col, input)
		prev_row, prev_col = pos_row, pos_col
		pos_row, pos_col = tmp_row, tmp_col

		path = append(path, position{
			y: prev_col,
			x: prev_row,
		})
		loop[prev_row] = append(loop[prev_row], prev_col)
		steps += 1
	}

	if steps%2 == 0 {
		return loop, path, steps / 2
	} else {
		return loop, path, steps/2 - 1
	}
}

type position struct {
	x int
	y int
}

func part2(input []string) int {

	replaceWith := map[string]string{}
	replaceWith["J"] = "┘"
	replaceWith["L"] = "└"
	replaceWith["7"] = "┐"
	replaceWith["F"] = "┌"
	replaceWith["|"] = "│"
	replaceWith["-"] = "─"

	loop, path, _ := part1(input)

	total := 0
	for k, v := range loop {

		a := strings.Split(input[k], "")
		for _, j := range v {
			if a[j] != "S" {
				a[j] = replaceWith[a[j]]
			} else {
				a[j] = replaceWith[string(replaceSWith(path[0], path[1], path[len(path)-1], input))]
			}
		}

		count := 0
		isInside := 0
		last := "-"
		for _, char := range a {

			if char == "│" {
				isInside++
				continue
			}
			if char == "─" {
				continue
			}

			if last == "-" && (char == "┘" || char == "└" || char == "┐" || char == "┌") {
				last = char
				continue
			} else if last != "-" && (char == "┘" || char == "└" || char == "┐" || char == "┌") {
				if last == "└" && char == "┐" {
					isInside++
				}

				if last == "┌" && char == "┘" {
					isInside++
				}
				last = "-"
				continue
			}

			if isInside%2 != 0 {
				count++
			}

		}
		total += count

	}
	return total
}

// --- Stolen ---
// I manually replaced S with the correct character
// but then tests would fail because input was different for part 1
// and part 2
// already spent too much time on this, so I just stole this code to
// replace S with the correct character
func replaceSWith(s, first, last position, input []string) byte {

	for _, char := range "-|JL7F" {
		duplicateInput := input
		a := strings.Split(duplicateInput[s.x], "")
		a[s.y] = string(char)
		duplicateInput[s.x] = strings.Join(a, "")
		x, y := findNext(s.x, s.y, duplicateInput, last.x, last.y)
		if x == first.x && y == first.y {
			return byte(char)
		}
	}

	return 'S'
}

func findNext(x, y int, input []string, lastx, lasty int) (int, int) {
	if input[x][y] == '-' {
		if y != 0 && lasty == (y-1) {
			y++
			return x, y
		}

		if y != (len(input[x])-1) && lasty == (y+1) {
			y--
			return x, y
		}
	}

	if input[x][y] == 'J' {
		if x != 0 && y != 0 && lastx == (x-1) {
			y--
			return x, y
		}

		if x != 0 && y != 0 && lasty == (y-1) {
			x--
			return x, y
		}
	}

	if input[x][y] == '|' {
		if x != 0 && x != len(input)-1 && lastx == (x-1) {
			x++
			return x, y
		}

		if x != 0 && x != len(input)-1 && lastx == (x+1) {
			x--
			return x, y
		}
	}

	if input[x][y] == 'L' {
		if x != 0 && y != len(input[x])-1 && lastx == (x-1) {
			y++
			return x, y
		}

		if x != 0 && y != len(input[x])-1 && lasty == (y+1) {
			x--
			return x, y
		}
	}

	if input[x][y] == '7' {
		if y != 0 && x != (len(input)-1) && lasty == (y-1) {
			x++
			return x, y
		}

		if x != (len(input)-1) && y != 0 && lastx == (x+1) {
			y--
			return x, y
		}
	}

	if input[x][y] == 'F' {
		if x != (len(input)-1) && y != (len(input[x])-1) && lasty == (y+1) {
			x++
			return x, y
		}

		if x != (len(input)-1) && y != (len(input[x])-1) && lastx == (x+1) {
			y++
			return x, y
		}
	}

	return x, y
}
