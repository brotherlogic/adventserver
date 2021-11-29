package main

import "testing"

func Test2016Day3P1(t *testing.T) {
	if isTriangle(5, 10, 25) {
		t.Errorf("Bad triangle")
	}

	if validTriangles("5 10 25\n3 3 3") != 1 {
		t.Errorf("Bad all triangles")
	}
}
