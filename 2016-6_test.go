package main

import "testing"

func Test2016Day6Part1(t *testing.T) {
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
}
