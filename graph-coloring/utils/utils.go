package utils

import (
	"moduliertersingvogel.de/graph-coloring/model"
	"strings"
)

func Parse(tuples string) map[*model.Node]bool {
	ids2node := make(map[string]*model.Node)
	for _, a := range strings.Split(tuples, " ") {
		split := strings.Split(a, "-")
		if len(split) > 2 {
			panic("Invalid argument")
		}

		name1 := split[0]
		n1, ok1 := ids2node[name1]
		if !ok1 {
			n1 = model.NewNode(name1)
			ids2node[name1] = n1
		}
		if len(split) == 2 {
			name2 := split[1]
			n2, ok2 := ids2node[name2]
			if !ok2 {
				n2 = model.NewNode(name2)
				ids2node[name2] = n2
			}
			n1.AddNeighbor(n2)
			n2.AddNeighbor(n1)
		}
	}

	nodelist := make(map[*model.Node]bool)
	for _, n := range ids2node {
		nodelist[n] = true
	}

	return nodelist
}

func findSubgraphs(startNode *model.Node, visited map[*model.Node]bool) {
	visited[startNode] = true
	for _, n := range startNode.Neighbors() {
		if _, ok := visited[n]; !ok {
			findSubgraphs(n, visited)
		}
	}
}

func Subgraphs(nodelist map[*model.Node]bool) []map[*model.Node]bool {
	remaining := make(map[*model.Node]bool)
	for k, v := range nodelist {
		remaining[k] = v
	}
	var result []map[*model.Node]bool
	for len(remaining) > 0 {
		for n := range remaining {
			subgraph := make(map[*model.Node]bool)
			findSubgraphs(n, subgraph)
			result = append(result, subgraph)
			for s := range subgraph {
				delete(remaining, s)
			}
		}
	}
	return result
}
