package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	input := CleanInput("Hello World")
	expeted := []string{"hello", "world"}
	if len(input) != len(expeted) {
		t.Errorf("Lengths not matching")
	} else {
		for i, _ := range input {
			if input[i] != expeted[i] {
				t.Errorf("Inputs not matching")
			}
		}
	}

}
