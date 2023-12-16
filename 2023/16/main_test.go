package main

import "testing"

var example = []string{
	".|...\\....",
	"|.-.\\.....",
	".....|-...",
	"........|.",
	"..........",
	".........\\",
	"..../.\\\\..",
	".-.-/..|..",
	".|....-|.\\",
	"..//.|....",
}

func maybe_complain(t *testing.T, output, result int) {
	if result != output {
		t.Fatalf("Obtained %d but correct value is %d\n", result, output)
	}
}

func TestPart1Small1(t *testing.T) {
	maybe_complain(t, 46, part1(example))
}

func TestPart1Big(t *testing.T) {
	maybe_complain(t, 7477, part1(parsed_input))
}

func TestPart2Small1(t *testing.T) {
	maybe_complain(t, 51, part2(example))
}

func TestPart2Big(t *testing.T) {
	maybe_complain(t, 7853, part2(parsed_input))
}
