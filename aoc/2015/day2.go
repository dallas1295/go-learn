package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func min(a, b, c int) int {
	if a < b && a < c {
		return a
	}
	if b < c {
		return b
	}
	return c
}

func getSquareFootage(s string) int {
	dim := strings.Split(s, "x")

	l, _ := strconv.Atoi(dim[0])
	w, _ := strconv.Atoi(dim[1])
	h, _ := strconv.Atoi(dim[2])

	lw := l * w
	wh := w * h
	hl := h * l

	res := 2*lw + 2*wh + 2*hl + min(lw, wh, hl)

	return res
}

func day2() {
	inputDir := "readthis"
	inputFile := filepath.Join(inputDir, "day2.txt")

	content, _ := os.ReadFile(inputFile)

	str := strings.TrimSpace(string(content))
	lines := strings.Split(str, "\n")
	// fmt.Println(lines)

	total := 0
	for _, line := range lines {
		total += getSquareFootage(line)
	}

	fmt.Println("Part 1 Answer:", total)
}
