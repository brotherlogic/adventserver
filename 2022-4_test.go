package main

import "testing"

func Test2022_4_1(t *testing.T) {
	data := `2-4,6-8
	2-3,4-5
	5-7,7-9
	2-8,3-7
	6-6,4-6
	2-6,4-8`

	count := countFullyOverlapping(data)
	if count != 2 {
		t.Errorf("Not the right overlap: %v (2)", count)
	}
}

func Test2022_4_2(t *testing.T) {
	data := `2-4,6-8
	2-3,4-5
	5-7,7-9
	2-8,3-7
	6-6,4-6
	2-6,4-8`

	count := countOverlapping(data)
	if count != 4 {
		t.Errorf("Not the right overlap: %v (2)", count)
	}
}

func Test2022_4_1_Overlap(t *testing.T) {
	tests := []struct {
		in      string
		overlap bool
	}{
		{"2-8,3-7", true},
		{"6-6,4-6", true},
		{"2-6,4-8", false},
		{"2-4,6-8", false},
		{"", false},
	}

	for _, test := range tests {
		overlap := doesOverlap(test.in)
		if overlap != test.overlap {
			t.Errorf("Bad overlap determineation: %v -> %v / %v", test.in, overlap, test.overlap)
		}
	}
}

func Test2022_4_2_Overlap(t *testing.T) {
	tests := []struct {
		in      string
		overlap bool
	}{
		{"2-8,3-7", true},
		{"6-6,4-6", true},
		{"2-6,4-8", true},
		{"2-4,6-8", false},
		{"", false},
	}

	for _, test := range tests {
		overlap := doesOverlapEven(test.in)
		if overlap != test.overlap {
			t.Errorf("Bad overlap determineation: %v -> %v / %v", test.in, overlap, test.overlap)
		}
	}
}
