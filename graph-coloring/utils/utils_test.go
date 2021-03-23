package utils

import (
	"testing"
)

func TestParse(t *testing.T) {
	test := "1-2 1-3 2-3 4"
	parsed := Parse(test)

	if len(parsed) != 4 {
		t.Errorf("Wrong number of nodes parsed. Expected 4, got %d", len(parsed))
	}

	zero_counter := 0
	two_counter := 0
	for _, n := range parsed {
		if len(n.Neighbors()) == 0 {
			zero_counter = zero_counter + 1
		}
		if len(n.Neighbors()) == 2 {
			two_counter = two_counter + 1
		}
	}
	if zero_counter != 1 {
		t.Errorf("Expecting one node without neighbors but didn't found one")
	}
	if two_counter != 3 {
		t.Errorf("Expecting three nodes with 2 neighbors but found %d", two_counter)
	}
}
