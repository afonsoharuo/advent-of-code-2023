package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var numbers []int
	for scanner.Scan() {
		line := scanner.Text()
		var digits []string
		for _, ch := range line {
			if unicode.IsDigit(ch) {
				digits = append(
					digits,
					string(ch),
				)
			}
		}
		nStr := digits[0] + digits[len(digits)-1]
		n, err := strconv.Atoi(nStr)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(line)
		fmt.Println(digits)
		fmt.Println(n)
		numbers = append(numbers, n)
	}

	sum := 0
	for _, value := range numbers {
		sum += value
	}

	fmt.Println(sum)
}
