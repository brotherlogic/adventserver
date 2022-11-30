package main

import "testing"

func Test2016_12_1(t *testing.T) {
	data := `cpy 41 a
	inc a
	inc a
	dec a
	jnz a 2
	dec a`

	state := runMonorailProgram(data, false)

	if state.a != 42 {
		t.Errorf("Bad program state: %+v", state)
	}
}

func Test2016_12_1_jump(t *testing.T) {
	data := `cpy 41 a
	inc a
	inc a
	dec a
	jnz d 2
	dec a`

	state := runMonorailProgram(data, false)

	if state.a != 41 {
		t.Errorf("Bad program state: %+v", state)
	}
}

func Test2016_12_1_blank(t *testing.T) {
	data := `cpy 41 a
	inc a
	inc a
	dec a
	jnz d 2
	dec a
	`

	state := runMonorailProgram(data, false)

	if state.a != 41 {
		t.Errorf("Bad program state: %+v", state)
	}
}
