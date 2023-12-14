package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Schematic []string

func readSchematic(filename string) Schematic {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var schematic []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		schematic = append(schematic, line)
	}

	return schematic
}

func (s Schematic) findParts() []int {
	var parts []int

	re := regexp.MustCompile(`\d+`)
	for i, line := range s {
		matches := re.FindAllStringIndex(line, -1) // Find all matches
		if matches == nil {
			continue
		}

		for _, idx := range matches {
			part := s.getValidPart(i, idx[0], idx[1])
			if part >= 0 {
				parts = append(parts, part)
			}
		}
	}

	return parts
}

/*
 * Return part number if valid, -1 otherwise.
 */
func (s Schematic) getValidPart(row, colStart, colEnd int) int {
	isValidPart := false

	m := len(s)
	n := len(s[row])

	var j0, j1 int
	if colStart == 0 {
		j0 = 0
	} else {
		j0 = colStart - 1
	}
	if colEnd == n {
		j1 = n
	} else {
		j1 = colEnd + 1
	}

	if row > 0 {
		for j := j0; j < j1; j++ {
			ch := s[row-1][j : j+1]
			if ch != "." {
				isValidPart = true
			}
		}
	}

	if row < m-1 {
		for j := j0; j < j1; j++ {
			ch := s[row+1][j : j+1]
			if ch != "." {
				isValidPart = true
			}
		}
	}

	if colStart != 0 && s[row][colStart-1:colStart] != "." {
		isValidPart = true
	}

	if colEnd != n && s[row][colEnd:colEnd+1] != "." {
		isValidPart = true
	}

	if !isValidPart {
		return -1
	}

	partNumber, err := strconv.Atoi(s[row][colStart:colEnd])
	if err != nil {
		log.Fatal(err)
	}
	return partNumber
}

func main() {
	s := readSchematic("input.txt")
	parts := s.findParts()
	sum := 0
	for _, part := range parts {
		sum += part
	}
	fmt.Println(sum)
}
