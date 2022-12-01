package main

import (
	"log"
	"strings"
	"testing"
)

func Test2016_14_1(t *testing.T) {
	indexes := buildKeys("abc")

	if indexes[64] != 22728 {
		t.Errorf("Bad Key: %v", indexes[64])
	}
}

func Test2016Sup_14_1(t *testing.T) {
	key := makeKey("abc", 18)
	if !strings.Contains(key, "cc38887a5") {
		t.Errorf("Bad key: %v", key)
	}

	threes := countThrees(key)
	found := false
	if threes == "8" {
		found = true
	}

	if !found {
		t.Errorf("No three: %v", threes)
	}
}

func Test2016SupFive_14_1(t *testing.T) {
	key := makeKey("abc", 39)

	threes := countThrees(key)
	found := false
	if threes == "e" {
		found = true
	}

	if !found {
		t.Errorf("No e: %v", threes)
	}

	key2 := makeKey("abc", 816)
	fives := countFives(key2)
	found = false
	for _, val := range fives {
		if val == "e" {
			found = true
		}
	}

	if !found {
		t.Errorf("No e: %v", key2)
	}
}

func Test2016SupKey_14_1(t *testing.T) {
	fives := make(map[string][]int)

	threes, _ := buildKey("abc", 39, fives)

	found := false
	if threes == "e" {
		found = true
	}

	if !found {
		t.Errorf("e not found: %v", threes)
	}

	found = false
	for k, v := range fives {
		if k == "e" {
			for _, val := range v {
				if val == 816 {
					found = true
				}
			}
		}
	}

	if !found {
		t.Errorf("Build key failed: %v and %v", threes, fives)
	}
}

func Test2016SupDiff(t *testing.T) {
	f := findFives("abc")
	if f != 22728 {
		t.Errorf("Bad: %v", f)
	}
}

func Test2016SupKeyLast_14_1(t *testing.T) {
	fives := make(map[string][]int)

	key := makeKey("abc", 22728)
	threes, f := buildKey("abc", 22728, fives)
	log.Printf("F %v", f)
	fives = f

	found := false
	if threes == "c" {
		found = true
	}

	if !found {
		t.Errorf("e not found: %v -> %v", threes, key)
	}

	found = false
	for k, v := range fives {
		if k == "c" {
			for _, val := range v {
				if val == 22804 {
					found = true
				}
			}
		}
	}

	if !found {
		t.Errorf("Build key failed: %v and %v", threes, fives)
	}
}
