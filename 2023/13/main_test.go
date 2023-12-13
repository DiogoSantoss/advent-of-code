package main

import "testing"

var example = []string{
	"#.##..##.",
	"..#.##.#.",
	"##......#",
	"##......#",
	"..#.##.#.",
	"..##..##.",
	"#.#.##.#.",
	"",
	"#...##..#",
	"#....#..#",
	"..##..###",
	"#####.##.",
	"#####.##.",
	"..##..###",
	"#....#..#",
}

func maybe_complain(t *testing.T, output, result int) {
	if result != output {
		t.Fatalf("Obtained %d but correct value is %d\n", result, output)
	}
}

func TestPart1Small1(t *testing.T) {
	maybe_complain(t, 405, part1(example))
}

func TestPart1Big(t *testing.T) {
	maybe_complain(t, 34821, part1(parsed_input))
}

func TestPart2Small1(t *testing.T) {
	maybe_complain(t, 400, part2(example))
}

func TestPart2Big(t *testing.T) {
	maybe_complain(t, 36919, part2(parsed_input))
}
