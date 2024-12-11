package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func validUpdate(update []int, rules map[int][]int) bool {

	for i, value := range update {
		for _, ahead := range update[i+1:] {
			// If value in front should be before current then fail
			if slices.Contains(rules[ahead], value) {
				return false
			}
		}
	}

	return true
}

func fixUpdate(update []int, rules map[int][]int) int {
	
	// Bubble sort
	for {
		updated := false
		for i := 0; i < len(update)-1; i++ {
			curr := update[i]
			next := update[i+1]
			if slices.Contains(rules[next],curr) {
				update[i+1] = curr
				update[i] = next
				updated = true
			}
		}
		if !updated {
			break
		}
	}

	return update[len(update)/2]
}

func main() {
	bytes, _ := os.ReadFile("input.txt")
	text := strings.Split(string(bytes), "\n\n")

	rules_string := strings.Fields(text[0])
	updates_string := strings.Fields(text[1])

	// A: [B,C]  A should be before B,C
	rules := make(map[int][]int)
	for _, rule := range rules_string {
		v := strings.Split(rule, "|")
		v0, _ := strconv.Atoi(v[0])
		v1, _ := strconv.Atoi(v[1])
		rules[v0] = append(rules[v0], v1)
	}

	fmt.Println(rules)

	updates := make([][]int, len(updates_string))
	for i, update := range updates_string {
		values := strings.Split(update, ",")
		u := make([]int, len(values))
		for j, v := range values {
			u[j], _ = strconv.Atoi(v)
		}
		updates[i] = u
	}

	total_correct := 0
	total_incorrect := 0

	for _, update := range updates {
		if validUpdate(update, rules) {
			total_correct += update[len(update)/2]
		} else {
			total_incorrect += fixUpdate(update, rules)
		}
	}

	fmt.Println(total_correct)
	fmt.Println(total_incorrect)
}
