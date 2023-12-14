package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Race struct {
	time int
	dist int
}

func parsePaper(filename string) []Race {
	races := make([]Race, 0)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`\d+`)

	// Read time
	scanner.Scan()
	timeLine := scanner.Text()

	scanner.Scan()
	distLine := scanner.Text()

	timeMatches := re.FindAllString(timeLine, -1)
	distMatches := re.FindAllString(distLine, -1)

	n := len(timeMatches)
	for i := 0; i < n; i++ {
		time, err := strconv.Atoi(timeMatches[i])
		if err != nil {
			log.Fatal(err)
		}

		dist, err := strconv.Atoi(distMatches[i])
		if err != nil {
			log.Fatal(err)
		}

		races = append(races, Race{time, dist})
	}

	return races
}

func calcDist(pressedTime, totalTime int) int {
	moveTime := totalTime - pressedTime
	speed := pressedTime
	return speed * moveTime
}

func calcNumWinPossibilities(race Race) int {
	count := 0
	for t := 1; t < race.time; t++ {
		dist := calcDist(t, race.time)
		if dist > race.dist {
			count++
		}
	}
	return count
}

func main() {
	races := parsePaper("input.txt")

	winPossibilities := 1
	for _, race := range races {
		winPossibilities *= calcNumWinPossibilities(race)
	}

	fmt.Println(winPossibilities)
}
