package main

import "testing"

func Test2022_21_1_Main(t *testing.T) {
	data := `root: pppw + sjmn
dbpl: 5
cczh: sllz + lgvd
zczc: 2
ptdq: humn - dvpt
dvpt: 3
lfqf: 4
humn: 5
ljgn: 2
sjmn: drzm * dbpl
sllz: 4
pppw: cczh / lfqf
lgvd: ljgn * ptdq
drzm: hmdt - zczc
hmdt: 32`

	program := buildProgram(data)
	res, _ := evalProg(program, "root", "")
	if res != 152 {
		t.Errorf("Bad prog run: %v (152)", res)
	}
}

func Test2022_21_2_Main(t *testing.T) {
	data := `root: pppw + sjmn
dbpl: 5
cczh: sllz + lgvd
zczc: 2
ptdq: humn - dvpt
dvpt: 3
lfqf: 4
humn: 5
ljgn: 2
sjmn: drzm * dbpl
sllz: 4
pppw: cczh / lfqf
lgvd: ljgn * ptdq
drzm: hmdt - zczc
hmdt: 32`

	program := buildProgram(data)
	program.progs["humn"].unknown = true

	val, _ := evalProg(program, program.progs["root"].right, "")

	result := findUnknown(program, program.progs["root"].left, val)

	if result != 301 {
		t.Errorf("Bad prog run: %v (301)", result)
	}
}
