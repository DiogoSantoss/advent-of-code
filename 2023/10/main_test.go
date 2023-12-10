package main

import "testing"

var example = []string{
	"..F7.",
	".FJ|.",
	"SJ.L7",
	"|F--J",
	"LJ...",
}

var example1 = []string{
	".....",
	".S-7.",
	".|.|.",
	".L-J.",
	".....",
}

var example2 = []string{
	"...........",
	".S-------7.",
	".|F-----7|.",
	".||.....||.",
	".||.....||.",
	".|L-7.F-J|.",
	".|..|.|..|.",
	".L--J.L--J.",
	"...........",
}

var example3 = []string{
	".F----7F7F7F7F-7....",
	".|F--7||||||||FJ....",
	".||.FJ||||||||L7....",
	"FJL7L7LJLJ||LJ.L-7..",
	"L--J.L7...LJS7F-7L7.",
	"....F-J..F7FJ|L7L7L7",
	"....L7.F7||L7|.L7L7|",
	".....|FJLJ|FJ|F7|.LJ",
	"....FJL-7.||.||||...",
	"....L---J.LJ.LJLJ...",
}

var example4 = []string{
	"FF7FSF7F7F7F7F7F---7",
	"L|LJ||||||||||||F--J",
	"FL-7LJLJ||||||LJL-77",
	"F--JF--7||LJLJ7F7FJ-",
	"L---JF-JLJ.||-FJLJJ7",
	"|F|F-JF---7F7-L7L|7|",
	"|FFJF7L7F-JF7|JL---7",
	"7-L-JL7||F7|L7F-7F7|",
	"L.L7LFJ|||||FJL7||LJ",
	"L7JLJL-JLJLJL--JLJ.L",
}

func maybe_complain(t *testing.T, output, result int) {
	if result != output {
		t.Fatalf("Obtained %d but correct value is %d\n", result, output)
	}
}

func TestPart1Small1(t *testing.T) {
	_,_,r := part1(example)
	maybe_complain(t, 8, r)
}

func TestPart1Small2(t *testing.T) {
	_,_,r := part1(example1)
	maybe_complain(t, 4, r)
}

func TestPart1Big(t *testing.T) {
	_,_,r := part1(parsed_input)
	maybe_complain(t, 6942, r)
}

func TestPart2Small1(t *testing.T) {
	maybe_complain(t, 4, part2(example2))
}

func TestPart2Small2(t *testing.T) {
	maybe_complain(t, 8, part2(example3))
}

func TestPart2Small3(t *testing.T) {
	maybe_complain(t, 10, part2(example4))
}

func TestPart2Big(t *testing.T) {
	maybe_complain(t, 297, part2(parsed_input))
}
