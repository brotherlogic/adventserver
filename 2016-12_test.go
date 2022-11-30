package main

import "testing"

func Test2016_answer(t *testing.T) {
	data := `cpy 1 a
	cpy 1 b
	cpy 26 d
	jnz c 2
	jnz 1 5
	cpy 7 c
	inc d
	dec c
	jnz c -2
	cpy a c
	inc a
	dec b
	jnz b -2
	cpy c b
	dec d
	jnz d -6
	cpy 19 c
	cpy 11 d
	inc a
	dec d
	jnz d -2
	dec c
	jnz c -5`

	state := runMonorailProgram(data)

	if state.a != 42 {
		t.Errorf("Bad program state: %+v", state)
	}

}

func Test2016_12_1(t *testing.T) {
	data := `cpy 41 a
	inc a
	inc a
	dec a
	jnz a 2
	dec a`

	state := runMonorailProgram(data)

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

	state := runMonorailProgram(data)

	if state.a != 41 {
		t.Errorf("Bad program state: %+v", state)
	}
}
