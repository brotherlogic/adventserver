package main

import "testing"

func Test2016_14_1(t *testing.T) {
	indexes := buildKeys("abc")

	if indexes[64] != 22728 {
		t.Errorf("Bad Key: %v", indexes[64])
	}
}
