package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type Directions []rune

type Node struct {
	label string
	left  *Node
	right *Node
}

func readMap(filename string) (Directions, []*Node) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read directions
	scanner.Scan()
	line := scanner.Text()
	ds := make(Directions, 0)
	for _, d := range line {
		ds = append(ds, d)
	}

	// Skip empty line
	scanner.Scan()

	re := regexp.MustCompile(`[A-Z|0-9]{3}`)

	// Read map
	nodes := make(map[string]*Node)
	initialNodes := make([]*Node, 0)
	for scanner.Scan() {
		line := scanner.Text()

		matches := re.FindAllString(line, -1)
		if matches == nil {
			log.Fatal("did not match nodes")
		}

		label := matches[0]
		left := matches[1]
		right := matches[2]

		node, ok := nodes[label]
		if !ok {
			node = &Node{label, nil, nil}
			nodes[label] = node
		}

		leftNode, ok := nodes[left]
		if !ok {
			leftNode = &Node{left, nil, nil}
			nodes[left] = leftNode
		}

		rightNode, ok := nodes[right]
		if !ok {
			rightNode = &Node{right, nil, nil}
			nodes[right] = rightNode
		}

		node.left = leftNode
		node.right = rightNode

		if strings.HasSuffix(node.label, "A") {
			initialNodes = append(initialNodes, node)
		}
	}

	return ds, initialNodes
}

func main() {
	ds, nodes := readMap("input.txt")
	fmt.Println(ds)
	for _, n := range nodes {
		fmt.Println(n.label)
	}

	i := 0
	n := len(ds)
	for {
		switch ds[i%n] {
		case 'L':
			for j, node := range nodes {
				nodes[j] = node.left
			}
		case 'R':
			for j, node := range nodes {
				nodes[j] = node.right
			}
		}

		numNodesEndWithZ := 0
		for _, node := range nodes {
			if strings.HasSuffix(node.label, "Z") {
				numNodesEndWithZ++
			}
		}

		i++

		if numNodesEndWithZ == len(nodes) {
			break
		}
	}
	fmt.Println(i)
}
