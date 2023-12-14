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
	cards []rune
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

func parseCards(cardsStr string) []rune {
	cards := make([]rune, 5)
	for i, card := range cardsStr {
		cards[i] = card
	}

	return cards
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
		cmp := compareCards(a.cards[i], b.cards[i])
		if cmp != 0 {
			return cmp
		}
	}

	return 0
}

// Assign a value to a type of hand for comparison.
func getTypeValue(h Hand) int {
	counts := make(map[rune]int)
	counts['J'] = 0
	for _, card := range h.cards {
		_, ok := counts[card]
		if !ok {
			counts[card] = 0
		}
		counts[card]++
	}
	numJokers := counts['J']

	hasFive := false
	hasFour := false
	hasThree := false
	hasFirstPair := false
	hasSecondPair := false
	for card, count := range counts {
		if card == 'J' {
			continue
		} else if count == 5 {
			hasFive = true
		} else if count == 4 {
			hasFour = true
		} else if count == 3 {
			hasThree = true
		} else if count == 2 {
			if !hasFirstPair {
				hasFirstPair = true
			} else {
				hasSecondPair = true
			}
		}
	}

	const fiveOfAKindScore int = 6
	const fourOfAKindScore int = 5
	const fullHouseScore int = 4
	const threeOfAKindScore int = 3
	const twoPairScore int = 2
	const onePairScore int = 1
	const highCardScore int = 0

	if hasFive {
		return fiveOfAKindScore
	} else if hasFour {
		if numJokers == 1 {
			return fiveOfAKindScore
		} else {
			return fourOfAKindScore
		}
	} else if hasThree && hasFirstPair {
		return fullHouseScore
	} else if hasThree {
		if numJokers == 2 {
			return fiveOfAKindScore
		} else if numJokers == 1 {
			return fourOfAKindScore
		} else {
			return threeOfAKindScore
		}
	} else if hasSecondPair {
		if numJokers == 1 {
			return fullHouseScore
		} else {
			return twoPairScore
		}
	} else if hasFirstPair {
		if numJokers == 3 {
			return fiveOfAKindScore
		} else if numJokers == 2 {
			return fourOfAKindScore
		} else if numJokers == 1 {
			return threeOfAKindScore
		} else {
			return onePairScore
		}
	} else {
		if numJokers >= 4 {
			return fiveOfAKindScore
		} else if numJokers == 3 {
			return fourOfAKindScore
		} else if numJokers == 2 {
			return threeOfAKindScore
		} else if numJokers == 1 {
			return onePairScore
		} else {
			return highCardScore
		}
	}
}

func compareCards(a, b rune) int {
	aValue := getCardValue(a)
	bValue := getCardValue(b)

	if aValue < bValue {
		return -1
	} else if aValue > bValue {
		return 1
	} else {
		return 0
	}
}

func getCardValue(card rune) int {
	var value int

	switch card {
	case 'A':
		value = 14
	case 'K':
		value = 13
	case 'Q':
		value = 12
	case 'T':
		value = 10
	case '9':
		value = 9
	case '8':
		value = 8
	case '7':
		value = 7
	case '6':
		value = 6
	case '5':
		value = 5
	case '4':
		value = 4
	case '3':
		value = 3
	case '2':
		value = 2
	case 'J':
		value = 1
	}

	return value
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
