package main

func TestFullP1(t *testing.T)){
	data :=`Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
	Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.`

	winner := getDistance(data, 1000)
	if winner != 1120 {
		t.Errorf("Bad Distance: %v should have been 1120", distance)
	}
}