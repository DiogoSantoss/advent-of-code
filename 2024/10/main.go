package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type position struct {
	x, y int
}

var dirs = []position{{x: 1, y: 0}, {x: 0, y: 1}, {x: -1, y: 0}, {x: 0, y: -1}}

var m [][]int
var reached_top_positions []position

func dfs(c position, d position) {
	c.x += d.x
	c.y += d.y

	if m[c.y][c.x] == 9 && !slices.Contains(reached_top_positions, c) {
		reached_top_positions = append(reached_top_positions, c)
		return
	}

	for _, d := range dirs {
		new_y, new_x := c.y+d.y, c.x+d.x
		if new_y < 0 || new_y >= len(m) || new_x < 0 || new_x >= len(m[0]) {
			continue
		}
		if m[new_y][new_x] == m[c.y][c.x]+1 {
			dfs(c, d)
		}
	}
}

func dfs2(c position, d position) int {
	c.x += d.x
	c.y += d.y

	if m[c.y][c.x] == 9 {
		return 1
	}

	total := 0
	for _, d := range dirs {
		new_y, new_x := c.y+d.y, c.x+d.x
		if new_y < 0 || new_y >= len(m) || new_x < 0 || new_x >= len(m[0]) {
			continue
		}
		if m[new_y][new_x] == m[c.y][c.x]+1 {
			total += dfs2(c, d)
		}
	}

	return total
}

func main() {
	bytes, _ := os.ReadFile(os.Args[1])
	text := strings.Fields(string(bytes))

	m = make([][]int, len(text))
	trailheads := make([]position, 0)

	for i, line := range text {
		m[i] = make([]int, len(line))
		for j, r := range line {
			n, _ := strconv.Atoi(string(r))
			m[i][j] = n
			if n == 0 {
				trailheads = append(trailheads, position{x: j, y: i})
			}
		}
	}

	total := 0
	for _, th := range trailheads {
		for _, d := range dirs {
			new_y := th.y + d.y
			new_x := th.x + d.x
			if new_y < 0 || new_y >= len(m) || new_x < 0 || new_x >= len(m[0]) {
				continue
			}
			if m[new_y][new_x] == m[th.y][th.x]+1 {
				dfs(th, d)
			}
		}
		total += len(reached_top_positions)
		reached_top_positions = nil
	}

	fmt.Println(total)

	total = 0
	for _, th := range trailheads {
		for _, d := range dirs {
			new_y := th.y + d.y
			new_x := th.x + d.x
			if new_y < 0 || new_y >= len(m) || new_x < 0 || new_x >= len(m[0]) {
				continue
			}
			if m[new_y][new_x] == m[th.y][th.x]+1 {
				total += dfs2(th, d)
			}
		}
	}

	fmt.Println(total)
}
