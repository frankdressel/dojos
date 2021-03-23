package main

import (
	"fmt"
	"moduliertersingvogel.de/graph-coloring/model"
	"moduliertersingvogel.de/graph-coloring/utils"
	"os"
	"strings"
)

func contains(nodes map[*model.Node]bool, node *model.Node) bool {
	for n, _ := range nodes {
		if n == node {
			return true
		}
	}
	return false
}

func search(node *model.Node, visited map[*model.Node]bool, colors []int) {
	if !contains(visited, node) {
		used_for_neighbors := make(map[int]bool, len(node.Neighbors()))
		for _, neigh := range node.Neighbors() {
			used_for_neighbors[neigh.Color] = true
		}
		for _, c := range colors {
			if _, ok := used_for_neighbors[c]; !ok {
				node.Color = c
				break
			}
		}
		visited[node] = true
		for _, n := range node.Neighbors() {
			search(n, visited, colors)
		}
	}
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: graph-coloring <Node1>-<Node2> ...")
	}
	nodelist := utils.Parse(strings.Join(args, " "))
	colors := make([]int, len(nodelist))
	for i, _ := range colors {
		colors[i] = i
	}

	visited := make(map[*model.Node]bool)
	for len(visited) < len(nodelist) {
		for _, n := range nodelist {
			_, ok := visited[n]
			if !ok {
				search(n, visited, colors)
				break
			}
		}
	}
	colorNum := make(map[int]int)
	for _, n := range nodelist {
		if num, ok := colorNum[n.Color]; ok {
			colorNum[n.Color] = num + 1
		} else {
			colorNum[n.Color] = 1
		}
	}
	fmt.Printf("%d colors used\n", len(colorNum))

	for _, n := range nodelist {
		fmt.Printf("Node %s: Color: %d\n", n.ID, n.Color)
	}
}
