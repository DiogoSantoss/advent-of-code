package main

import "testing"

var example = []string{
	"O....#....",
	"O.OO#....#",
	".....##...",
	"OO.#O....O",
	".O.....O#.",
	"O.#..O.#.#",
	"..O..#O..O",
	".......O..",
	"#....###..",
	"#OO..#....",
}

func maybe_complain(t *testing.T, output, result int) {
	if result != output {
		t.Fatalf("Obtained %d but correct value is %d\n", result, output)
	}
}

func TestPart1Small1(t *testing.T) {
	maybe_complain(t, 136, part1(example))
}

func TestPart1Big(t *testing.T) {
	maybe_complain(t, 109665, part1(parsed_input))
}

func TestPart2Small1(t *testing.T) {
	maybe_complain(t, 64, part2(example))
}

func TestPart2Big(t *testing.T) {
	maybe_complain(t, 96061, part2(parsed_input))
}
