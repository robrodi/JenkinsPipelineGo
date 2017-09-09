package main

import "testing"

func TestAdd(t *testing.T) {
	expected := 5
  actual := add(3, 2)
  if (expected != actual) {
    t.Error("Nope")
  }
}
