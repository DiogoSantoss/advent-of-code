package main

import "testing"

func TestPart1Small(t *testing.T) {
	input := []string{
		"32T3K 765",
		"T55J5 684",
		"KK677 28",
		"KTJJT 220",
		"QQQJA 483",
	}
	output := 6440
	result := part1(input)
	if result != output {
		t.Fatalf("Obtained %d but correct value is %d\n", result, output)
	}
}

func TestPart1Big(t *testing.T) {
	output := 249726565
	result := part1(parsed_input)
	if result != output {
		t.Fatalf("Obtained %d but correct value is %d\n", result, output)
	}
}

func TestPart2Small(t *testing.T) {
	input := []string{
		"32T3K 765",
		"T55J5 684",
		"KK677 28",
		"KTJJT 220",
		"QQQJA 483",
	}
	output := 5905
	result := part2(input)
	if result != output {
		t.Fatalf("Obtained %d but correct value is %d\n", result, output)
	}
}

func TestPart2Big(t *testing.T) {
	output := 251135960
	result := part2(parsed_input)
	if result != output {
		t.Fatalf("Obtained %d but correct value is %d\n", result, output)
	}
}
