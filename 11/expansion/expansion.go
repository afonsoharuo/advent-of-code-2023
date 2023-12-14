package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Space [][]bool

func (s *Space) calcDistSum(factor int) int {
	m := len(*s)
	n := len((*s)[0])

	points := make([]Point, 0)
	galaxyInRow := make([]bool, m)
	galaxyInCol := make([]bool, n)

	// Extract points and find empty rows and columns
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (*s)[i][j] {
				points = append(points, Point{i, j})
				galaxyInRow[i] = true
				galaxyInCol[j] = true
			}
		}
	}

	// Calculate sum of distances with expansion factor
	sum := 0
	for i, pa := range points {
		for _, pb := range points[i+1:] {
			d := dist(pa, pb)

			ia := min(pa.I, pb.I)
			ib := max(pa.I, pb.I)
			numEmptyRows := 0
			for k := ia + 1; k < ib; k++ {
				if !galaxyInRow[k] {
					numEmptyRows++
				}
			}

			ja := min(pa.J, pb.J)
			jb := max(pa.J, pb.J)
			numEmptyCols := 0
			for k := ja + 1; k < jb; k++ {
				if !galaxyInCol[k] {
					numEmptyCols++
				}
			}

			sum += d + (numEmptyRows+numEmptyCols)*(factor-1)
		}
	}

	return sum
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
	sum := space.calcDistSum(1000000)
	fmt.Println(sum)
}
