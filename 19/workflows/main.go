package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Workflow []func(Part) string

func (w *Workflow) process(p Part) string {
	var r string
	for _, f := range *w {
		r = f(p)
		if r != "" {
			return r
		}
	}

	return ""
}

type Part struct {
	x int
	m int
	a int
	s int
}

func readInput(filename string) (map[string]Workflow, []Part) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read workflows
	workflows := make(map[string]Workflow)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			// Finished reading workflows
			break
		}

		topElems := strings.Split(line, "{")
		name := topElems[0]

		elems := strings.Split(topElems[1][:len(topElems[1])-1], ",")
		workflow := make(Workflow, 0)
		for _, elem := range elems {
			parts := strings.Split(elem, ":")
			if len(parts) > 1 {
				condition := parts[0]
				output := parts[1]
				var condElems []string
				var comparison int
				if strings.Contains(condition, "<") {
					condElems = strings.Split(condition, "<")
					comparison = -1
				} else if strings.Contains(condition, ">") {
					condElems = strings.Split(condition, ">")
					comparison = 1
				} else {
					log.Fatal("cannot read condition")
				}

				f := func(p Part) string {
					var value int
					var passed bool

					switch condElems[0] {
					case "x":
						value = p.x
					case "m":
						value = p.m
					case "a":
						value = p.a
					case "s":
						value = p.s
					}

					condValue, err := strconv.Atoi(condElems[1])
					if err != nil {
						log.Fatal(err)
					}

					if comparison < 0 {
						passed = value < condValue
					} else {
						passed = value > condValue
					}

					if passed {
						return output
					} else {
						return ""
					}
				}

				workflow = append(workflow, f)
			} else {
				workflow = append(workflow, func(p Part) string { return parts[0] })
			}
		}
		workflows[name] = workflow
	}

	// Read parts
	parts := make([]Part, 0)
	for scanner.Scan() {
		line := scanner.Text()
		trimmedLine := line[1 : len(line)-1]

		ps := strings.Split(trimmedLine, ",")
		partValues := make([]int, 0)
		for _, p := range ps {
			valueElems := strings.Split(p, "=")
			value, err := strconv.Atoi(valueElems[1])
			if err != nil {
				log.Fatal(err)
			}

			partValues = append(partValues, value)
		}

		part := Part{
			x: partValues[0],
			m: partValues[1],
			a: partValues[2],
			s: partValues[3],
		}
		parts = append(parts, part)
	}

	return workflows, parts
}

func findAccepted(workflows map[string]Workflow, parts []Part) []Part {
	accepted := make([]Part, 0)

	for _, part := range parts {
		curWorkflow := "in"
		for {
			workflow := workflows[curWorkflow]
			r := workflow.process(part)
			if r == "A" {
				accepted = append(accepted, part)
				break
			} else if r == "R" {
				break
			} else if r != "" {
				curWorkflow = r
			} else {
				log.Fatal("failed to process part", part)
			}
		}
	}

	return accepted
}

func sumParts(parts []Part) int {
	sum := 0
	for _, part := range parts {
		sum += part.x
		sum += part.m
		sum += part.a
		sum += part.s
	}

	return sum
}

func main() {
	workflows, parts := readInput("input.txt")
	accepted := findAccepted(workflows, parts)
	sum := sumParts(accepted)
	fmt.Println(sum)
}
