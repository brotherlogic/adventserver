package main

import "testing"

func Test2022_9_1_Main(t *testing.T) {
	data := `R 4
	U 4
	L 3
	D 1
	R 4
	D 1
	L 5
	R 2`

	result := runRopeBridge(data, 2)

	if result != 13 {
		t.Errorf("Bad rope bridge run: %v", result)
	}
}

func Test2022_9_2_Main(t *testing.T) {
	data := `R 4
	U 4
	L 3
	D 1
	R 4
	D 1
	L 5
	R 2`

	result := runRopeBridge(data, 9)

	if result != 1 {
		t.Errorf("Bad rope bridge run part 2: %v", result)
	}
}

func Test2022_9_2_Extra(t *testing.T) {
	data := `R 5
	U 8
	L 8
	D 3
	R 17
	D 10
	L 25
	U 20`

	result := runRopeBridge(data, 10)

	if result != 36 {
		t.Errorf("Bad rope bridge run part 2 extra: %v", result)
	}
}

func Test2022_9_2_Sup(t *testing.T) {
	m1, m2 := ropeMove(2, -2, 0, 0)
	if m1 != 1 && m2 != -1 {
		t.Errorf("Bad move, %v %v", m1, m2)
	}
}
