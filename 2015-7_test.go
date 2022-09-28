package main

import "testing"

func TestDay7P1(t *testing.T) {

	cases := []struct {
		in   string
		want uint
	}{
		{"d", 72},
		{"e", 507},
		{"f", 492},
		{"g", 114},
		{"h", 65412},
		{"i", 65079},
		{"j", 123},
		{"lj", 123},
		{"x", 123},
		{"y", 456},
	}

	var rules map[string]string
	rules = make(map[string]string)

	rules["x"] = "123"
	rules["y"] = "456"
	rules["d"] = "x AND y"
	rules["e"] = "x OR y"
	rules["f"] = "x LSHIFT 2"
	rules["g"] = "y RSHIFT 2"
	rules["h"] = "NOT x"
	rules["i"] = "NOT y"
	rules["j"] = "x"
	rules["lj"] = "ky"
	rules["ky"] = "x"

	for _, c := range cases {
		answer := WorkRules(rules, c.in)

		if answer != c.want {
			t.Errorf("Spec(%q) == %d, want %d", c.in, answer, c.want)
		}
	}
}
