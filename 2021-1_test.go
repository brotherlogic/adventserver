package main

import "testing"

func Test2021Day1Part1(t *testing.T) {
	input := "199\n200\n208\n210\n200\n207\n240\n269\n260\n263"
	count := countInc(input)

	if count != int32(7) {
		t.Errorf("Bad input: %v vs %v", count, 7)
	}
}
