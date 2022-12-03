package main

import "testing"

func Test2022_3_1(t *testing.T) {
	data := `vJrwpWtwJgWrhcsFMMfFFhFp
	jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
	PmmdzqPrVvPwwTWBwg
	wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
	ttgJtRGJQctTZtZT
	CrZsJsPPZsGzwwsLwLmpwMDw`

	s := sumOfPriorities(data)

	if s != 157 {
		t.Errorf("Bad sum of priorities: %v (157)", s)
	}
}
