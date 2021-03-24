package main

import (
	"fmt"
	"moduliertersingvogel.de/graph-coloring/model"
	"moduliertersingvogel.de/graph-coloring/utils"
	"os"
	"strings"
)

func contains(nodes map[*model.Node]bool, node *model.Node) bool {
	for n := range nodes {
		if n == node {
			return true
		}
	}
	return false
}

func colorAllowed(node *model.Node) bool {
	for _, n := range node.Neighbors() {
		if node.Color == n.Color {
			return false
		}
	}
	return true
}

func countColorsMap(nodelist map[*model.Node]bool) map[int]int {
	colorNum := make(map[int]int)
	for n := range nodelist {
		if num, ok := colorNum[n.Color]; ok {
			colorNum[n.Color] = num + 1
		} else {
			colorNum[n.Color] = 1
		}
	}

	return colorNum
}

func countColorsList(nodelist []*model.Node) map[int]int {
	colorNum := make(map[int]int)
	for _, n := range nodelist {
		if num, ok := colorNum[n.Color]; ok {
			colorNum[n.Color] = num + 1
		} else {
			colorNum[n.Color] = 1
		}
	}

	return colorNum
}

func greedy_search(node *model.Node, visited map[*model.Node]bool, colors []int) {
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
			greedy_search(n, visited, colors)
		}
	}
}

func full_search(remaining []*model.Node, visited []*model.Node, max_color int) {
	if len(remaining) == 0 {
		colorNum := countColorsList(visited)
		fmt.Printf("%d colors used\n", len(colorNum))
		return
	}
	last := remaining[len(remaining)-1]
	visited_new := append(visited, last)
	remaining_new := remaining[0 : len(remaining)-1]
	for _, r := range remaining {
		r.Color = -1
	}
	for c := last.Color + 1; c < max_color; c++ {
		last.Color = c
		if colorAllowed(last) {
			full_search(remaining_new, visited_new, max_color)
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
	for i := range colors {
		colors[i] = i
	}

	subgraphs := utils.Subgraphs(nodelist)
	for _, s := range subgraphs {
		for n := range s {
			fmt.Println(n)
			visited := make(map[*model.Node]bool)
			greedy_search(n, visited, colors)

			break
		}
	}

	colorNum := countColorsMap(nodelist)
	fmt.Printf("%d colors used\n", len(colorNum))

	for n := range nodelist {
		fmt.Printf("Node %s: Color: %d\n", n.ID, n.Color)
	}

	for _, s := range subgraphs {
		visited := make([]*model.Node, 0)
		remaining := make([]*model.Node, 0)
		for n := range s {
			remaining = append(remaining, n)
		}
		full_search(remaining, visited, len(s))
	}
}
