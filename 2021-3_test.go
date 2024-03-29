package main

import "testing"

func TestDay3Part1(t *testing.T) {
	data := "00100\n11110\n10110\n10111\n10101\n01111\n00111\n11100\n10000\n11001\n00010\n01010"

	val := computePower(data)
	if val != 198 {
		t.Errorf("Bad power %v -> %v", val, 198)
	}
}

func TestDay3Part2(t *testing.T) {
	data := "00100\n11110\n10110\n10111\n10101\n01111\n00111\n11100\n10000\n11001\n00010\n01010"

	val := computeLSR(data) * computeOGR(data)
	if val != 230 {
		t.Errorf("Bad power %v -> %v", val, 230)
	}
}
