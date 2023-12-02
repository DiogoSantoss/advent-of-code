package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}
	output := 142
	result := part1(input)
	if result != output {
		t.Fatalf("Obtained %d but correct value is %d\n", result, output)
	}
}

func TestPart2(t *testing.T) {
	input := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four", "4nineeightseven2", "zoneight234", "7pqrstsixteen",
	}
	output := 281
	result := part2(input)
	if result != output {
		t.Fatalf("Obtained %d but correct value is %d\n", result, output)
	}
}
