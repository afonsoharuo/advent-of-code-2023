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

outerLoop:
	for k := 1; k < n; k++ {
		i := k - 1
		j := k
		numSmudges := 0
		for i >= 0 && j < n {
			rowsEqual, hasSmudge := m.areColsEqual(i, j)
			if hasSmudge {
				numSmudges++
			}
			if !rowsEqual || numSmudges > 1 {
				continue outerLoop
			}

			i--
			j++
		}

		if numSmudges == 1 {
			return k
		}
	}

	return -1
}

/*
 * Test for equality allowing up to one smudge, returning equality and presence of smudge.
 */
func (m FloorMap) areColsEqual(i, j int) (bool, bool) {
	n := len(m)

	nSmudges := 0
	for k := 0; k < n; k++ {
		if m[k][i] != m[k][j] {
			nSmudges++
		}
	}

	if nSmudges > 1 {
		return false, false
	} else if nSmudges == 1 {
		return true, true
	} else {
		return true, false
	}
}

/*
 * Return greatest index before first row reflection, or -1 if not found.
 */
func (m FloorMap) findRowReflection() int {
	n := len(m)

outerLoop:
	for k := 1; k < n; k++ {
		i := k - 1
		j := k
		numSmudges := 0
		for i >= 0 && j < n {
			rowsEqual, hasSmudge := m.areRowsEqual(i, j)
			if hasSmudge {
				numSmudges++
			}
			if !rowsEqual || numSmudges > 1 {
				continue outerLoop
			}

			i--
			j++
		}

		if numSmudges == 1 {
			return k
		}
	}

	return -1
}

/*
 * Test for equality allowing up to one smudge, returning equality and presence of smudge.
 */
func (m FloorMap) areRowsEqual(i, j int) (bool, bool) {
	n := len(m[0])

	nSmudges := 0
	for k := 0; k < n; k++ {
		if m[i][k] != m[j][k] {
			nSmudges++
		}
	}

	if nSmudges > 1 {
		return false, false
	} else if nSmudges == 1 {
		return true, true
	} else {
		return true, false
	}
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
