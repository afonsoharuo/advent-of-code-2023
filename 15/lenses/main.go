package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Lens struct {
	Label       string
	FocalLength int
}

type Box []Lens

func readInitSeq(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	initSeq := strings.Split(line, ",")

	return initSeq
}

func calcHash(s string) int {
	value := 0
	for i := 0; i < len(s); i++ {
		value += int(s[i])
		value *= 17
		value %= 256
	}

	return value
}

func initialise(initSeq []string) []Box {
	boxes := make([]Box, 256)
	for i, _ := range boxes {
		boxes[i] = make([]Lens, 0)
	}

	for _, step := range initSeq {
		if strings.Contains(step, "-") {
			label := step[:len(step)-1]
			boxIndex := calcHash(label)
			lensIndex := findLens(boxes[boxIndex], label)
			if lensIndex > -1 {
				boxes[boxIndex] = slices.Delete(boxes[boxIndex], lensIndex, lensIndex+1)
			}
		} else if strings.Contains(step, "=") {
			elems := strings.Split(step, "=")
			label := elems[0]
			focalLength, err := strconv.Atoi(elems[1])
			if err != nil {
				log.Fatal(err)
			}

			lens := Lens{
				Label:       label,
				FocalLength: focalLength,
			}
			boxIndex := calcHash(label)
			lensIndex := findLens(boxes[boxIndex], label)
			if lensIndex > -1 {
				boxes[boxIndex][lensIndex] = lens
			} else {
				boxes[boxIndex] = append(boxes[boxIndex], lens)
			}
		}
	}

	return boxes
}

func findLens(lenses []Lens, label string) int {
	for i, lens := range lenses {
		if lens.Label == label {
			return i
		}
	}

	return -1
}

func calcFocusingPower(boxes []Box) int {
	sum := 0
	for i, box := range boxes {
		for j, lens := range box {
			sum += (i + 1) * (j + 1) * lens.FocalLength
		}
	}

	return sum
}

func main() {
	initSeq := readInitSeq("input.txt")
	boxes := initialise(initSeq)
	power := calcFocusingPower(boxes)
	fmt.Println(power)
}
