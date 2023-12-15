package main

import "testing"

var example = []string{
	"rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7",
}

func maybe_complain(t *testing.T, output, result int) {
	if result != output {
		t.Fatalf("Obtained %d but correct value is %d\n", result, output)
	}
}

func TestPart1Small1(t *testing.T) {
	maybe_complain(t, 1320, part1(example))
}

func TestPart1Big(t *testing.T) {
	maybe_complain(t, 509152, part1(parsed_input))
}

func TestPart2Small1(t *testing.T) {
	maybe_complain(t, 145, part2(example))
}

func TestPart2Big(t *testing.T) {
	maybe_complain(t, 244403, part2(parsed_input))
}
