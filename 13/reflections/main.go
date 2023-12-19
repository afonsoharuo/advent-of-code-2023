package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Floor int

const (
	Ash Floor = iota
	Rock
)

type FloorMap [][]Floor

func (m FloorMap) print() {
	for _, row := range m {
		for _, floor := range row {
			switch floor {
			case Ash:
				fmt.Print(".")
			case Rock:
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

func readNotes(filename string) []FloorMap {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var floorMap FloorMap
	floorMap = make([][]Floor, 0)
	floorMaps := make([]FloorMap, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) > 0 {
			row := make([]Floor, len(line))
			for i, ch := range line {
				var floor Floor
				switch ch {
				case '.':
					floor = Ash
				case '#':
					floor = Rock
				default:
					log.Fatal("invalid map element")
				}
				row[i] = floor
			}
			floorMap = append(floorMap, row)
		} else {
			// Finished reading current map
			floorMaps = append(floorMaps, floorMap)
			floorMap = make([][]Floor, 0)
		}
	}

	// Append last map
	floorMaps = append(floorMaps, floorMap)

	return floorMaps
}

/*
 * Return greatest index before first column reflection, or -1 if not found.
 */
func (m FloorMap) findColReflection() int {
	n := len(m[0])

	colsEqual := make([][]bool, n)
	for i := 0; i < n; i++ {
		colsEqual[i] = make([]bool, n)
	}

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			colsEqual[i][j] = m.areColsEqual(i, j)
		}
	}

outerLoop:
	for k := 1; k < n; k++ {
		i := k - 1
		j := k
		for i >= 0 && j < n {
			if !colsEqual[i][j] {
				continue outerLoop
			}

			i--
			j++
		}

		return k
	}

	return -1
}

func (m FloorMap) areColsEqual(i, j int) bool {
	n := len(m)

	for k := 0; k < n; k++ {
		if m[k][i] != m[k][j] {
			return false
		}
	}

	return true
}

/*
 * Return greatest index before first row reflection, or -1 if not found.
 */
func (m FloorMap) findRowReflection() int {
	n := len(m)

	rowsEqual := make([][]bool, n)
	for i := 0; i < n; i++ {
		rowsEqual[i] = make([]bool, n)
	}

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			rowsEqual[i][j] = m.areRowsEqual(i, j)
		}
	}

outerLoop:
	for k := 1; k < n; k++ {
		i := k - 1
		j := k
		for i >= 0 && j < n {
			if !rowsEqual[i][j] {
				continue outerLoop
			}

			i--
			j++
		}

		return k
	}

	return -1
}

func (m FloorMap) areRowsEqual(i, j int) bool {
	n := len(m[0])

	for k := 0; k < n; k++ {
		if m[i][k] != m[j][k] {
			return false
		}
	}

	return true
}

func (m FloorMap) summarise() int {
	const rowFactor = 100
	var r int

	r = m.findColReflection()
	if r != -1 {
		return r
	}

	r = m.findRowReflection()
	if r != -1 {
		return rowFactor * r
	}

	return -1
}

func summariseNotes(maps []FloorMap) int {
	sum := 0
	for _, m := range maps {
		r := m.summarise()
		if r == -1 {
			log.Fatal("didn't find reflection")
		}

		sum += r
	}

	return sum
}

func main() {
	maps := readNotes("input.txt")
	sum := summariseNotes(maps)
	fmt.Println(sum)
}
