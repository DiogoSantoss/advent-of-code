package main

import (
	"fmt"
	"os"
	"strings"
)

type position struct {
	x, y int
	dir  int // ^0 >1 v2 <3
}

var dirs = map[int][]int{0: {0, -1}, 1: {1, 0}, 2: {0, 1}, 3: {-1, 0}}

func (p position) nextPosition() (int, int) {
	return p.x + dirs[p.dir][0], p.y + dirs[p.dir][1]
}

func (p *position) step() {
	p.x += dirs[p.dir][0]
	p.y += dirs[p.dir][1]
}

func (p *position) nextDirection() {
	p.dir = (p.dir + 1) % 4
}

func part1(maze []string) {
	max_y := len(maze)
	max_x := len(maze[0])
	visited := make([][]rune, len(maze))

	// find guard and initiate visited
	var c position
	for i := 0; i < len(maze); i++ {
		visited[i] = make([]rune, len(maze[0]))
		for j := 0; j < len(maze[i]); j++ {
			visited[i][j] = '.'
			if maze[i][j] != '.' && maze[i][j] != '#' {
				var d int
				switch maze[i][j] {
				case '^':
					d = 0
				case '>':
					d = 1
				case 'v':
					d = 2
				case '<':
					d = 3
				}
				c = position{dir: d, y: i, x: j}
			}
		}
	}

	// walk guard
	total := 1
	visited[c.y][c.x] = 'X'
	for {
		new_x, new_y := c.nextPosition()
		if new_x >= max_x || new_y >= max_y || new_x < 0 || new_y < 0 {
			break
		}
		if maze[new_y][new_x] == '#' {
			c.nextDirection()
		} else {
			c.step()
			if visited[c.y][c.x] != 'X' {
				visited[c.y][c.x] = 'X'
				total += 1
			}
		}
	}

	fmt.Println(total)
}

func main() {
	bytes, _ := os.ReadFile(os.Args[1])
	maze := strings.Fields(string(bytes))

	part1(maze)
}
