package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func day1() {
	inputDir := "readthis"
	inputFile := filepath.Join(inputDir, "day1.txt")

	content, _ := os.ReadFile(inputFile)

	// fmt.Println(string(content))
	str := strings.TrimSpace(string(content))
	lines := strings.Split(str, "")

	// fmt.Println(lines)

	floor := 0

	for _, char := range lines {
		if char == "(" {
			floor++
		} else if char == ")" {
			floor--
		}
	}

	fmt.Println("Part 1 Answer:", floor)

	// part 2

	floor = 0
	step := 0

	for _, char := range lines {
		if char == "(" {
			floor++
			step++
		} else if char == ")" {
			floor--
			step++
			if floor == -1 {
				break
			}
		}
	}

	fmt.Println("Part 2 Answer:", step)
}
