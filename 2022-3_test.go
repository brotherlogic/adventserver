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

func Test2022_3_1_Pick(t *testing.T) {
	tests := []struct {
		in    string
		out   string
		value int
	}{
		{"vJrwpWtwJgWrhcsFMMfFFhFp", "p", 16},
		{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "L", 38},
		{"PmmdzqPrVvPwwTWBwg", "P", 42},
		{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "v", 22},
		{"ttgJtRGJQctTZtZT", "t", 20},
		{"CrZsJsPPZsGzwwsLwLmpwMDw", "s", 19},
	}

	for _, test := range tests {
		let := getPCommon(test.in)
		if let != test.out {
			t.Errorf("Bad get common: %v (%v) -> %v", let, test.out, test.in)
		} else {
			val := getPriority(let)
			if val != test.value {
				t.Errorf("Bad get priority: %v => %v (%v) -> %v", let, val, test.value, test.in)
			}
		}
	}
}
