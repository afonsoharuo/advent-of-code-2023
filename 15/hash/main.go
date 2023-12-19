package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readInitSeq(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	initSeq := strings.Split(line, ",")

	return initSeq
}

func calcHashSum(seq []string) int {
	sum := 0

	for _, s := range seq {
		sum += calcHash(s)
	}

	return sum
}

func calcHash(s string) int {
	value := 0
	for i := 0; i < len(s); i++ {
		value += int(s[i])
		value *= 17
		value %= 256
	}

	return value
}

func main() {
	initSeq := readInitSeq("input.txt")
	sum := calcHashSum(initSeq)
	fmt.Println(sum)
}
