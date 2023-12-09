package main

import "testing"

var example = []string{
	"0 3 6 9 12 15",
	"1 3 6 10 15 21",
	"10 13 16 21 30 45",
}

func maybe_complain(t *testing.T, output, result int) {
	if result != output {
		t.Fatalf("Obtained %d but correct value is %d\n", result, output)
	}
}

func TestPart1Small(t *testing.T) {
	maybe_complain(t, 114, part1(example))
}

func TestPart1Big(t *testing.T) {
	maybe_complain(t, 1647269739, part1(parsed_input))
}

func TestPart2Small(t *testing.T) {
	maybe_complain(t, 2, part2(example))
}

func TestPart2Big(t *testing.T) {
	maybe_complain(t, 864, part2(parsed_input))
}
