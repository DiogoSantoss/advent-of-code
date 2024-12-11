package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func validWord(text []string, i, j int) int {

	partial := 0
	where := make([]string, 8)

	// left right
	if j >= 3 {
		if string(text[i][j-1])+string(text[i][j-2])+string(text[i][j-3]) == "MAS" {
			partial += 1
			where[0] = "left"
		}
	}
	if j <= len(text[i])-4 {
		if string(text[i][j+1])+string(text[i][j+2])+string(text[i][j+3]) == "MAS" {
			partial += 1
			where[1] = "right"
		}
	}

	// up down
	if i >= 3 {
		if string(text[i-1][j])+string(text[i-2][j])+string(text[i-3][j]) == "MAS" {
			partial += 1
			where[2] = "up"
		}
	}
	if i <= len(text)-4 {
		if string(text[i+1][j])+string(text[i+2][j])+string(text[i+3][j]) == "MAS" {
			partial += 1
			where[3] = "down"
		}
	}

	// top diag
	if j >= 3 {
		if i >= 3 {
			if string(text[i-1][j-1])+string(text[i-2][j-2])+string(text[i-3][j-3]) == "MAS" {
				partial += 1
				where[4] = "up left"
			}
		}
		if i <= len(text)-4 {
			if string(text[i+1][j-1])+string(text[i+2][j-2])+string(text[i+3][j-3]) == "MAS" {
				partial += 1
				where[5] = "up right"
			}
		}
	}

	// bot diag
	if j <= len(text[i])-4 {
		if i >= 3 {
			if string(text[i-1][j+1])+string(text[i-2][j+2])+string(text[i-3][j+3]) == "MAS" {
				partial += 1
				where[6] = "bot left"
			}
		}
		if i <= len(text)-4 {
			if string(text[i+1][j+1])+string(text[i+2][j+2])+string(text[i+3][j+3]) == "MAS" {
				partial += 1
				where[7] = "bot right"
			}
		}
	}

	if partial != 0 {
		fmt.Println(i, j, partial, where)
	}

	return partial
}

func validCross(text []string, i, j int) bool {

	// cant be on edge
	if i < 1 || j < 1 || i > len(text)-2 || j > len(text[i])-2 {
		return false
	}

	// A B
	// C D
	// A,B,C,D
	crosses := []string{"MMSS", "MSMS", "SMSM", "SSMM"}
	pattern := string(text[i-1][j-1]) + string(text[i-1][j+1]) + string(text[i+1][j-1]) + string(text[i+1][j+1])

	if slices.Contains(crosses, pattern) {
		fmt.Println(i,j)
		return true
	}

	return false
}

func part1(text []string) {
	total := 0
	for i := 0; i < len(text); i++ {
		for j := 0; j < len(text[i]); j++ {
			if text[i][j] == 'X' {
				total += validWord(text, i, j)
			}
		}
	}

	fmt.Println(total)
}

func part2(text []string) {
	total := 0
	for i := 0; i < len(text); i++ {
		for j := 0; j < len(text[i]); j++ {
			if text[i][j] == 'A' && validCross(text,i,j) {
				total += 1
			}
		}
	}

	fmt.Println(total)
}

func main() {
	bytes, _ := os.ReadFile("input.txt")
	text := strings.Fields(string(bytes))
	part1(text)
	part2(text)
}
