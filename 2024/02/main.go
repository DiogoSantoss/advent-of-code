package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func verifyLine(line []int) bool {
	slope := line[1] - line[0]
	for i := 1; i < len(line); i++ {
		res := slope * (line[i] - line[i-1])
		if res == 0 {
			return false
		} else if res > 0 && res > 3 {
			return false
		}
	}
	return true
}

func part1(matrix [][]int) {
	safeReports := 0
	for _, line := range matrix {
		if verifyLine(line) {
			safeReports += 1
		}
	}
	fmt.Println(safeReports)
}

func part2(matrix [][]int) {
	safeReports := 0
	for _, line := range matrix {
		if !verifyLine(line) {
			fmt.Print(line, " not safe")
			safeAfter := false
			for i := range line {
				newLine := make([]int, 0)
				newLine = append(newLine, line[:i]...)
				newLine = append(newLine, line[i+1:]...)
				if verifyLine(newLine) {
					safeAfter = true
					safeReports += 1
					break
				}
			}
			if safeAfter {
				fmt.Println(" but can be")
			} else {
				fmt.Println()
			}
		} else {
			fmt.Println(line, "safe")
			safeReports += 1
		}
	}
	fmt.Println(safeReports)
}

func main() {
	bytes, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(bytes), "\n")
	matrix := make([][]int, len(lines))
	for i, line := range lines {
		numbers := strings.Fields(line)
		matrix[i] = make([]int, len(numbers))
		for j, number := range numbers {
			n, _ := strconv.Atoi(number)
			matrix[i][j] = n
		}
	}

	part1(matrix)
	part2(matrix)
}
