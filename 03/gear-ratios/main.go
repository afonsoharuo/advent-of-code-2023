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

type StarCoordinates struct {
	I int
	J int
}

type Part struct {
	Number            int
	NeighbouringStars []StarCoordinates
}

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

func (s Schematic) findParts() []Part {
	var parts []Part

	re := regexp.MustCompile(`\d+`)
	for i, line := range s {
		matches := re.FindAllStringIndex(line, -1) // Find all matches
		if matches == nil {
			continue
		}

		for _, idx := range matches {
			part := s.getValidPart(i, idx[0], idx[1])
			if part != nil {
				parts = append(parts, *part)
			}
		}
	}

	return parts
}

/*
 * Return part number if valid, -1 otherwise.
 */
func (s Schematic) getValidPart(row, colStart, colEnd int) *Part {
	partNumber, err := strconv.Atoi(s[row][colStart:colEnd])
	if err != nil {
		log.Fatal(err)
	}

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

	var neighbouringStars []StarCoordinates

	if row > 0 {
		for j := j0; j < j1; j++ {
			ch := s[row-1][j : j+1]
			if ch != "." {
				isValidPart = true
			}
			if ch == "*" {
				neighbouringStars = append(neighbouringStars, StarCoordinates{row - 1, j})
			}
		}
	}

	if row < m-1 {
		for j := j0; j < j1; j++ {
			ch := s[row+1][j : j+1]
			if ch != "." {
				isValidPart = true
			}
			if ch == "*" {
				neighbouringStars = append(neighbouringStars, StarCoordinates{row + 1, j})
			}
		}

	}

	if colStart != 0 && s[row][colStart-1:colStart] != "." {
		isValidPart = true
	}

	if colStart != 0 && s[row][colStart-1:colStart] == "*" {
		neighbouringStars = append(neighbouringStars, StarCoordinates{row, colStart - 1})
	}

	if colEnd != n && s[row][colEnd:colEnd+1] != "." {
		isValidPart = true
	}

	if colEnd != n && s[row][colEnd:colEnd+1] == "*" {
		neighbouringStars = append(neighbouringStars, StarCoordinates{row, colEnd})
	}

	if !isValidPart {
		return nil
	}

	return &Part{partNumber, neighbouringStars}
}

func main() {
	s := readSchematic("input.txt")
	parts := s.findParts()

	sum := 0
	for i, p1 := range parts {
		for j := i + 1; j < len(parts); j++ {
			p2 := parts[j]
			coincidences := 0
			for _, s1 := range p1.NeighbouringStars {
				for _, s2 := range p2.NeighbouringStars {
					if s1.I == s2.I && s1.J == s2.J {
						coincidences++
					}
				}
			}
			if coincidences == 1 {
				sum += p1.Number * p2.Number
			}
		}
	}

	fmt.Println(sum)
}
