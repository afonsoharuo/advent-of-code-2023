package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
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

func calcLCM(a, b *big.Int) *big.Int {
	gcd := new(big.Int)
	div := new(big.Int)
	prod := new(big.Int)

	gcd.GCD(nil, nil, a, b)
	div.Div(b, gcd)
	prod.Mul(a, div)

	return prod
}

func main() {
	ds, nodes := readMap("input.txt")
	for _, n := range nodes {
		fmt.Println(n.label)
	}

	n := len(ds)
	indexAtZ := make([]*big.Int, len(nodes))
	for k, node := range nodes {
		i := 0
		for {
			switch ds[i%n] {
			case 'L':
				node = node.left
			case 'R':
				node = node.right
			}

			i++

			if strings.HasSuffix(node.label, "Z") {
				break
			}
		}
		indexAtZ[k] = big.NewInt(int64(i))
	}

	lcm := indexAtZ[0]
	for _, value := range indexAtZ[1:] {
		lcm = calcLCM(lcm, value)
	}

	fmt.Println(indexAtZ)
	fmt.Println(lcm)
}
