package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type nb struct {
	n,b int
}

var memo map[nb]int = make(map[nb]int)

func transformStone(stone int, blinks int) int {
	if blinks == 0 {
		return 1
	}

	res,ok := memo[nb{n:stone,b:blinks}]
	if ok {
		return res
	}

	if stone == 0 {
		res = transformStone(1, blinks-1)
	} else if is_even, n1, n2 := isEvenLength(stone); is_even {
		res = transformStone(n1,blinks-1)+transformStone(n2,blinks-1)
	} else {
		res = transformStone(stone*2024, blinks-1)
	}

	memo[nb{n:stone,b: blinks}] = res
	return res
}

func isEvenLength(stone int) (bool,int,int) {
	s := strconv.Itoa(stone)
	even := len(s)%2 == 0
	s1,s2 := s[:len(s)/2],s[len(s)/2:]
	n1,_:=strconv.Atoi(s1)
	n2,_:=strconv.Atoi(s2)
	return even,n1,n2
}

func main() {
	bytes, _ := os.ReadFile(os.Args[1])
	line := strings.Fields(string(bytes))

	total_blinks := 75

	total := 0
	for _,stone := range line {
		s,_ := strconv.Atoi(stone)
		total += transformStone(s,total_blinks)
	}
	fmt.Println(total)
}
