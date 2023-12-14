package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Scratchcard struct {
	WinningNumbers map[int]bool
	CardNumbers    []int
}

func readScratchcards(filename string) []Scratchcard {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var scratchcards []Scratchcard
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, ":")
		numbers := strings.Split(values[1], "|")

		s := Scratchcard{
			make(map[int]bool),
			make([]int, 0),
		}

		re := regexp.MustCompile(`\d+`)

		winningNumbersStr := numbers[0]
		winningMatches := re.FindAllStringIndex(winningNumbersStr, -1) // Find all matches
		if winningMatches == nil {
			log.Fatal("could not match winning numbers")
		}

		for _, i := range winningMatches {
			n, err := strconv.Atoi(winningNumbersStr[i[0]:i[1]])
			if err != nil {
				log.Fatal("could not convert number")
			}
			s.WinningNumbers[n] = true
		}

		cardNumbersStr := numbers[1]
		cardMatches := re.FindAllStringIndex(cardNumbersStr, -1) // Find all matches
		if cardMatches == nil {
			log.Fatal("could not match card numbers")
		}

		for _, i := range cardMatches {
			n, err := strconv.Atoi(cardNumbersStr[i[0]:i[1]])
			if err != nil {
				log.Fatal("could not convert number")
			}
			s.CardNumbers = append(s.CardNumbers, n)
		}

		scratchcards = append(scratchcards, s)
	}

	return scratchcards
}

func main() {
	scratchcards := readScratchcards("input.txt")

	sum := 0
	for _, s := range scratchcards {
		score := 0
		for _, n := range s.CardNumbers {
			_, present := s.WinningNumbers[n]
			if present {
				if score > 0 {
					score = 2 * score
				} else {
					score = 1
				}
			}
		}
		sum += score
	}

	fmt.Println(sum)
}
