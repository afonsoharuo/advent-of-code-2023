package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readSensor(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sensorReadings := make([][]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		readingsStr := strings.Split(line, " ")

		sequence := make([]int, 0)
		for _, r := range readingsStr {
			value, err := strconv.Atoi(r)
			if err != nil {
				log.Fatal(err)
			}
			sequence = append(sequence, value)
		}
		sensorReadings = append(sensorReadings, sequence)
	}

	return sensorReadings
}

func extrapolateBack(seq []int) int {
	diff := calcDiff(seq)
	if isAllZeros(diff) {
		return seq[0]
	} else {
		return seq[0] - extrapolateBack(diff)
	}
}

func calcDiff(seq []int) []int {
	n := len(seq)
	diff := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		diff[i] = seq[i+1] - seq[i]
	}
	return diff
}

func isAllZeros(seq []int) bool {
	for i, _ := range seq {
		if seq[i] != 0 {
			return false
		}
	}

	return true
}

func main() {
	readings := readSensor("input.txt")

	sum := 0
	for _, r := range readings {
		sum += extrapolateBack(r)
	}
	fmt.Println(sum)
}
