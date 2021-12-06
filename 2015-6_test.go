package main

import "testing"

func TestDay6P2(t *testing.T) {

	cases := []struct {
		in   string
		want int
	}{
		{"turn on 0,0 through 0,0", 1},
		{"toggle 0,0 through 999,999", 2000000},
	}

	for _, c := range cases {
		lights := MakeBrightBoard(1000)
		lights = ToggleBrightness(c.in, lights)

		count := 0
		size := 0
		for i := 0; i < len(lights); i++ {
			for j := 0; j < len(lights[i]); j++ {
				size++
				count += lights[i][j]
			}
		}

		if count != c.want {
			t.Errorf("Spec(%q) == %d, want %d", c.in, count, c.want)
		}

		if size != 1000*1000 {
			t.Errorf("Size is %d, should be %d", size, 1000*1000)
		}
	}
}

func TestDay6P1(t *testing.T) {

	cases := []struct {
		in   string
		want int
	}{
		{"turn on 0,0 through 999,999", 1000 * 1000},
		{"turn on 0,0 through 999,0", 1000},
		{"turn off 499,499 through 500,500", 0},
		{"toggle 499,499 through 500,500", 4},
	}

	for _, c := range cases {
		lights := MakeBoard(1000)
		lights = ToggleLights(c.in, lights)

		count := 0
		size := 0
		for i := 0; i < len(lights); i++ {
			for j := 0; j < len(lights[i]); j++ {
				size++
				if lights[i][j] {
					count++
				}
			}
		}

		if count != c.want {
			t.Errorf("Spec(%q) == %d, want %d", c.in, count, c.want)
		}

		if size != 1000*1000 {
			t.Errorf("Size is %d, should be %d", size, 1000*1000)
		}
	}
}
