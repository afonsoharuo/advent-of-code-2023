package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type MapRange struct {
	start  int
	end    int
	offset int
}

func (r MapRange) inRange(item int) bool {
	return r.start <= item && item <= r.end
}

func (r MapRange) convert(item int) int {
	return item + r.offset
}

type Map struct {
	ranges []MapRange
}

func NewMap() *Map {
	return &Map{make([]MapRange, 0)}
}

func (m *Map) addRange(r MapRange) {
	m.ranges = append(m.ranges, r)
}

func (m *Map) convert(item int) int {
	for _, r := range m.ranges {
		if r.inRange(item) {
			return r.convert(item)
		}
	}

	// Not in any range
	return item
}

func parseAlmanac(filename string) ([]int, map[string]*Map) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	seeds := make([]int, 0)
	maps := make(map[string]*Map)

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`\d+`)

	// Read seeds
	scanner.Scan()
	line := scanner.Text()
	matches := re.FindAllString(line, -1)
	if matches == nil {
		log.Fatal("Did not match any seeds")
	}

	for _, s := range matches {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		seeds = append(seeds, n)
	}

	// Skip empty line
	scanner.Scan()

	// Read maps
	type Mode int
	const (
		Title Mode = iota
		Range
	)
	var mode Mode = Title
	currentTitle := ""
	for scanner.Scan() {
		line := scanner.Text()

		switch mode {
		case Title:
			titleLineElems := strings.Split(line, " ")
			currentTitle = titleLineElems[0]
			maps[currentTitle] = NewMap()
			mode = Range
		case Range:
			if line == "" {
				mode = Title
				continue
			}

			rangeLineElems := strings.Split(line, " ")
			ns := make([]int, 0)
			for _, nStr := range rangeLineElems {
				n, err := strconv.Atoi(nStr)
				if err != nil {
					log.Fatal(err)
				}
				ns = append(ns, n)
			}

			r := MapRange{
				ns[1],
				ns[1] + ns[2] - 1,
				ns[0] - ns[1],
			}
			maps[currentTitle].addRange(r)
		}
	}

	return seeds, maps
}

func main() {
	seeds, maps := parseAlmanac("input.txt")

	minLocation := math.MaxInt
	for _, seed := range seeds {
		soil := maps["seed-to-soil"].convert(seed)
		fertilizer := maps["soil-to-fertilizer"].convert(soil)
		water := maps["fertilizer-to-water"].convert(fertilizer)
		light := maps["water-to-light"].convert(water)
		temperature := maps["light-to-temperature"].convert(light)
		humidity := maps["temperature-to-humidity"].convert(temperature)
		location := maps["humidity-to-location"].convert(humidity)

		if location < minLocation {
			minLocation = location
		}
	}

	fmt.Println(minLocation)
}
