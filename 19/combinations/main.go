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

type Queue[T any] struct {
	elems []T
}

func (q *Queue[T]) size() int {
	return len(q.elems)
}

func (q *Queue[T]) empty() bool {
	return len(q.elems) == 0
}

func (q *Queue[T]) put(elem T) {
	q.elems = append(q.elems, elem)
}

func (q *Queue[T]) get() T {
	elem := q.elems[0]
	q.elems = slices.Delete(q.elems, 0, 1)
	return elem
}

type Comparison int

const (
	LessThan Comparison = iota
	GreaterThan
)

type WorkflowResult struct {
	part   Part
	result string
}

type Workflow []func(Part) []WorkflowResult

func (w *Workflow) process(p Part) []WorkflowResult {
	queue := Queue[Part]{}
	queue.put(p)
	wrs := make([]WorkflowResult, 0)
	for _, f := range *w {
		n := queue.size()
		for i := 0; i < n; i++ {
			curPart := queue.get()
			curWrs := f(curPart)
			for _, wr := range curWrs {
				if wr.result != "" {
					wrs = append(wrs, wr)
				} else {
					queue.put(wr.part)
				}
			}
		}
	}

	if !queue.empty() {
		log.Fatal("did not process all elements in queue")
	}

	return wrs
}

type Part [4][2]int

func readInput(filename string) map[string]Workflow {
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
				var comparison Comparison
				if strings.Contains(condition, "<") {
					condElems = strings.Split(condition, "<")
					comparison = LessThan
				} else if strings.Contains(condition, ">") {
					condElems = strings.Split(condition, ">")
					comparison = GreaterThan
				} else {
					log.Fatal("cannot read condition")
				}

				f := func(p Part) []WorkflowResult {
					result := make([]WorkflowResult, 0)

					var index int
					switch condElems[0] {
					case "x":
						index = 0
					case "m":
						index = 1
					case "a":
						index = 2
					case "s":
						index = 3
					}

					condValue, err := strconv.Atoi(condElems[1])
					if err != nil {
						log.Fatal(err)
					}

					if comparison == LessThan {
						if p[index][1] < condValue {
							result = append(result, WorkflowResult{p, output})
						} else if p[index][0] < condValue {
							var lowPart, highPart Part
							for i := 0; i < 4; i++ {
								if i == index {
									lowPart[i][0] = p[i][0]
									lowPart[i][1] = condValue - 1
									highPart[i][0] = condValue
									highPart[i][1] = p[i][1]
								} else {
									lowPart[i] = p[i]
									highPart[i] = p[i]
								}
							}
							result = append(result, WorkflowResult{lowPart, output})
							result = append(result, WorkflowResult{highPart, ""})
						} else {
							result = append(result, WorkflowResult{p, ""})
						}
					} else {
						if p[index][0] > condValue {
							result = append(result, WorkflowResult{p, output})
						} else if p[index][1] > condValue {
							var lowPart, highPart Part
							for i := 0; i < 4; i++ {
								if i == index {
									lowPart[i][0] = p[i][0]
									lowPart[i][1] = condValue
									highPart[i][0] = condValue + 1
									highPart[i][1] = p[i][1]
								} else {
									lowPart[i] = p[i]
									highPart[i] = p[i]
								}
							}
							result = append(result, WorkflowResult{lowPart, ""})
							result = append(result, WorkflowResult{highPart, output})
						} else {
							result = append(result, WorkflowResult{p, ""})
						}
					}

					return result
				}

				workflow = append(workflow, f)
			} else {
				f := func(p Part) []WorkflowResult {
					result := make([]WorkflowResult, 1)
					result[0] = WorkflowResult{p, parts[0]}
					return result
				}
				workflow = append(workflow, f)
			}
		}
		workflows[name] = workflow
	}

	return workflows
}

func findCombinations(workflows map[string]Workflow) int {
	initialPart := Part{
		[2]int{1, 4000},
		[2]int{1, 4000},
		[2]int{1, 4000},
		[2]int{1, 4000},
	}

	workflowResult := WorkflowResult{initialPart, "in"}

	combinations := 0
	queue := Queue[WorkflowResult]{}
	queue.put(workflowResult)
	for !queue.empty() {
		wr := queue.get()
		w := workflows[wr.result]
		newWrs := w.process(wr.part)
		for _, newWr := range newWrs {
			if newWr.result == "A" {
				combinations += calcCombinations(newWr.part)
			} else if newWr.result != "" && newWr.result != "R" {
				queue.put(newWr)
			} else if newWr.result != "R" {
				log.Fatal("failed to process")
			}
		}
	}

	return combinations
}

func calcCombinations(p Part) int {
	combinations := 1
	for i := 0; i < 4; i++ {
		combinations *= p[i][1] - p[i][0] + 1
	}

	return combinations
}

func main() {
	workflows := readInput("input.txt")
	combinations := findCombinations(workflows)
	fmt.Println(combinations)
}
