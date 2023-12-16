package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"strings"
	"sync"
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

type pos struct {
	row, col int
}

type dir struct {
	row, col int
}

type vec struct {
	pos pos
	dir dir
}

var (
	left  = dir{row: 0, col: -1}
	right = dir{row: 0, col: 1}
	up    = dir{row: -1, col: 0}
	down  = dir{row: 1, col: 0}
)

func part1(input []string) int {
	return megaLazer(input, vec{pos: pos{0, 0}, dir: right})
}

func part2(input []string) int {
	total := 0
	m := len(input)
	n := len(input[0])

	var wg sync.WaitGroup
	res := make(chan int, 2*m+2*n)

	wg.Add(2*m + 2*n)

	for i := 0; i < m; i++ {
		go func(i int) {
			defer wg.Done()
			count := megaLazer(input, vec{pos: pos{i, 0}, dir: right})
			res <- count
		}(i)
		go func(i int) {
			defer wg.Done()
			count := megaLazer(input, vec{pos: pos{i, n - 1}, dir: left})
			res <- count
		}(i)
	}

	for j := 0; j < n; j++ {
		go func(j int) {
			defer wg.Done()
			count := megaLazer(input, vec{pos: pos{0, j}, dir: down})
			res <- count
		}(j)
		go func(j int) {
			defer wg.Done()
			count := megaLazer(input, vec{pos: pos{m - 1, j}, dir: up})
			res <- count
		}(j)
	}

	wg.Wait()
	close(res)

	for count := range res {
		total = max(total, count)
	}

	return total
}

func megaLazer(grid []string, start vec) int {
	visited := map[vec]bool{start: true}
	queue := []vec{start}
	res := map[pos]bool{start.pos: true}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, nei := range getNeighbours(grid, curr) {
			if !visited[nei] {
				visited[nei] = true
				res[nei.pos] = true
				queue = append(queue, nei)
			}
		}
	}

	return len(res)
}

func getNeighbours(grid []string, curr vec) []vec {
	r, c := curr.pos.row, curr.pos.col
	m, n := len(grid), len(grid[0])
	nr, nc := curr.pos.row+curr.dir.row, curr.pos.col+curr.dir.col

	switch curr.dir {
	case left, right:
		// beam passes through
		if grid[r][c] == '-' || grid[r][c] == '.' {
			if nr < 0 || nr >= m || nc < 0 || nc >= n {
				return nil
			}
			return []vec{{pos: pos{nr, nc}, dir: curr.dir}}
		}
		// split beam
		if grid[r][c] == '|' {
			var neighbours []vec
			for _, dir := range []dir{up, down} {
				if r+dir.row >= 0 && r+dir.row < m {
					neighbours = append(neighbours, vec{pos: pos{r + dir.row, c}, dir: dir})
				}
			}
			return neighbours
		}

		// reflect beam
		if r+down.row < m && (grid[r][c] == '/' && curr.dir == left || grid[r][c] == '\\' && curr.dir == right) {
			return []vec{{pos: pos{r + down.row, c}, dir: down}}
		}
		if r+up.row >= 0 && (grid[r][c] == '/' && curr.dir == right || grid[r][c] == '\\' && curr.dir == left) {
			return []vec{{pos: pos{r + up.row, c}, dir: up}}
		}
		return nil

	case up, down:
		// beam passes through
		if grid[r][c] == '|' || grid[r][c] == '.' {
			if nr < 0 || nr >= m || nc < 0 || nc >= n {
				return nil
			}
			return []vec{{pos: pos{nr, nc}, dir: curr.dir}}
		}

		// split beam
		if grid[r][c] == '-' {
			var neighbours []vec
			for _, dir := range []dir{left, right} {
				if c+dir.col >= 0 && c+dir.col < n {
					neighbours = append(neighbours, vec{pos: pos{r, c + dir.col}, dir: dir})
				}
			}
			return neighbours
		}

		// reflect beam
		if c+right.col < n && (grid[r][c] == '/' && curr.dir == up || grid[r][c] == '\\' && curr.dir == down) {
			return []vec{{pos: pos{r, c + right.col}, dir: right}}
		}
		if c+left.col >= 0 && (grid[r][c] == '/' && curr.dir == down || grid[r][c] == '\\' && curr.dir == up) {
			return []vec{{pos: pos{r, c + left.col}, dir: left}}
		}
	}

	return nil
}

func max(nums ...int) int {
	res := math.MinInt
	for _, num := range nums {
		if num > res {
			res = num
		}
	}
	return res
}
