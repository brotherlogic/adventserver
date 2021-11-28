package main

import "testing"

func Test2016Day1Part1(t *testing.T) {
	code := "ULL\nRRDDD\nLURDL\nUUUUD"
	result := procCode(code, 1, 1)
	if result != 1985 {
		t.Errorf("Bad code: %v should have been %v", result, 1985)
	}
}
