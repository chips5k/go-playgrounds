package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func buildAndActivateGrid(file string) (*[][]byte, int, error) {
	f, err := os.Open(file)

	if err != nil {
		return nil, 0, fmt.Errorf("Failed to open file with error: %s", err)
	}

	scanner := bufio.NewScanner(f)

	grid := make([][]byte, 0, 200)

	beams := make([]int, 0, 200)

	splitterActivations := 0

	row := 0
	for scanner.Scan() {
		l := scanner.Bytes()

		bb := make([]byte, len(l))
		copy(bb, l)

		// render beams for current row
		if len(beams) > 0 {
			for i, s := range beams {
				if s != -1 {
					if bb[s] != '^' && bb[s] != 'S' && bb[s] != '|' {
						bb[s] = '|'
					} else {
						beams[i] = -1
					}
				}
			}
		}

		// calculate beams for next row
		for i, v := range bb {
			if v == 'S' {
				beams = append(beams, i)
			}

			if v == '^' && grid[row-1][i] == '|' {
				splitterActivations++
				if i-1 >= 0 && grid[row-1][i-1] != '|' {
					bb[i-1] = '|'
					beams = append(beams, i-1)
				}
				if i+1 < len(bb) && grid[row-1][i+1] != '|' {
					bb[i+1] = '|'
					beams = append(beams, i+1)
				}
			}
		}

		grid = append(grid, bb)

		row++
	}

	return &grid, splitterActivations, nil
}
func main_part_1() {

	grid, splitterActivations, err := buildAndActivateGrid("input")

	if err != nil {
		log.Fatal(err)
	}

	for _, r := range *grid {
		for _, c := range r {
			fmt.Printf("%c", c)
		}
		fmt.Println()
	}

	fmt.Printf("Splitter Activations: %d\n", splitterActivations)
}

type NodeValue struct {
	x int
	y int
	v byte
}
type Node struct {
	Left  *Node
	Value NodeValue
	Right *Node
}

func main_part_2() {
	grid, _, err := buildAndActivateGrid("input")
	if err != nil {
		log.Fatal(err)
	}

	root := &Node{nil, NodeValue{}, nil}
	buildNode(root, grid, -1, -1)

	//fmt.Printf("%s\n\n", *grid)

	//PrintTree(root)

	fmt.Printf("Total Paths: %d\n", countPaths(root, 0))
}

var countPathsCache map[string]int64 = make(map[string]int64)

func countPaths(node *Node, total int64) int64 {

	if node != nil {
		cacheKey := fmt.Sprintf("%d-%d-%c", node.Value.x, node.Value.y, node.Value.v)
		cacheValue, cacheHit := countPathsCache[cacheKey]
		if node.Left == nil && node.Right == nil {
			return 1
		}

		if cacheHit {
			return cacheValue
		}

		amt := total + countPaths(node.Left, total) + countPaths(node.Right, total)

		countPathsCache[cacheKey] = amt
		return amt

	}
	return 0
}

var buildNodeCache map[string]*Node = make(map[string]*Node)

func buildNode(parent *Node, grid *[][]byte, startX int, startY int) {

	g := *grid

	// find entry point
	for y, row := range g {

		for x, cell := range row {

			if y > startY && !(x > startX && y == startY) {
				if cell == 'S' {
					parent.Value = NodeValue{
						x: x,
						y: y,
						v: cell,
					}
					buildNode(parent, grid, 0, y)
					return
				}

				if cell == '|' {

					// if aligns with parent, its a regular beam continuation
					if x == parent.Value.x && y-1 == parent.Value.y {
						n := &Node{
							nil,
							NodeValue{
								x: x,
								y: y,
								v: cell,
							},
							nil,
						}

						parent.Left = n
						buildNode(parent.Left, grid, x, y-1)
						return
					}
				}

				if cell == '^' {
					if x == parent.Value.x && y-1 == parent.Value.y && parent.Value.v == '|' {

						cacheKey := fmt.Sprintf("%d-%d-%c-%d-%d-%c", parent.Value.x, parent.Value.y, parent.Value.v, startX, startY, cell)
						cacheValue, cacheHit := buildNodeCache[cacheKey]

						if cacheHit {
							parent.Left = cacheValue.Left
							parent.Right = cacheValue.Right
							return
						}

						parent.Left = &Node{
							nil,
							NodeValue{
								x: x - 1,
								y: y,
								v: '|',
							},
							nil,
						}
						buildNode(parent.Left, grid, x, y)

						parent.Right = &Node{
							nil,
							NodeValue{
								x: x + 1,
								y: y,
								v: '|',
							},
							nil,
						}
						buildNode(parent.Right, grid, x, y)

						buildNodeCache[cacheKey] = parent

						return
					}
				}
			}
		}
		//fmt.Println()

	}

}

// PrintTree initiates the recursive printing
func PrintTree(root *Node) {
	printRecursive(root, "", true, "Root")
}

func printRecursive(n *Node, indent string, isLast bool, label string) {
	if n == nil {
		return
	}

	// Print the current node with a branch prefix
	marker := "├── "
	if isLast {
		marker = "└── "
	}

	fmt.Printf("%s%s[%s] \"%c\" (%d, %d)\n", indent, marker, label, n.Value.v, n.Value.y+1, n.Value.x+1)

	// Update indent for children
	newIndent := indent
	if isLast {
		newIndent += "    "
	} else {
		newIndent += "│   "
	}

	// Recurse: Right child first if you want it on top,
	// or Left first for standard order.
	if n.Left != nil || n.Right != nil {
		printRecursive(n.Left, newIndent, n.Right == nil, "L")
		printRecursive(n.Right, newIndent, true, "R")
	}
}

func main() {
	//main_part_1()
	main_part_2()

}
