package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Rock int

const (
	Empty Rock = iota
	Cubic
	Round
)

type Platform [][]Rock

func (p *Platform) tiltNorth() {
	m := len(*p)
	n := len((*p)[0])

	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if (*p)[i][j] == Round {
				p.moveNorth(i, j)
			}
		}
	}
}

func (p *Platform) moveNorth(i, j int) {
	k := i - 1
	for k >= 0 && (*p)[k][j] == Empty {
		k--
	}
	// Set original position to empty first, as i and k+1 might be equal
	(*p)[i][j] = Empty
	(*p)[k+1][j] = Round
}

func (p *Platform) calcLoad() int {
	m := len(*p)
	n := len((*p)[0])

	load := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (*p)[i][j] == Round {
				load += m - i
			}
		}
	}

	return load
}

func (p *Platform) print() {
	for _, row := range *p {
		for _, elem := range row {
			switch elem {
			case Empty:
				fmt.Print(".")
			case Cubic:
				fmt.Print("#")
			case Round:
				fmt.Print("O")
			}
		}
		fmt.Println()
	}
}

func readInput(filename string) Platform {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	platform := make(Platform, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		row := make([]Rock, len(line))
		for i, elem := range line {
			var r Rock
			switch elem {
			case '.':
				r = Empty
			case '#':
				r = Cubic
			case 'O':
				r = Round
			default:
				log.Fatal("couldn't identify rock")
			}
			row[i] = r
		}
		platform = append(platform, row)
	}

	return platform
}

func main() {
	p := readInput("input.txt")
	p.tiltNorth()
	load := p.calcLoad()
	fmt.Println(load)
}
