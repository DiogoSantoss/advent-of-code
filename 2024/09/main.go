package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(disk []int) {

	expandedDisk := make([]string, 0)

	id := 0
	file := true
	for _, n := range disk {
		if file {
			for i := 0; i < n; i++ {
				expandedDisk = append(expandedDisk, strconv.Itoa(id))
			}
			id += 1
		} else {
			for i := 0; i < n; i++ {
				expandedDisk = append(expandedDisk, ".")
			}
		}
		file = !file
	}

	left, right := 0, len(expandedDisk)-1
	for left < right {
		for expandedDisk[left] != "." && left < right {
			left += 1
		}

		for expandedDisk[right] == "." && right > left {
			right -= 1
		}
		expandedDisk[left] = expandedDisk[right]
		expandedDisk[right] = "."
	}

	total := 0
	i := 0
	for expandedDisk[i] != "." {
		n, _ := strconv.Atoi(expandedDisk[i])
		total += i * n
		i++
	}
	fmt.Println(total)
}

type gap struct {
	pos    int
	length int
}

type file struct {
	id     int
	pos    int
	length int
}

func part2(disk []int) {

	files := make([]file, 0)
	gaps := make([]gap, 0)

	id := 0
	next_pos := 0
	is_file := true
	for _, n := range disk {
		if is_file {
			files = append(files, file{id: id, length: n, pos: next_pos})
			next_pos += n
			id += 1
		} else {
			gaps = append(gaps, gap{length: n, pos: next_pos})
			next_pos += n
		}
		is_file = !is_file
	}

	for i := len(files) - 1; i >= 0; i-- {
		for j := 0; j < len(gaps); j++ {
			// not only length but also dont move further from current pos
			if gaps[j].length >= files[i].length && gaps[j].pos < files[i].pos {
				// new file location
				files[i].pos = gaps[j].pos
				// update gap info
				gaps[j].length -= files[i].length
				gaps[j].pos += files[i].length
				fmt.Printf("moved file %d to pos %d\n",files[i].id, files[i].pos)
				break
			}
		}
	}

	total := 0
	for _, file := range files {
		for i := 0; i < file.length; i++ {
			total += (i + file.pos) * file.id
		}
	}
	fmt.Println(total)
}

func main() {
	bytes, _ := os.ReadFile("input.txt")
	text := strings.Fields(string(bytes))

	disk := make([]int, len(text[0]))
	for i, r := range text[0] {
		n, _ := strconv.Atoi(string(r))
		disk[i] = n
	}

	part1(disk)
	part2(disk)
}
