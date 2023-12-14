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

func parseAlmanac(filename string) ([][]int, map[string]*Map) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`\d+`)

	// Read seed ranges
	scanner.Scan()
	line := scanner.Text()
	matches := re.FindAllString(line, -1)
	if matches == nil {
		log.Fatal("Did not match any seeds")
	}

	seeds := make([][]int, len(matches)/2)
	for i := 0; i < len(seeds); i++ {
		seeds[i] = make([]int, 2)
	}
	for i := 0; i < len(matches); i += 2 {
		nStart, err := strconv.Atoi(matches[i])
		if err != nil {
			log.Fatal(err)
		}
		nCount, err := strconv.Atoi(matches[i+1])
		if err != nil {
			log.Fatal(err)
		}
		seeds[i/2][0] = nStart
		seeds[i/2][1] = nStart + nCount
	}

	// Skip empty line
	scanner.Scan()

	// Read maps
	maps := make(map[string]*Map)
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

func findRangeMin(seedRange []int, maps map[string]*Map, c chan int) {
	minLocation := math.MaxInt

	for seed := seedRange[0]; seed <= seedRange[1]; seed++ {
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

	c <- minLocation
}

func main() {
	seedRanges, maps := parseAlmanac("input.txt")

	c := make(chan int)
	for _, seedRange := range seedRanges {
		go findRangeMin(seedRange, maps, c)
	}

	n := len(seedRanges)
	mins := make([]int, n)
	for i := 0; i < n; i++ {
		m := <-c
		mins[i] = m
	}

	minLocation := math.MaxInt
	for _, m := range mins {
		if m < minLocation {
			minLocation = m
		}
	}
	fmt.Println(minLocation)
}
