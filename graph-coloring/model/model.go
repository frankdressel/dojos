package model

import "strings"

type Node struct {
	neighbors []Node
}

func (n *Node) Neighbors() []Node {
	return n.neighbors
}

func (n *Node) Equal(other Node) bool {
	if len(n.neighbors) != len(other.neighbors) {
		return false
	}

	for i, _ := range n.neighbors {
		if !n.neighbors[i].Equal(other.neighbors[i]) {
			return false
		}
	}

	return true
}

func (n *Node) AddNeighbor(neigh Node) {
	n.neighbors = append(n.neighbors, neigh)
}
