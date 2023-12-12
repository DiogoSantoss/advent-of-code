package main

import (
	_ "embed"
	"flag"
	"fmt"
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

type cache_key struct {
	dr string
	g string // string(..) for easy compare but must use []uint8
}

var cache = make(map[cache_key]int)

func fd(dr string, g []uint8) int {

	// base case 1
	if len(dr) == 0 && len(g) == 0 {
		return 1
	}
	// base case 2
	if len(dr) == 0 {
		return 0
	}

	if v,found := cache[cache_key{dr,string(g)}]; found {
		return v
	}

	// try both '.' and '#'
	// with the optimization that '.' will be ignored
	// and therefore is only dr[1:]
	if dr[0] == '?' {
		r := fd(dr[1:], g) + fd("#"+dr[1:], g)
		cache[cache_key{dr,string(g)}] = r
		return r
	}

	// ignore '.'
	if dr[0] == '.' {
		r := fd(dr[1:], g)
		cache[cache_key{dr,string(g)}] = r
		return r
	}

	// verify group stuff
	if dr[0] == '#' {
		// no groups left and still # left
		// fail
		if len(g) == 0 {
			cache[cache_key{dr,string(g)}] = 0
			return 0
		}

		// get len of group
		n := g[0]
		// get max record group
		// either end at '.' or end of string
		idxDot := strings.Index(dr, ".")
		if idxDot == -1 {
			idxDot = len(dr)
		}

		// not enough to build group
		if idxDot < int(n) {
			cache[cache_key{dr,string(g)}] = 0
			return 0
		}

		// consume n to create group
		remaining := dr[n:]

		// end of record
		if len(remaining) == 0 {
			r := fd(remaining, g[1:])
			cache[cache_key{dr,string(g)}] = r
			return r
		}

		// should be . to create group, if not fail
		if remaining[0] == '#' {
			cache[cache_key{dr,string(g)}] = 0
			return 0
		}

		// continue from remaining[1:] since
		// remaining[0] is '.' or '?'
		// force '?' to be '.' by consuming it 
		r := fd(remaining[1:],g[1:])
		cache[cache_key{dr,string(g)}] = r
		return r
	}
	
	panic("you shouldn't be here")
}

func part1(input []string) int {
	total := 0

	damaged_records := []string{}
	groups_records := [][]uint8{}
	for _, line := range input {
		records := strings.Fields(line)
		damaged_records = append(damaged_records, records[0])
		groups := strings.Split(records[1], ",")
		group_record := []uint8{}
		for _, number := range groups {
			n, _ := strconv.Atoi(number)
			group_record = append(group_record, uint8(n))
		}
		groups_records = append(groups_records, group_record)
	}

	for i, damaged_record := range damaged_records {
		total += fd(damaged_record, groups_records[i])
	}

	return total
}

func part2(input []string) int {
	total := 0

	damaged_records := []string{}
	groups_records := [][]uint8{}
	for _, line := range input {

		records := strings.Fields(line)
		
		pattern := records[0]
		for i:=0;i<4;i++{
			records[0] = records[0] + "?" + pattern
		}
		
		damaged_records = append(damaged_records, records[0])

		groups := strings.Split(records[1], ",")
		group_record := []uint8{}
		for _, number := range groups {
			n, _ := strconv.Atoi(number)
			group_record = append(group_record, uint8(n))
		}

		gc := []uint8{}
		for i:=0;i<5;i++ {
			gc = append(gc, group_record...)
		}

		groups_records = append(groups_records, gc)
	}

	for i, damaged_record := range damaged_records {
		total += fd(damaged_record, groups_records[i])
	}

	return total
}
