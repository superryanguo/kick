package main

import "testing"

func TestParseFile(t *testing.T) {
	order, err := parseFile(file)

	if err != nil {
		t.Errorf("Could not parse file: %v", err)
	}

}
