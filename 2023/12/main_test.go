package main

import "testing"

var example = []string{
	"???.### 1,1,3",
	".??..??...?##. 1,1,3",
	"?#?#?#?#?#?#?#? 1,3,1,6",
	"????.#...#... 4,1,1",
	"????.######..#####. 1,6,5",
	"?###???????? 3,2,1",
}

func maybe_complain(t *testing.T, output, result int) {
	if result != output {
		t.Fatalf("Obtained %d but correct value is %d\n", result, output)
	}
}

func TestPart1Small1(t *testing.T) {
	maybe_complain(t, 21, part1(example))
}

func TestPart1Big(t *testing.T) {
	maybe_complain(t, 8193, part1(parsed_input))
}

func TestPart2Small1(t *testing.T) {
	maybe_complain(t, 525152, part2(example))
}

func TestPart2Big(t *testing.T) {
	maybe_complain(t, 45322533163795, part2(parsed_input))
}
