package main

import (
	"fmt"
	"testing"
)

func TestParseFile(t *testing.T) {
	order, err := parseFile(defaultFilename)

	if err != nil {
		t.Errorf("Could not parse file: %v\n", err)
	} else {
		fmt.Printf("Parse Json Ok->order=%v\n", order)
	}

}
