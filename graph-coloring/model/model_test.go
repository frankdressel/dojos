package model

import "testing"

func TestEmptyNode(t *testing.T) {
	n := Node{}
	neigh := n.Neighbors()

	if len(neigh) != 0 {
		t.Errorf("Empty node should have no neighbors but have: %q", len(neigh))
	}
}

func TestEqual(t *testing.T) {
	n1 := Node{}
	n2 := Node{}

	if !n1.Equal(n2) {
		t.Errorf("Equal nodes are not equal")
	}
	if !n2.Equal(n1) {
		t.Errorf("Equal nodes are not equal")
	}

	n3 := Node{}
	n1 = Node{[]Node{n3}}
	n2 = Node{[]Node{n3}}

	if !n1.Equal(n2) {
		t.Errorf("Equal nodes are not equal")
	}
	if !n2.Equal(n1) {
		t.Errorf("Equal nodes are not equal")
	}
}

func TestNeighborAddinf(t *testing.T) {
	n1 := Node{}
	n2 := Node{}

	if len(n1.Neighbors()) != 0 {
		t.Errorf("Neighbor list for clean node not empty")
	}
	n1.AddNeighbor(n2)
	if len(n1.Neighbors()) != 1 {
		t.Errorf("Added neighbor not in neighbors list")
	}
	if len(n2.Neighbors()) != 0 {
		t.Errorf("Added neighbor not added directional")
	}
}
