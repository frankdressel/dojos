package utils

import (
	"moduliertersingvogel.de/graph-coloring/model"
	"strings"
)

func Parse(tuples string) []*model.Node {
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

	nodelist := make([]*model.Node, 0)
	for _, n := range ids2node {
		nodelist = append(nodelist, n)
	}

	return nodelist
}
