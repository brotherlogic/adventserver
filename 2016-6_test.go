package main

import "testing"

func Test2016Day6(t *testing.T) {
	data := `eedadn
drvtee
eandsr
raavrd
atevrs
tsrnev
sdttsa
rasrtv
nssdts
ntnada
svetve
tesnvt
vntsnd
vrdear
dvrsen
enarar`

	common := getCommon(data)
	if common != "easter" {
		t.Errorf("Bad trans %v vs easter", common)
	}

	leastCommon := getLeastCommon(data)
	if leastCommon != "advent" {
		t.Errorf("Bad trans %v vs advert", leastCommon)
	}
}
