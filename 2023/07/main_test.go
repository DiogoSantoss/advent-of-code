package main

import "testing"

var example = []string{
	"32T3K 765",
	"T55J5 684",
	"KK677 28",
	"KTJJT 220",
	"QQQJA 483",
}

func maybe_complain(t *testing.T, result, output int) {
	if result != output {
		t.Fatalf("Obtained %d but correct value is %d\n", result, output)
	}
}

func TestPart1Small(t *testing.T) {
	maybe_complain(t, 6440, part1(example))
}

func TestPart1Big(t *testing.T) {
	maybe_complain(t, 249726565, part1(parsed_input))
}

func TestPart2Small(t *testing.T) {
	maybe_complain(t, 5905, part2(example))
}

func TestPart2Big(t *testing.T) {
	maybe_complain(t, 251135960, part2(parsed_input))
}
