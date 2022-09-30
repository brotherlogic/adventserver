package main

import (
	"context"
	"testing"
)

func Test2015Day9Part1(t *testing.T) {
	details := `London to Dublin = 464
	London to Belfast = 518
	Dublin to Belfast = 141`

	server := Init()
	result := server.computeBestDistance(context.Background(), details)
	if result != 605 {
		t.Errorf("Wrong result %v vs 605", result)
	}
}
