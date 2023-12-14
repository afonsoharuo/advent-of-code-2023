package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

type Direction int

const (
	Invalid Direction = iota
	North
	East
	South
	West
)

type Side int

const (
	Left Side = iota
	Right
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
	i         int
	j         int
}

func newNode(pipe rune, i, j int) *Node {
	n := Node{pipe, make([]*Node, 0), false, -1, i, j}
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
			grid[i][j] = newNode(pipe, i, j)
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

func (grid *Grid) connectAllExceptLoop(loop []*Node) {
	m := len(*grid)
	n := len((*grid)[0])

	for i, _ := range *grid {
		for j, _ := range (*grid)[i] {
			node := (*grid)[i][j]

			if nodeInLoop(node, loop) {
				continue
			}

			// Connect up
			if i > 0 {
				upNode := (*grid)[i-1][j]
				if !nodeInLoop(upNode, loop) {
					node.addEdge(upNode)
				}
			}

			// Connect down
			if i < m-1 {
				downNode := (*grid)[i+1][j]
				if !nodeInLoop(downNode, loop) {
					node.addEdge(downNode)
				}
			}

			// Connect left
			if j > 0 {
				leftNode := (*grid)[i][j-1]
				if !nodeInLoop(leftNode, loop) {
					node.addEdge(leftNode)
				}
			}

			// Connect right
			if j < n-1 {
				rightNode := (*grid)[i][j+1]
				if !nodeInLoop(rightNode, loop) {
					node.addEdge(rightNode)
				}
			}
		}
	}
}

func nodeInLoop(node *Node, loop []*Node) bool {
	for _, loopNode := range loop {
		if node == loopNode {
			return true
		}
	}
	return false
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

func countEnclosed(grid *Grid, loop []*Node) int {
	m := len(*grid)
	n := len((*grid)[0])

	// Reset explored
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			(*grid)[i][j].explored = false
			(*grid)[i][j].component = -1
		}
	}

	// Find inside of loop
	positions := make([][]int, 0)
	node := loop[0]
	numExplored := 0
	loopSize := len(loop)
	for numExplored != loopSize {
		p := make([]int, 2)
		p[0] = node.i
		p[1] = node.j
		positions = append(positions, p)
		node.explored = true
		numExplored++
		for _, edge := range node.edges {
			if !edge.explored {
				node = edge
				break
			}
		}
	}

	nPos := len(positions)
	leftTurns := 0
	rightTurns := 0
	lastPos := positions[nPos-1]
	lastDirection := getDirection(positions[nPos-2], positions[nPos-1])
	for i := 0; i < len(positions); i++ {
		curPos := positions[i]
		curDirection := getDirection(lastPos, curPos)
		if (lastDirection == North && curDirection == East) || (lastDirection == East && curDirection == South) || (lastDirection == South && curDirection == West) || (lastDirection == West && curDirection == North) {
			rightTurns++
		} else if (lastDirection == North && curDirection == West) || (lastDirection == West && curDirection == South) || (lastDirection == South && curDirection == East) || (lastDirection == East && curDirection == North) {
			leftTurns++
		}
		lastPos = curPos
		lastDirection = curDirection
	}

	var side Side
	if leftTurns > rightTurns {
		side = Left
	} else {
		side = Right
	}

	// Remake connections
	grid.connectAllExceptLoop(loop)

	components := grid.findEnclosedConnected(positions, side)

	count := 0
	for _, component := range components {
		count += len(component)
	}

	return count
}

func getDirection(lastPos, curPos []int) Direction {
	if curPos[0] == lastPos[0] {
		if curPos[1] > lastPos[1] {
			return East
		} else if curPos[1] < lastPos[1] {
			return West
		}
	} else if curPos[1] == lastPos[1] {
		if curPos[0] > lastPos[0] {
			return South
		} else if curPos[0] < lastPos[0] {
			return North
		}
	}
	return Invalid
}

func (grid *Grid) findEnclosedConnected(positions [][]int, side Side) [][]*Node {
	m := len(*grid)
	n := len((*grid)[0])

	component := 0
	nPos := len(positions)
	lastPos := positions[nPos-1]
	for i := 0; i < len(positions); i++ {
		curPos := positions[i]
		direction := getDirection(lastPos, curPos)
		startPos := make([]int, 2)
		if direction == North {
			if side == Left {
				startPos[0] = curPos[0]
				startPos[1] = curPos[1] - 1
			} else {
				startPos[0] = curPos[0]
				startPos[1] = curPos[1] + 1
			}
		} else if direction == West {
			if side == Left {
				startPos[0] = curPos[0] + 1
				startPos[1] = curPos[1]
			} else {
				startPos[0] = curPos[0] - 1
				startPos[1] = curPos[1]
			}
		} else if direction == South {
			if side == Left {
				startPos[0] = curPos[0]
				startPos[1] = curPos[1] + 1
			} else {
				startPos[0] = curPos[0]
				startPos[1] = curPos[1] - 1
			}
		} else if direction == East {
			if side == Left {
				startPos[0] = curPos[0] - 1
				startPos[1] = curPos[1]
			} else {
				startPos[0] = curPos[0] + 1
				startPos[1] = curPos[1]
			}
		} else {
			log.Fatal("invalid direction")
		}

		startNode := (*grid)[startPos[0]][startPos[1]]
		if !startNode.explored {
			startNode.bFS(component)
			component++
		}

		lastPos = curPos
	}

	components := make([][]*Node, component)
	for i, _ := range components {
		components[i] = make([]*Node, 0)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			node := (*grid)[i][j]
			if node.component != -1 {
				components[node.component] = append(components[node.component], node)
			}
		}
	}

	return components
}

func main() {
	grid := readGrid("complexinput.txt")
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

	numEnclosed := countEnclosed(grid, components[startComponent])

	for i := 0; i < len(*grid); i++ {
		for j := 0; j < len((*grid)[i]); j++ {
			node := (*grid)[i][j]
			if nodeInLoop(node, components[startComponent]) {
				switch node.pipe {
				case '|':
					node.pipe = '║'
				case '-':
					node.pipe = '═'
				case 'L':
					node.pipe = '╚'
				case 'J':
					node.pipe = '╝'
				case '7':
					node.pipe = '╗'
				case 'F':
					node.pipe = '╔'
				}
			} else if node.component != -1 {
				node.pipe = 'I'
			} else {
				node.pipe = 'O'
			}
			fmt.Print(string(node.pipe))
		}
		fmt.Println()
	}

	fmt.Println(numEnclosed)
}
