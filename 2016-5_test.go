package main

import "testing"

func Test2016Day5Part1(t *testing.T) {
	password := computeHashPass("abc")
	if password != "18f47a30" {
		t.Errorf("Bad password %v vs %v", password, "18f47a30")
	}
}
func Test2016Day5Part2(t *testing.T) {
	password := computeHashPass2("abc")
	if password != "05ace8e3" {
		t.Errorf("Bad password %v vs %v", password, "05ace8e3")
	}
}
