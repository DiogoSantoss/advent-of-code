package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func part1(text string) {
	total := 0

	stateValue := ""
	state := ""

	for _, c := range text {

		switch state {
		case "":
			switch c {
			case 'm':
				state = "m"
			default:
				state = ""
			}
		case "m":
			switch c {
			case 'u':
				state = "u"
			default:
				state = ""
			}
		case "u":
			switch c {
			case 'l':
				state = "l"
			default:
				state = ""
			}
		case "l":
			switch c {
			case '(':
				state = "("
			default:
				state = ""
			}
		case "(":
			switch c {
			case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
				state = "d11"
			default:
				state = ""
			}
		case "d11":
			switch c {
			case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
				state = "d12"
			case ',':
				state = ","
			default:
				state = ""
			}
		case "d12":
			switch c {
			case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
				state = "d13"
			case ',':
				state = ","
			default:
				state = ""
			}
		case "d13":
			switch c {
			case ',':
				state = ","
			default:
				state = ""
			}
		case ",":
			switch c {
			case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
				state = "d21"
			default:
				state = ""
			}
		case "d21":
			switch c {
			case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
				state = "d22"
			case ')':
				state = ")"
			default:
				state = ""
			}
		case "d22":
			switch c {
			case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
				state = "d23"
			case ')':
				state = ")"
			default:
				state = ""
			}
		case "d23":
			switch c {
			case ')':
				state = ")"
			default:
				state = ""
			}
		case ")":

			// final state, record value
			var v1, v2 int
			fmt.Sscanf(stateValue, "mul(%d,%d)", &v1, &v2)
			fmt.Println(stateValue, v1, v2)
			total += v1 * v2

			// if m start from beginning and reset
			// stateValue
			switch c {
			case 'm':
				state = "m"
				stateValue = ""
			default:
				state = ""
			}
		}

		// if state is not empty then add character
		if state != "" {
			stateValue += string(c)
			// else state got reseted so reset state
		} else {
			stateValue = ""
		}
	}

	// if ends on ) then is complete sequence
	if state == ")" {
		f1 := strings.Split(stateValue, ",")
		value1 := strings.Split(f1[0], "(")
		value2 := strings.Split(f1[1], ")")
		v1, _ := strconv.Atoi(value1[1])
		v2, _ := strconv.Atoi(value2[0])
		fmt.Println(stateValue, v1, v2)
		total += v1 * v2
	}

	fmt.Println(total)
}

func part2(text string) {
	total := 0

	r := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|don't\(\)|do\(\)`)
	matches := r.FindAllStringSubmatch(text,-1)

	var v1,v2 int
	do := true
	for _,match := range matches {
		n,_ := fmt.Sscanf(match[0],"mul(%d,%d)",&v1,&v2)
		if n > 0 && do {
			total += v1*v2
		} else {
			do = (match[0] == "do()")
		}
	}

	fmt.Println(total)
}

func main() {
	bytes, _ := os.ReadFile("input.txt")
	text := string(bytes)

	//part1(text)
	part2(text)
}
