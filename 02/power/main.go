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

type CubeSet struct {
	Red   int
	Green int
	Blue  int
}

type Game struct {
	Id       int
	CubeSets []CubeSet
}

func parseGames(filename string) []Game {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var games []Game
	for scanner.Scan() {
		line := scanner.Text()

		game := Game{}

		reGame := regexp.MustCompile(`\d+`)
		match := reGame.FindString(line)
		id, err := strconv.Atoi(match)
		if err != nil {
			log.Fatal(err)
		}
		game.Id = id

		setStr := strings.Split(line, ":")
		sets := strings.Split(setStr[1], ";")

		for _, set := range sets {
			cubeSet := CubeSet{}

			reRed := regexp.MustCompile(`(\d+) red`)
			redSubMatches := reRed.FindStringSubmatch(set)
			if redSubMatches != nil {
				n, err := strconv.Atoi(redSubMatches[1])
				if err != nil {
					log.Fatal(err)
				}
				cubeSet.Red = n
			}

			reGreen := regexp.MustCompile(`(\d+) green`)
			greenSubMatches := reGreen.FindStringSubmatch(set)
			if greenSubMatches != nil {
				n, err := strconv.Atoi(greenSubMatches[1])
				if err != nil {
					log.Fatal(err)
				}
				cubeSet.Green = n
			}

			reBlue := regexp.MustCompile(`(\d+) blue`)
			blueSubMatches := reBlue.FindStringSubmatch(set)
			if blueSubMatches != nil {
				n, err := strconv.Atoi(blueSubMatches[1])
				if err != nil {
					log.Fatal(err)
				}
				cubeSet.Blue = n
			}

			game.CubeSets = append(game.CubeSets, cubeSet)
		}

		games = append(games, game)
	}

	return games
}

func main() {
	games := parseGames("input.txt")

	sum := 0
	for _, game := range games {
		var minCubeSet CubeSet
		for _, cubeSet := range game.CubeSets {
			if cubeSet.Red > minCubeSet.Red {
				minCubeSet.Red = cubeSet.Red
			}
			if cubeSet.Green > minCubeSet.Green {
				minCubeSet.Green = cubeSet.Green
			}
			if cubeSet.Blue > minCubeSet.Blue {
				minCubeSet.Blue = cubeSet.Blue
			}
		}

		sum += minCubeSet.Red * minCubeSet.Green * minCubeSet.Blue
	}

	fmt.Println(sum)
}
