package main

import (
	"context"
	"log"
	"testing"
)

func tlog(ctx context.Context, lstr string) {
	log.Printf(lstr)
}

func Test2015Day23Part1(t *testing.T) {
	program := `inc a
	jio a, +2
	tpl a
	inc a`

	result := runProgram(context.Background(), program, tlog)
	if result.a != 2 {
		t.Errorf("Bad result %v (2)", result.a)
	}
}
