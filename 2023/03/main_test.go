package main

import "testing"

func TestPart1(t *testing.T) {
	input := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}
	output := 4361
	result := part1(input)
	if result != output {
		t.Fatalf("Obtained %d but correct value is %d\n", result, output)
	}
}

func TestPart2(t *testing.T) {
	input := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}
	output := 467835
	result := part2(input)
	if result != output {
		t.Fatalf("Obtained %d but correct value is %d\n", result, output)
	}
}
