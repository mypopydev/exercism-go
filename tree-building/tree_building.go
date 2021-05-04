package tree

import (
	"sort"
	"errors"
)

// Record is a struct containing int fields ID and Parent
type Record struct {
	ID      int
	Parent  int
}

// Node is a struct containing int field ID and []*Node field Children.
type Node struct {
	ID       int
	Children []*Node
}

// Build will be to implement the tree building logic for these records.
func Build(records []Record) (*Node, error) {
	// check empty record
	if records == nil {
		return nil, nil
	}

	// sort records
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	recordMap := make(map[int]*Node)

	for i, r := range records {
		// check for id continuity
		if r.ID != i {
			return nil, errors.New("Non continous record id")
		}

		// check for id duplication
		if _, ok := recordMap[r.ID]; ok {
			return nil, errors.New("Duplications are not allowed")
		}

		// current ID < Paraent ID
		if r.ID != 0 && r.ID <= r.Parent {
			return nil, errors.New("Invalid ID number")
		}

		// Root id have a paraent
		if r.ID == 0 && r.Parent != 0 {
			return nil, errors.New("Root node cannot have a parent")
		}

		n := &Node{ID: r.ID}
		recordMap[r.ID] = n

		if _, ok := recordMap[r.Parent]; ok && r.ID != 0 {
			recordMap[r.Parent].Children = append(recordMap[r.Parent].Children, n)
		}
	}

	return recordMap[0], nil
}
