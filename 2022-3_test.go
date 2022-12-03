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

func Test2022_3_2(t *testing.T) {
	data := `vJrwpWtwJgWrhcsFMMfFFhFp
	jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
	PmmdzqPrVvPwwTWBwg
	wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
	ttgJtRGJQctTZtZT
	CrZsJsPPZsGzwwsLwLmpwMDw`

	s := sumOfCommons(data)

	if s != 70 {
		t.Errorf("Bad sum of commons: %v (70)", s)
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

func Test2022_3_2_Pick(t *testing.T) {
	tests := []struct {
		in1, in2, in3 string
		out           string
		value         int
	}{
		{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg", "r", 18},
		{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw", "Z", 52},
	}

	for _, test := range tests {
		let := getFCommon(test.in1, test.in2, test.in3)
		if let != test.out {
			t.Errorf("Bad get common: %v (%v) -> %v", let, test.out, test.in1)
		} else {
			val := getPriority(let)
			if val != test.value {
				t.Errorf("Bad get priority: %v => %v (%v) -> %v", let, val, test.value, test.in1)
			}
		}
	}
}
