package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

func replaceDigits(line string) string {
	re := regexp.MustCompile("one|two|three|four|five|six|seven|eight|nine")
	for {
		idx := re.FindStringIndex(line)
		if idx == nil {
			break
		}
		line = line[:idx[0]] + getDigit(line[idx[0]:idx[1]]) + line[idx[0]+1:]
	}

	return line
}

func getDigit(digit string) string {
	var r string
	switch digit {
	case "one":
		r = "1"
	case "two":
		r = "2"
	case "three":
		r = "3"
	case "four":
		r = "4"
	case "five":
		r = "5"
	case "six":
		r = "6"
	case "seven":
		r = "7"
	case "eight":
		r = "8"
	case "nine":
		r = "9"
	default:
		log.Fatal("No case found")
	}

	return r
}

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
		fmt.Println(line)
		line = replaceDigits(line)
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
		fmt.Println()
		numbers = append(numbers, n)
	}

	sum := 0
	for _, value := range numbers {
		sum += value
	}

	fmt.Println(sum)
}
