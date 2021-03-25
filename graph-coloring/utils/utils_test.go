package utils

import (
	"moduliertersingvogel.de/graph-coloring/model"
	"sort"
	"testing"
)

type byNumberOfArguments []map[*model.Node]bool

func (a byNumberOfArguments) Len() int           { return len(a) }
func (a byNumberOfArguments) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byNumberOfArguments) Less(i, j int) bool { return len(a[i]) < len(a[j]) }

func TestParse(t *testing.T) {
	test := "1-2 1-3 2-3 4"
	parsed := Parse(test)

	if len(parsed) != 4 {
		t.Errorf("Wrong number of nodes parsed. Expected 4, got %d", len(parsed))
	}

	zero_counter := 0
	two_counter := 0
	for n := range parsed {
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

func TestSubgraphs(t *testing.T) {
	parsed := Parse("1-2 1-3 2-3 4")

	listOfSubs := Subgraphs(parsed)

	if len(listOfSubs) != 2 {
		t.Errorf("Wrong number of subgraphs found: Expected 2 but got %d", len(listOfSubs))
	}

	sort.Sort(byNumberOfArguments(listOfSubs))
	if len(listOfSubs[0]) != 1 {
		t.Errorf("Expecting 1 element in smallest cluster but was %d", len(listOfSubs[0]))
	}
	for n := range listOfSubs[0] {
		if n.ID != "4" {
			t.Errorf("Expecting node 4 but was %s", n.ID)
		}
	}
	if len(listOfSubs[1]) != 3 {
		t.Errorf("Expecting 3 element in smallest cluster but was %d", len(listOfSubs[1]))
	}
}

func TestInfiniteIntStreamNext(t *testing.T) {
	i := NewInfiniteIntStream(0)

	next := i.GetAndIncrement()
	if next != 0 {
		t.Errorf("Expected 0 as next but got %d", next)
	}

	next = i.GetAndIncrement()
	if next != 1 {
		t.Errorf("Expected 1 as next but got %d", next)
	}

	i = NewInfiniteIntStream(6)

	next = i.GetAndIncrement()
	if next != 6 {
		t.Errorf("Expected 6 as next but got %d", next)
	}
}
