package main

import "testing"

func Test2017Day4Part2(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"aaaaa-bbb-z-y-x-123[abxyz]", 123},
		{"a-b-c-d-e-f-g-h-987[abcde]", 987},
		{"not-a-real-room-404[oarel]", 404},
		{"totally-real-room-200[decoy]", 0},
	}

	for _, c := range cases {
		got := isRealRoom(c.in)
		if got != c.want {
			t.Errorf("Spec(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}
