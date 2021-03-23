package model

import (
	"fmt"
	"strings"
)

/*
 * Nodes are by default mutable and should be passed as pointers.
 */
type Node struct {
	neighbors []*Node
	ID        string
	Color     int
}

func (n *Node) Neighbors() []*Node {
	return n.neighbors
}

func (n *Node) AddNeighbor(neigh *Node) {
	n.neighbors = append(n.neighbors, neigh)
}

func NewNode(ID string) *Node {
	n := Node{}
	n.ID = ID
	n.Color = -1
	return &n
}

func (n *Node) String() string {
	toStringArray := func(nodes []*Node) []string {
		strarr := []string{}
		for _, n := range nodes {
			strarr = append(strarr, n.ID)
		}

		return strarr
	}
	return fmt.Sprintf("{ID: %s, Color: %d, neighbors: [%s]}", n.ID, n.Color, strings.Join(toStringArray(n.neighbors), ","))
}
