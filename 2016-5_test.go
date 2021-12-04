package main

import "testing"

func Test2016Day5Part1(t *testing.T) {
	password := computeHashPass("abc")
	if password != "18f47a30" {
		t.Errorf("Bad password %v vs %v", password, "18f47a30")
	}
}
