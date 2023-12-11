package main

import "testing"

var example = []string{
	"...#......",
	".......#..",
	"#.........",
	"..........",
	"......#...",
	".#........",
	".........#",
	"..........",
	".......#..",
	"#...#.....",
}

func maybe_complain(t *testing.T, output, result int) {
	if result != output {
		t.Fatalf("Obtained %d but correct value is %d\n", result, output)
	}
}

func TestPart1Small1(t *testing.T) {
	maybe_complain(t, 374, part1(example))
}

func TestPart1Big(t *testing.T) {
	maybe_complain(t, 9623138, part1(parsed_input))
}

func TestPart2Small1(t *testing.T) {
	maybe_complain(t, 8410, part2(example,100))
}

func TestPart2Big(t *testing.T) {
	maybe_complain(t, 726820169514, part2(parsed_input,1000000))
}
