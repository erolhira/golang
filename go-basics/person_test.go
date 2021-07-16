package main

import "testing"

/**
Test functions starts with uppercase and with Test
 */
func TestNewPeople(t *testing.T) {

	p := createPeople()
	if len(p) != 8 {
		t.Errorf("Expected length is 8 but got %d", len(p))
	}
}