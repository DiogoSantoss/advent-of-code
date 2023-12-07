package main

import (
	_ "embed"
	"flag"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

//go:embed input.txt
var input string

// input splitted on \n
var parsed_input []string

func init() {
	// do this in init (not main) so test file has same input
	parsed_input = strings.Split(strings.TrimRight(input, "\n"), "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		games := parsed_input
		result := part1(games)
		fmt.Printf("Output: %d\n", result)
	} else if part == 2 {
		games := parsed_input
		result := part2(games)
		fmt.Printf("Output: %d\n", result)
	}
}

type hand_type int

const (
	high hand_type = iota
	one
	two
	three
	full
	four
	five
)

type play struct {
	t    hand_type
	hand string
	bid  int
}

func classify_part1(p *play) {

	dict := make(map[rune]int)
	for _, c := range p.hand {
		dict[c]++
	}

	highest_frequency := 0
	second_highest_frequency := 0
	for _, v := range dict {
		if v > highest_frequency {
			second_highest_frequency = highest_frequency
			highest_frequency = v
		} else if v > second_highest_frequency {
			second_highest_frequency = v
		}
	}

	if highest_frequency == 5 {
		p.t = five
		return
	} else if highest_frequency == 4 {
		p.t = four
		return
	} else if (highest_frequency == 3) && (second_highest_frequency == 2) {
		p.t = full
		return
	} else if highest_frequency == 3 {
		p.t = three
		return
	} else if (highest_frequency == 2) && (second_highest_frequency == 2) {
		p.t = two
		return
	} else if highest_frequency == 2 {
		p.t = one
		return
	} else {
		p.t = high
		return
	}
}

func classify_part2(p *play) {

	// special case bc we dont count 'J'
	// in frequency loop
	if p.hand == "JJJJJ" {
		p.t = five
		return
	}

	dict := make(map[rune]int)
	for _, c := range p.hand {
		dict[c]++
	}

	highest_frequency := 0
	second_highest_frequency := 0
	for k, v := range dict {
		if k == 'J' {
			continue
		}
		if v > highest_frequency {
			second_highest_frequency = highest_frequency
			highest_frequency = v
		} else if v > second_highest_frequency {
			second_highest_frequency = v
		}
	}

	// 'J' increases whatever hand they had before
	// but only once, i.e. only highest and not second_highest
	if highest_frequency+dict['J'] == 5 {
		p.t = five
		return
	} else if highest_frequency+dict['J'] == 4 {
		p.t = four
		return
	} else if (highest_frequency+dict['J'] == 3) && (second_highest_frequency == 2) {
		p.t = full
		return
	} else if highest_frequency+dict['J'] == 3 {
		p.t = three
		return
	} else if (highest_frequency+dict['J'] == 2) && (second_highest_frequency == 2) {
		p.t = two
		return
	} else if highest_frequency+dict['J'] == 2 {
		p.t = one
		return
	} else {
		p.t = high
		return
	}
}

func get_special_char(a rune) int {
	switch a {
	case 'T':
		return 0
	case 'J': // part 2 handled earlier in good_compare
		return 1
	case 'Q':
		return 2
	case 'K':
		return 3
	case 'A':
		return 4
	default:
		panic(1)
	}
}

func good_compare(a, b string, part int) int {
	for i := 0; i < len(a); i++ {
		rune_a := rune(a[i])
		rune_b := rune(b[i])

		if part == 2 {
			if (rune_a == 'J') && (rune_b == 'J') {
				continue
			} else if rune_a == 'J' {
				return 1
			} else if rune_b == 'J' {
				return -1
			}
		}

		if unicode.IsNumber(rune_a) && unicode.IsNumber(rune_b) {
			n_a, _ := strconv.Atoi(string(rune_a))
			n_b, _ := strconv.Atoi(string(rune_b))

			if n_a > n_b {
				return -1
			} else if n_a == n_b {
				continue
			} else {
				return 1
			}

		} else if unicode.IsNumber(rune_a) {
			return 1
		} else if unicode.IsNumber(rune_b) {
			return -1
		} else {

			t_a := get_special_char(rune_a)
			t_b := get_special_char(rune_b)

			if t_a > t_b {
				return -1
			} else if t_a == t_b {
				continue
			} else {
				return 1
			}
		}
	}
	return 0
}

func part(input []string, part int) int {
	total := 0
	ordered_plays := make(map[hand_type][]play, 0)

	for _, line := range input {
		stuff := strings.Fields(line)
		bid, _ := strconv.Atoi(stuff[1])
		play := play{
			hand: stuff[0],
			bid:  bid,
		}
		if part == 1 {
			classify_part1(&play)
		} else {
			classify_part2(&play)
		}
		ordered_plays[play.t] = append(ordered_plays[play.t], play)
	}

	// must be ordered
	hands := []hand_type{
		five,
		four,
		full,
		three,
		two,
		one,
		high,
	}
	rank := len(input)
	for _, hand := range hands {

		slices.SortFunc(ordered_plays[hand],
			func(a, b play) int {
				return good_compare(a.hand, b.hand, part)
			})

		for _, play := range ordered_plays[hand] {
			total += (rank * play.bid)
			rank--
		}
	}

	return total
}

func part1(input []string) int {
	return part(input, 1)
}

func part2(input []string) int {
	return part(input, 2)
}
