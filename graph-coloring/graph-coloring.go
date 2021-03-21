package main

import (
	"fmt"
	"moduliertersingvogel.de/graph-coloring/model"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: graph-coloring <Node1>-<Node2> ...")
	}
	ids2node := make(map[string]model.Node)
	for _, a := range args {
		split := strings.Split(a, "-")
		if len(split) > 2 {
			panic("Invalid argument")
		}

		n1, ok1 := ids2node[split[0]]
		n2, ok2 := ids2node[split[1]]
		if !ok1 {
			n1 = model.Node{}
			ids2node[split[0]] = n1
		}
		if !ok2 {
			n2 = model.Node{}
			ids2node[split[1]] = n2
		}
		n1.AddNeighbor(n2)
		fmt.Println(n1)
		n2.AddNeighbor(n1)
	}

	fmt.Println(ids2node)
}
