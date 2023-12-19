package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Contraption struct {
	Grid [][]ContraptionElem
}

func newContraption() Contraption {
	grid := make([][]ContraptionElem, 0)
	return Contraption{grid}
}

func (c *Contraption) propagateBeam(i, j int, dir Direction) {
	m := len(c.Grid)
	n := len(c.Grid[0])

	for i >= 0 && i < m && j >= 0 && j < n {
		elem := &c.Grid[i][j]
		elem.Energised = true
		_, ok := elem.BeamDirections[dir]
		if !ok {
			elem.BeamDirections[dir] = true
		} else {
			// Beam has already been propagated in this direction through this element
			break
		}

		switch elem.Type {
		case Empty:
		case SWNEMirror:
			switch dir {
			case North:
				dir = East
			case West:
				dir = South
			case South:
				dir = West
			case East:
				dir = North
			}
		case NWSEMirror:
			switch dir {
			case North:
				dir = West
			case West:
				dir = North
			case South:
				dir = East
			case East:
				dir = South
			}
		case VertSplitter:
			if dir == West || dir == East {
				c.propagateBeam(i-1, j, North)
				c.propagateBeam(i+1, j, South)
				return
			}
		case HorzSplitter:
			if dir == North || dir == South {
				c.propagateBeam(i, j-1, West)
				c.propagateBeam(i, j+1, East)
				return
			}
		}

		switch dir {
		case North:
			i -= 1
		case West:
			j -= 1
		case South:
			i += 1
		case East:
			j += 1
		}
	}
}

func (c *Contraption) print() {
	for _, row := range c.Grid {
		for _, elem := range row {
			switch elem.Type {
			case Empty:
				fmt.Print(".")
			case SWNEMirror:
				fmt.Print("/")
			case NWSEMirror:
				fmt.Print("\\")
			case VertSplitter:
				fmt.Print("|")
			case HorzSplitter:
				fmt.Print("-")
			}
		}
		fmt.Println()
	}
}

func (c *Contraption) printEnergised() {
	for _, row := range c.Grid {
		for _, elem := range row {
			if elem.Energised {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (c *Contraption) countEnergised() int {
	numEnergised := 0
	for _, row := range c.Grid {
		for _, elem := range row {
			if elem.Energised {
				numEnergised += 1
			}
		}
	}

	return numEnergised
}

type ContraptionElemType int

type ContraptionElem struct {
	Type           ContraptionElemType
	Energised      bool
	BeamDirections map[Direction]bool
}

const (
	Empty ContraptionElemType = iota
	SWNEMirror
	NWSEMirror
	VertSplitter
	HorzSplitter
)

type Direction int

const (
	North Direction = iota
	West
	South
	East
)

func readLayout(filename string) Contraption {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	contraption := newContraption()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]ContraptionElem, len(line))
		for i, _ := range row {
			row[i].BeamDirections = make(map[Direction]bool)
		}
		for i, elem := range line {
			var t ContraptionElemType
			switch elem {
			case '.':
				t = Empty
			case '/':
				t = SWNEMirror
			case '\\':
				t = NWSEMirror
			case '|':
				t = VertSplitter
			case '-':
				t = HorzSplitter
			default:
				log.Panic("could not read element")
			}
			row[i].Type = t
		}
		contraption.Grid = append(contraption.Grid, row)
	}

	return contraption
}

func main() {
	contraption := readLayout("input.txt")
	contraption.propagateBeam(0, 0, East)
	numEnergised := contraption.countEnergised()
	fmt.Println(numEnergised)
}
