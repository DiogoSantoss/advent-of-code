package main

import "testing"

var example1 = []string{
	"LLR",
	"",
	"AAA = (BBB, BBB)",
	"BBB = (AAA, ZZZ)",
	"ZZZ = (ZZZ, ZZZ)",
}

var example2 = []string{
	"RL",
	"",
	"AAA = (BBB, CCC)",
	"BBB = (DDD, EEE)",
	"CCC = (ZZZ, GGG)",
	"DDD = (DDD, DDD)",
	"EEE = (EEE, EEE)",
	"GGG = (GGG, GGG)",
	"ZZZ = (ZZZ, ZZZ)",
}

var example3 = []string{
	"LR",
	"",
	"11A = (11B, XXX)",
	"11B = (XXX, 11Z)",
	"11Z = (11B, XXX)",
	"22A = (22B, XXX)",
	"22B = (22C, 22C)",
	"22C = (22Z, 22Z)",
	"22Z = (22B, 22B)",
	"XXX = (XXX, XXX)",
}

func maybe_complain(t *testing.T, result, output int) {
	if result != output {
		t.Fatalf("Obtained %d but correct value is %d\n", result, output)
	}
}

func TestPart1Small1(t *testing.T) {
	maybe_complain(t, 6, part1(example1))
}

func TestPart1Small2(t *testing.T) {
	maybe_complain(t, 2, part1(example2))
}

func TestPart1Big(t *testing.T) {
	maybe_complain(t, 18673, part1(parsed_input))
}

func TestPart2Small(t *testing.T) {
	maybe_complain(t, 6, part2(example3))
}

func TestPart2Big(t *testing.T) {
	maybe_complain(t, 17972669116327, part2(parsed_input))
}
