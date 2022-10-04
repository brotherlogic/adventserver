package main

import "testing"

func TestFullP1(t *testing.T) {
	data := `Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
	Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.`

	winner := getDistance(data, 1000)
	if winner != 1120 {
		t.Errorf("Bad Distance: %v should have been 1120", winner)
	}
}

func TestPartialDay14(t *testing.T) {
	cases := []struct {
		spd      int
		blast    int
		wait     int
		time     int
		distance int
	}{
		{14, 10, 127, 1, 14},
		{16, 11, 162, 1, 16},
		{14, 10, 127, 11, 140},
		{16, 11, 162, 11, 176},
	}

	for _, c := range cases {
		distance := computeDistance(&reindeer{speed: c.spd, blastTime: c.blast, waitTime: c.wait}, c.time)
		if distance != c.distance {
			t.Errorf("%+v @ %v -> %v (should have been %v", &reindeer{speed: c.spd, blastTime: c.blast, waitTime: c.wait}, c.time, distance, c.distance)
		}
	}
}
