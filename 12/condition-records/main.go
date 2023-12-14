package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	Operational int = iota
	Damaged
	Unknown
)

type Record struct {
	Springs []int
	Damaged []int
}

func readRecords(filename string) []Record {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	records := make([]Record, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lineElems := strings.Split(line, " ")

		springs := make([]int, len(lineElems[0]))
		for i, elem := range lineElems[0] {
			var cond int
			switch elem {
			case '.':
				cond = Operational
			case '#':
				cond = Damaged
			case '?':
				cond = Unknown
			default:
				log.Fatal("invalid spring condition")
			}
			springs[i] = cond
		}

		damaged := make([]int, 0)
		damagedElems := strings.Split(lineElems[1], ",")
		for _, valueStr := range damagedElems {
			value, err := strconv.Atoi(valueStr)
			if err != nil {
				log.Fatal(err)
			}
			damaged = append(damaged, value)
		}

		records = append(records, Record{springs, damaged})
	}

	return records
}

func calcArrangements(records []Record) int {
	validArrangements := 0

	for _, record := range records {
		unknowns := findUnknowns(record.Springs)
		n := 1 << len(unknowns)

		arrangement := make([]int, len(record.Springs))
		for i := 0; i < n; i++ {
			seenUnknowns := 0
			for j, cond := range record.Springs {
				if cond == Unknown {
					arrangement[j] = (i >> seenUnknowns) % 2
					seenUnknowns++
				} else {
					arrangement[j] = cond
				}
			}
			if isValidArrangement(arrangement, record.Damaged) {
				validArrangements++
			}
		}
	}

	return validArrangements
}

func isValidArrangement(arrangement []int, expectedDamaged []int) bool {
	damaged := make([]int, 0)

	contiguousDamaged := 0
	lastValue := Operational
	for _, value := range arrangement {
		if value == Damaged {
			contiguousDamaged++
		} else if lastValue == Damaged && value == Operational {
			damaged = append(damaged, contiguousDamaged)
			contiguousDamaged = 0
		}

		lastValue = value
	}

	if lastValue == Damaged {
		damaged = append(damaged, contiguousDamaged)
	}

	if len(damaged) != len(expectedDamaged) {
		return false
	} else {
		for i := 0; i < len(damaged); i++ {
			if damaged[i] != expectedDamaged[i] {
				return false
			}
		}

		return true
	}
}

func findUnknowns(springs []int) []int {
	unknowns := make([]int, 0)
	for i, s := range springs {
		if s == Unknown {
			unknowns = append(unknowns, i)
		}
	}

	return unknowns
}

func main() {
	records := readRecords("input.txt")
	n := calcArrangements(records)
	fmt.Println(n)
}
