package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

type Queue[T any] struct {
	elems []T
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

type Node struct {
	pipe      rune
	edges     []*Node
	explored  bool
	component int
}

func newNode(pipe rune) *Node {
	n := Node{pipe, make([]*Node, 0), false, -1}
	return &n
}

func (node *Node) addEdge(edge *Node) {
	node.edges = append(node.edges, edge)
}

func (node *Node) bFS(component int) {
	node.explored = true
	node.component = component

	queue := Queue[*Node]{}
	queue.put(node)
	for !queue.empty() {
		v := queue.get()
		for _, w := range v.edges {
			if !w.explored {
				w.explored = true
				w.component = component
				queue.put(w)
			}
		}
	}
}

type Grid [][]*Node

func newGrid(gridMap []string) *Grid {
	m := len(gridMap)
	n := len(gridMap[0])

	// Initialise grid
	grid := make(Grid, m)
	for i, gridRow := range gridMap {
		grid[i] = make([]*Node, n)
		for j, pipe := range gridRow {
			grid[i][j] = newNode(pipe)
		}
	}

	grid.connectNodes()

	return &grid
}

func (grid *Grid) connectNodes() {
	// Assume all nodes have been initialised with respective pipes

	m := len(*grid)
	n := len((*grid)[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch (*grid)[i][j].pipe {
			case '|':
				grid.connectUpNode(i, j)
				grid.connectDownNode(i, j)
			case '-':
				grid.connectLeftNode(i, j)
				grid.connectRightNode(i, j)
			case 'L':
				grid.connectUpNode(i, j)
				grid.connectRightNode(i, j)
			case 'J':
				grid.connectUpNode(i, j)
				grid.connectLeftNode(i, j)
			case '7':
				grid.connectLeftNode(i, j)
				grid.connectDownNode(i, j)
			case 'F':
				grid.connectRightNode(i, j)
				grid.connectDownNode(i, j)
			case '.':
			case 'S':
				grid.connectUpNode(i, j)
				grid.connectDownNode(i, j)
				grid.connectLeftNode(i, j)
				grid.connectRightNode(i, j)
			}
		}
	}
}

func (grid *Grid) connectUpNode(i, j int) {
	if i > 0 {
		upPipe := (*grid)[i-1][j].pipe
		if upPipe == '|' || upPipe == '7' || upPipe == 'F' || upPipe == 'S' {
			(*grid)[i][j].addEdge((*grid)[i-1][j])
		}
	}
}

func (grid *Grid) connectDownNode(i, j int) {
	m := len(*grid)
	if i < m-1 {
		downPipe := (*grid)[i+1][j].pipe
		if downPipe == '|' || downPipe == 'L' || downPipe == 'J' || downPipe == 'S' {
			(*grid)[i][j].addEdge((*grid)[i+1][j])
		}
	}
}

func (grid *Grid) connectLeftNode(i, j int) {
	if j > 0 {
		leftPipe := (*grid)[i][j-1].pipe
		if leftPipe == '-' || leftPipe == 'L' || leftPipe == 'F' || leftPipe == 'S' {
			(*grid)[i][j].addEdge((*grid)[i][j-1])
		}
	}
}

func (grid *Grid) connectRightNode(i, j int) {
	n := len((*grid)[0])
	if j < n-1 {
		rightPipe := (*grid)[i][j+1].pipe
		if rightPipe == '-' || rightPipe == 'J' || rightPipe == '7' || rightPipe == 'S' {
			(*grid)[i][j].addEdge((*grid)[i][j+1])
		}
	}
}

func (grid *Grid) findConnectedComponents() [][]*Node {
	m := len(*grid)
	n := len((*grid)[0])

	component := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			node := (*grid)[i][j]
			if !node.explored {
				node.bFS(component)
				component++
			}
		}
	}

	components := make([][]*Node, component)
	for i, _ := range components {
		components[i] = make([]*Node, 0)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			node := (*grid)[i][j]
			components[node.component] = append(components[node.component], node)
		}
	}

	return components
}

func (grid *Grid) print() {
	for i, _ := range *grid {
		for j, _ := range (*grid)[i] {
			node := (*grid)[i][j]
			fmt.Println(string(node.pipe), node.explored, node.component)
		}
	}
}

func readGrid(filename string) *Grid {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read map
	gridMap := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		gridMap = append(gridMap, line)
	}

	grid := newGrid(gridMap)

	return grid
}

func main() {
	grid := readGrid("input.txt")
	components := grid.findConnectedComponents()

	var startComponent int
	for i, nodes := range components {
		for _, node := range nodes {
			if node.pipe == 'S' {
				startComponent = i
				break
			}
		}
	}

	fmt.Println(len(components[startComponent]) / 2)
}
