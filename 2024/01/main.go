package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func part1(line1 []int, line2 []int) {
	diff := 0

	slices.Sort(line1)
	slices.Sort(line2)

	for i:=0; i<len(line1);i++ {
		d := line1[i]-line2[i]
		if d < 0 {
			d = -d
		}
		diff += d
	}

	fmt.Println(diff)
}

func part2(line1 []int, line2 []int) {
	siml := 0
	
	freq := make(map[int]int, len(line1))

	for _,n := range line2 {
		freq[n] += 1
	}

	for _,n := range line1 {
		siml += n * freq[n] 
	}

	fmt.Println(siml)
}

func main() {
	bytes, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(bytes),"\n")
	
	line1 := make([]int, len(lines))
	line2 := make([]int, len(lines))

	for idx,line := range lines {
		numbers := strings.Fields(line)
		n,_ := strconv.Atoi(numbers[0])
		line1[idx] = n
		n,_ = strconv.Atoi(numbers[1])
		line2[idx] = n
	}

	part1(line1,line2)
	part2(line1,line2)
}
