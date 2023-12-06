package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"strconv"
	"strings"
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

func max_dist(v int, t_hold int, t_max int) int {
	t_rem := t_max - t_hold
	return v * t_rem
}

func part1(input []string) int {
	total := 0

	times := strings.Fields(input[0])
	record_distances := strings.Fields(input[1])

	for i := 1; i < len(times); i++ {
		rec_dist, _ := strconv.Atoi(record_distances[i])
		rec_dist_f := float64(rec_dist)

		t_total, _ := strconv.Atoi(times[i])
		t_total_f := float64(t_total)

		t_hold_f := math.Round(rec_dist_f / (t_total_f - 1))

		k := 0
		for i := int(t_hold_f); i < t_total; i++ {
			if max_dist(i, i, int(t_total_f)) > rec_dist {
				k++
			}
		}
		if total == 0 {
			total = k
		} else {
			total *= k
		}
	}

	return total
}

/*

Alternatively, you could solve using quadratic formula
but this was rounding errors and i don't feel
like finding the real problem

MATH:

(t_total - t_hold) * t_hold > rec_dist

want to solve with respect to t_hold

t_hold^2 - t_total*t_hold + rec_dist > 0

use quadratic to find roots

ax^2 + bx + c = 0

y = -b +/- (sqrt(b^2-4ac))/2a

CODE:

total := 0.0
rec_dist := float64(race_record)
t_total := float64(race_time)

det := math.Pow(t_total,2)-4*rec_dist
root_1 := (t_total - (math.Sqrt(det)/2.0))
root_2 := (t_total + (math.Sqrt(det)/2.0))

total = math.Abs(math.Ceil(root_1) - math.Floor(root_2)) + 1

*/

func part2(input []string) int {

	times := strings.Fields(input[0])
	record_distances := strings.Fields(input[1])

	race_time, _ := strconv.Atoi(strings.Join(times[1:], ""))
	race_record, _ := strconv.Atoi(strings.Join(record_distances[1:], ""))

	rec_dist := race_record
	rec_dist_f := float64(rec_dist)

	t_total := race_time
	t_total_f := float64(t_total)

	// dist = vel * t
	
	// vel = t_hold
	// t = t_total - t_hold
	
	// rec_dist = t_hold*t_total - t_hold^2
	// rec_dist = t_hold(t_total-1)
	
	// t_hold = rec_dist/(t_total-1)  min to achieve rec_dist
	t_hold_f := math.Round(rec_dist_f / (t_total_f - 1))

	k := 0
	for i := int(t_hold_f); i < t_total; i++ {
		if max_dist(i, i, int(t_total_f)) > rec_dist {
			k++
		}
	}
		
	return k
}
