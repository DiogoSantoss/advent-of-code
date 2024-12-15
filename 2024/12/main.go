package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type coord struct {
	x, y int
}

func partialPerim(m [][]rune, c coord) (total int) {

	r := m[c.y][c.x]
	dirs := []coord{{x: 1, y: 0}, {x: 0, y: 1}, {x: -1, y: 0}, {x: 0, y: -1}}

	for _, d := range dirs {
		nx, ny := d.x+c.x, d.y+c.y
		if nx < 0 || ny < 0 || nx > len(m[0])-1 || ny > len(m)-1 {
			total += 1
		} else if nx >= 0 && ny >= 0 && nx < len(m[0]) && ny < len(m) && m[ny][nx] != r {
			total += 1
		}
	}

	return total
}

func charAt(m [][]rune, x, y int) rune {
	if x < 0 || y < 0 || x >= len(m[0]) || y >= len(m) {
		return ' '
	} else {
		return m[y][x]
	}
}

func boolToInt(b bool) int {
	if b {
		return 0
	}
	return 1
}
func intToBool(i int) bool {
	return i == 1
}

/*
Source: https://github.com/derfritz/AoC24/blob/main/Day12/solution.js
Why?
every point can have some corners:

	XXX
	XOX
	XXX
has 4 corners

	XOX
	XOX
	XXX
has 3 corners

	XOX
	OOX
	XXX
has 2 corners

	XXX
	XOO
	XOO
has 1 corner

	XOX
	XOX
	XOX
has 0 corners
so compute how many corners each point has and this will
be equal to the number of sides
*/
func computeSides(m [][]rune, points []coord) (total int) {

	for _, p := range points {
		left := boolToInt(charAt(m, p.x-1, p.y) == m[p.y][p.x])
		right := boolToInt(charAt(m, p.x+1, p.y) == m[p.y][p.x])
		up := boolToInt(charAt(m, p.x, p.y-1) == m[p.y][p.x])
		down := boolToInt(charAt(m, p.x, p.y+1) == m[p.y][p.x])
		upLeft := boolToInt(charAt(m, p.x-1, p.y-1) == m[p.y][p.x])
		upRight := boolToInt(charAt(m, p.x+1, p.y-1) == m[p.y][p.x])
		downLeft := boolToInt(charAt(m, p.x-1, p.y+1) == m[p.y][p.x])
		downRight := boolToInt(charAt(m, p.x+1, p.y+1) == m[p.y][p.x])

		if left+up == 2 || (left+up == 0 && intToBool(upLeft)) {
			total += 1
		}
		if left+down == 2 || (left+down == 0 && intToBool(downLeft)) {
			total += 1
		}
		if right+up == 2 || (right+up == 0 && intToBool(upRight)) {
			total += 1
		}
		if right+down == 2 || (right+down == 0 && intToBool(downRight)) {
			total += 1
		}
	}

	return total
}

func main() {
	bytes, _ := os.ReadFile(os.Args[1])
	text := strings.Fields(string(bytes))

	m := make([][]rune, len(text))
	for i, l := range text {
		m[i] = make([]rune, len(l))
		for j, r := range l {
			m[i][j] = r
		}
	}

	total1 := 0
	total2 := 0

	dirs := []coord{{x: 1, y: 0}, {x: 0, y: 1}, {x: -1, y: 0}, {x: 0, y: -1}}

	visited := make(map[coord]bool)

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {

			c := coord{x: j, y: i}

			// dont visit again
			if visited[c] {
				continue
			}

			visited[c] = true
			area := 1
			perim := partialPerim(m, c)
			r := m[i][j]
			// part 2
			points := []coord{c}

			// fill stack with same region neighbors (and not visited)
			stack := make([]coord, 0)
			for _, d := range dirs {
				nx, ny := c.x+d.x, c.y+d.y
				if nx >= 0 && ny >= 0 &&
					nx < len(m[0]) && ny < len(m) &&
					!visited[coord{x: nx, y: ny}] &&
					m[ny][nx] == r {
					stack = append(stack, coord{x: nx, y: ny})
				}
			}

			for len(stack) > 0 {

				// increase area and mark as visited
				c = stack[len(stack)-1]
				area += 1
				perim += partialPerim(m, c)
				visited[c] = true
				stack = stack[:len(stack)-1]

				// part 2
				points = append(points, c)

				// fill stack with same region neighbors (and not visited)
				for _, d := range dirs {
					nx, ny := c.x+d.x, c.y+d.y
					if nx >= 0 && ny >= 0 &&
						nx < len(m[0]) && ny < len(m) &&
						!visited[coord{x: nx, y: ny}] &&
						!slices.Contains(stack, coord{x: nx, y: ny}) &&
						m[ny][nx] == r {
						stack = append(stack, coord{x: nx, y: ny})
					}
				}

			}
			sides := computeSides(m, points)
			total1 += area * perim
			total2 += area * sides
		}
	}

	fmt.Println(total1)
	fmt.Println(total2)
}
