package model

import "testing"

func TestEmptyNode(t *testing.T) {
	n := NewNode("1")
	neigh := n.Neighbors()

	if len(neigh) != 0 {
		t.Errorf("Empty node should have no neighbors but have: %q", len(neigh))
	}
}

func TestNeighborAddinf(t *testing.T) {
	n1 := NewNode("1")
	n2 := NewNode("2")
	n3 := NewNode("3")

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
	n1.AddNeighbor(n3)
	if len(n1.Neighbors()) != 2 {
		t.Errorf("Added neighbor not in neighbors list")
	}
}

func TestIDs(t *testing.T) {
	n1 := NewNode("1")
	n2 := NewNode("2")

	if n1.ID == n2.ID {
		t.Errorf("IDs of two distinct nodes are IDentical: %s vs. %s", n1.ID, n2.ID)
	}
}
