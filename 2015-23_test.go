package main

import "testing"

func Test2015Day23Part1(t *testing.T) {
	program := `inc a
	jio a, +2
	tpl a
	inc a`

	result := runProgram(program)
	if result.a != 2 {
		t.Errorf("Bad result %v (2)", result.a)
	}
}
