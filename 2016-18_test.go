package main

import "testing"

func Test2016_18_1_Transform(t *testing.T) {
	line := "..^^."

	trans := nextLine(line)

	if trans != ".^^^^" {
		t.Errorf("Bad next line: %v (%v)", trans, ".^^^^.")
	}
}

func Test2016_18_1(t *testing.T) {
	line := ".^^.^.^^^^"
	count := runTiles(line, 10)

	if count != 38 {
		t.Errorf("Bad count: %v (38)", count)
	}

}
