package main

import "testing"

func Test2016_19_1(t *testing.T) {
	p := runPresents(5)
	if p != 3 {
		t.Errorf("Bad present run: %v (3)", p)
	}
}

func Test2016_19_2(t *testing.T) {
	p := runCircularPresents(5)
	if p != 2 {
		t.Errorf("Bad present run: %v (3)", p)
	}
}
