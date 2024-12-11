package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type position struct {
	x, y int
}

func (p1 position) add(p2 position) position {
	return position{x: p1.x + p2.x, y: p1.y + p2.y}
}

func (p1 position) sub(p2 position) position {
	return position{x: p1.x - p2.x, y: p1.y - p2.y}
}

func part1(antennas map[rune][]position, max_x, max_y int) {

	total := 0
	antinodes := make([]position, 0)
	for _, antennas_list := range antennas {
		for i := 0; i < len(antennas_list); i++ {
			for j := i; j < len(antennas_list); j++ {
				if i == j {
					continue
				}

				a := antennas_list[i]
				b := antennas_list[j]
				// b-a hence vector from a to b
				v := position{x: b.x - a.x, y: b.y - a.y}

				// compute antinodes (an)
				// b+v
				an := b.add(v)
				if an.x < max_x && an.y < max_y && an.x >= 0 && an.y >= 0 {
					if !slices.Contains(antinodes, an) {
						total += 1
						antinodes = append(antinodes, an)
					}
				}
				// a-v
				an = a.sub(v)
				if an.x < max_x && an.y < max_y && an.x >= 0 && an.y >= 0 {
					if !slices.Contains(antinodes, an) {
						total += 1
						antinodes = append(antinodes, an)
					}
				}
			}
		}
	}

	fmt.Println(total)
}

func part2(antennas map[rune][]position, max_x, max_y int, matrix [][]rune) {

	total := 0
	antinodes := make([]position, 0)
	for _, antennas_list := range antennas {
		for i := 0; i < len(antennas_list); i++ {
			for j := i; j < len(antennas_list); j++ {
				if i == j {
					continue
				}

				a := antennas_list[i]
				b := antennas_list[j]
				if !slices.Contains(antinodes, a) {
					total += 1
					antinodes = append(antinodes, a)
				}
				if !slices.Contains(antinodes, b) {
					total += 1
					antinodes = append(antinodes, b)
				}
				// b-a hence vector from a to b
				v := position{x: b.x - a.x, y: b.y - a.y}

				// compute antinodes (an)
				// b+v
				base := b
				for {
					base = base.add(v)
					if base.x < max_x && base.y < max_y && base.x >= 0 && base.y >= 0 {
						if !slices.Contains(antinodes, base) {
							total += 1
							antinodes = append(antinodes, base)
						}
					} else {
						break
					}
				}
				// a-v
				base = a
				for {
					base = base.sub(v)
					if base.x < max_x && base.y < max_y && base.x >= 0 && base.y >= 0 {
						if !slices.Contains(antinodes, base) {
							total += 1
							antinodes = append(antinodes, base)
						}
					} else {
						break
					}
				}
			}
		}
	}

	for _, a := range antinodes {
		if matrix[a.y][a.x] == '.' {
			matrix[a.y][a.x] = '#'
		}
	}

	for i := 0; i < max_y; i++ {
		for j := 0; j < max_y; j++ {
			fmt.Print(string(matrix[i][j]))
		}
		fmt.Println()
	}

	fmt.Println(total)
}

func main() {
	bytes, _ := os.ReadFile("input.txt")
	text := strings.Fields(string(bytes))
	max_x, max_y := len(text[0]), len(text)

	matrix := make([][]rune, len(text))
	for i := 0; i < len(text); i++ {
		matrix[i] = make([]rune, len(text[i]))
		for j := 0; j < len(text[0]); j++ {
			matrix[i][j] = rune(text[i][j])
		}
	}

	antennas := make(map[rune][]position)
	for i, line := range text {
		for j, r := range line {
			if r != '.' {
				antennas[r] = append(antennas[r], position{x: j, y: i})
			}
		}
	}

	part1(antennas, max_x, max_y)
	part2(antennas, max_x, max_y, matrix)
}
