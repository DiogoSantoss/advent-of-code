package main

import "testing"

func TestPart1Small(t *testing.T) {
	input := []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}
	output := 288
	result := part1(input)
	if result != output {
		t.Fatalf("Obtained %d but correct value is %d\n", result, output)
	}
}

func TestPart1Big(t *testing.T) {
	output := 211904
	result := part1(parsed_input)
	if result != output {
		t.Fatalf("Obtained %d but correct value is %d\n", result, output)
	}
}

func TestPart2Small(t *testing.T) {
	input := []string{
		"Time:      7 15 30",
		"Distance:  9 40 200",
	}
	output := 71503
	result := part2(input)
	if result != output {
		t.Fatalf("Obtained %d but correct value is %d\n", result, output)
	}
}

func TestPart2Big(t *testing.T) {
	output := 43364472
	result := part2(parsed_input)
	if result != output {
		t.Fatalf("Obtained %d but correct value is %d\n", result, output)
	}
}
