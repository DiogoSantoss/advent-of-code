package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type equation struct {
	result  int
	numbers []int
}

func backtrack(eq equation, combination []string, operator string, idx int) bool {
	operators := []string{"+", "*", "||"}

	// Complete possible combination
	if idx >= len(combination) {
		total := eq.numbers[0]
		for i := 0; i < len(combination); i++ {
			switch combination[i] {
			case "+":
				total = total + eq.numbers[i+1]
			case "*":
				total = total * eq.numbers[i+1]
			case "||":
				
				//fmt.Println(total,eq.numbers[i+1],strconv.Itoa(total)+strconv.Itoa(eq.numbers[i+1]))
				total,_ = strconv.Atoi(strconv.Itoa(total)+strconv.Itoa(eq.numbers[i+1])) 
			}
		}
		
		//fmt.Println(eq, combination, total, eq.result)
		return total == eq.result
	}

	// Not complete combination
	combination[idx] = operator
	valid := false
	for _, operator := range operators {
		if backtrack(eq, combination, operator, idx+1) {
			valid = true
		}
	}
	return valid
}

func part1(equations []equation) {
	total := 0

	operators := []string{"+", "*", "||"}

	for _, eq := range equations {
		max_operators := len(eq.numbers) - 1
		combination := make([]string, max_operators)

		valid := false
		for _, operator := range operators {
			if backtrack(eq, combination, operator, 0) {
				valid = true
			}
		}

		if valid {
			total += eq.result
		}
	}

	fmt.Println(total)
}

func main() {
	bytes, _ := os.ReadFile("input.txt")
	text := strings.Split(string(bytes), "\n")

	equations := make([]equation, len(text))

	for i, eq := range text {
		temp := strings.Split(eq, ": ")
		result, _ := strconv.Atoi(temp[0])
		equations[i].result = result
		temp_values := strings.Fields(temp[1])
		equations[i].numbers = make([]int, len(temp_values))
		for j, v := range temp_values {
			result, _ = strconv.Atoi(v)
			equations[i].numbers[j] = result
		}
	}

	part1(equations)
}
