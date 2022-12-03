package main

import (
	"testing"
)

func TestExpansion(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"1", "100"},
		{"0", "001"},
		{"11111", "11111000000"},
		{"111100001010", "1111000010100101011110000"},
	}

	for _, test := range tests {
		expansion := dragonExpand(test.in)
		if expansion != test.out {
			t.Errorf("Bad expansion: %v -> %v (%v)", test.in, expansion, test.out)
		}
	}
}

func TestChecksum(t *testing.T) {
	cs := dragonChecksum(dragonChecksum("110010110100"))
	if cs != "100" {
		t.Errorf("bad checksum: %v", cs)
	}
}

func Test2016_16_1(t *testing.T) {
	result := dragonRun("10000", 20)
	if result != "01100" {
		t.Errorf("Bad dragon run %v (%v)", result, "01100")
	}
}
