package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

type Space [][]bool

func (s *Space) expand() {
	m := len(*s)
	n := len((*s)[0])

	galaxyInRow := make([]bool, m)
	galaxyInCol := make([]bool, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (*s)[i][j] {
				galaxyInRow[i] = true
				galaxyInCol[j] = true
			}
		}
	}

	numRowInserts := 0
	for i, hasGalaxy := range galaxyInRow {
		if !hasGalaxy {
			emptyRow := make([]bool, n)
			*s = slices.Insert(*s, i+numRowInserts, emptyRow)
			numRowInserts++
		}
	}

	numColInserts := 0
	for j, hasGalaxy := range galaxyInCol {
		if !hasGalaxy {
			for i := 0; i < m+numRowInserts; i++ {
				(*s)[i] = slices.Insert((*s)[i], j+numColInserts, false)
			}
			numColInserts++
		}
	}
}

func (s *Space) print() {
	for _, row := range *s {
		for _, elem := range row {
			var repr rune
			if elem {
				repr = '#'
			} else {
				repr = '.'
			}
			fmt.Print(string(repr))
		}
		fmt.Println()
	}
}

type Point struct {
	I int
	J int
}

func dist(a, b Point) int {
	return abs(a.I-b.I) + abs(a.J-b.J)
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func extractPoints(s Space) []Point {
	points := make([]Point, 0)
	for i, row := range s {
		for j, elem := range row {
			if elem {
				points = append(points, Point{i, j})
			}
		}
	}

	return points
}

func readData(filename string) Space {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	space := make([][]bool, 0)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]bool, len(line))
		for j, elem := range line {
			switch elem {
			case '#':
				row[j] = true
			case '.':
				row[j] = false
			}
		}
		space = append(space, row)
	}

	return space
}

func main() {
	space := readData("input.txt")
	space.expand()

	points := extractPoints(space)

	sum := 0
	for i, p1 := range points {
		for _, p2 := range points[i+1:] {
			sum += dist(p1, p2)
		}
	}

	fmt.Println(sum)
}
