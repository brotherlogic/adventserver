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

	result := runProgram(context.Background(), program, tlog, 0)
	if result.a != 2 {
		t.Errorf("Bad result %v (2)", result.a)
	}
}

func Test2015Day23Part1Sup(t *testing.T) {
	program := `jio a, +18
	inc a
	tpl a
	inc a
	tpl a
	tpl a
	tpl a
	inc a
	tpl a
	inc a
	tpl a
	inc a
	inc a
	tpl a
	tpl a
	tpl a
	inc a
	jmp +22
	tpl a
	inc a
	tpl a
	inc a
	inc a
	tpl a
	inc a
	tpl a
	inc a
	inc a
	tpl a
	tpl a
	inc a
	inc a
	tpl a
	inc a
	inc a
	tpl a
	inc a
	inc a
	tpl a
	jio a, +8
	inc b
	jie a, +4
	tpl a
	inc a
	jmp +2
	hlf a
	jmp -7`

	result := runProgram(context.Background(), program, tlog, 0)
	if result.a != 1 {
		t.Errorf("Bad result %v (2)", result)
	}
}
