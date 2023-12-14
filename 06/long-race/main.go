package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	time int
	dist int
}

func (race Race) calcDist(time int) int {
	moveTime := race.time - time
	speed := time
	return speed * moveTime
}

func (race Race) findFirstWin(initialTime, finalTime int) int {
	time := initialTime + (finalTime-initialTime)/2
	dist := race.calcDist(time)

	if dist > race.dist {
		// Victory, recurse on left half
		return race.findFirstWin(initialTime, time)
	} else {
		// Loss
		if race.calcDist(time+1) > race.dist {
			// Found first win
			return time + 1
		} else {
			// Also a loss, recurse on right half
			return race.findFirstWin(time, finalTime)
		}
	}
}

func parsePaper(filename string) Race {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read time
	scanner.Scan()
	timeLine := scanner.Text()
	timeWithSpaces := strings.Split(timeLine, ":")[1]
	timeStr := strings.ReplaceAll(timeWithSpaces, " ", "")

	// Read distance
	scanner.Scan()
	distLine := scanner.Text()
	distWithSpaces := strings.Split(distLine, ":")[1]
	distStr := strings.ReplaceAll(distWithSpaces, " ", "")

	time, err := strconv.Atoi(timeStr)
	if err != nil {
		log.Fatal(err)
	}

	dist, err := strconv.Atoi(distStr)
	if err != nil {
		log.Fatal(err)
	}

	return Race{time, dist}
}

func main() {
	race := parsePaper("input.txt")
	firstWin := race.findFirstWin(1, race.time/2+1)
	numPossibilities := race.time - 1
	numLossPossibilities := 2 * (firstWin - 1)
	numWinPossibilities := numPossibilities - numLossPossibilities
	fmt.Println(numWinPossibilities)
}
