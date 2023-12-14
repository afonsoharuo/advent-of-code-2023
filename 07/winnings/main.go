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

type Hand struct {
	cards []int
	bid   int
}

func parseHands(filename string) []Hand {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	hands := make([]Hand, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		elems := strings.Split(line, " ")
		cardsStr := elems[0]
		bidStr := elems[1]

		cards := parseCards(cardsStr)

		bid, err := strconv.Atoi(string(bidStr))
		if err != nil {
			log.Fatal(err)
		}

		hands = append(hands, Hand{cards, bid})
	}

	return hands
}

func parseCards(cardsStr string) []int {
	cards := make([]int, 5)
	for i, card := range cardsStr {
		cards[i] = getCardAsInt(card)
	}

	return cards
}

func getCardAsInt(card rune) int {
	var cardInt int

	switch card {
	case 'A':
		cardInt = 14
	case 'K':
		cardInt = 13
	case 'Q':
		cardInt = 12
	case 'J':
		cardInt = 11
	case 'T':
		cardInt = 10
	default:
		var err error
		cardInt, err = strconv.Atoi(string(card))
		if err != nil {
			log.Fatal(err)
		}
	}

	return cardInt
}

func compareHands(a, b Hand) int {
	aValue := getTypeValue(a)
	bValue := getTypeValue(b)

	if aValue < bValue {
		return -1
	} else if aValue > bValue {
		return 1
	}

	// Values for types are equal, compare cards.
	for i, _ := range a.cards {
		if a.cards[i] < b.cards[i] {
			return -1
		} else if a.cards[i] > b.cards[i] {
			return 1
		}
	}

	return 0
}

// Assign a value to a type of hand for comparison.
func getTypeValue(h Hand) int {
	counts := make(map[int]int)
	for _, card := range h.cards {
		_, ok := counts[card]
		if !ok {
			counts[card] = 0
		}
		counts[card]++
	}

	hasFive := false
	hasFour := false
	hasThree := false
	hasFirstPair := false
	hasSecondPair := false

	for _, count := range counts {
		if count == 5 {
			hasFive = true
			break
		} else if count == 4 {
			hasFour = true
			break
		} else if count == 3 {
			hasThree = true
		} else if count == 2 {
			if !hasFirstPair {
				hasFirstPair = true
				continue
			} else {
				hasSecondPair = true
				break
			}
		}
	}

	if hasFive {
		return 6
	} else if hasFour {
		return 5
	} else if hasThree && hasFirstPair {
		return 4
	} else if hasThree {
		return 3
	} else if hasSecondPair {
		return 2
	} else if hasFirstPair {
		return 1
	} else {
		return 0
	}
}

func main() {
	hands := parseHands("input.txt")
	slices.SortFunc(hands, compareHands)

	winnings := 0
	for i, hand := range hands {
		winnings += (i + 1) * hand.bid
		fmt.Println(i, hand)
	}

	fmt.Println(winnings)
}
