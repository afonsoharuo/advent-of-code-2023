package main

import (
	"01/internal/digits"
	"bufio"
	"fmt"
	"log"
	"os"
)

func readInput(filename string) []int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var numbers []int
	for scanner.Scan() {
		line := scanner.Text()
		n, err := digits.ExtractNumber(line)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, n)
	}

	return numbers
}

func main() {
	numbers := readInput("input.txt")

	sum := 0
	for _, value := range numbers {
		sum += value
	}

	fmt.Println(sum)
}
