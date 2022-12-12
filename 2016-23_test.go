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
