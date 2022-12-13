package main

import "testing"

func Test2016_23_1_Main(t *testing.T) {
	data := `cpy 2 a
	tgl a
	tgl a
	tgl a
	cpy 1 a
	dec a
	dec a`

	res := runToggleProgram(data, 7)
	if res.a != 3 {
		t.Errorf("Bad register: %+v", res)
	}
}

func Test2016_23_2_Full(t *testing.T) {
	data := `cpy a b
	dec b
	cpy a d
	cpy 0 a
	cpy b c
	inc a
	dec c
	jnz c -2
	dec d
	jnz d -5
	dec b
	cpy b c
	cpy c d
	dec d
	inc c
	jnz d -2
	tgl c
	cpy -16 c
	jnz 1 c
	cpy 96 c
	jnz 95 d
	inc a
	inc d
	jnz d -2
	inc c
	jnz c -5`

	res := runToggleProgram(data, 12)
	if res.a != 3 {
		t.Errorf("Bad register: %+v", res)
	}
}
