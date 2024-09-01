package main

import "testing"

func TestAdd(t *testing.T) {
	total := add(2, 2)
	if total != 4 {
		t.Errorf("Add(2, 2) = %d; want 4", total)
	}
}
