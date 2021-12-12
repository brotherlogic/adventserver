package main

import (
	"strings"
	"testing"
)

func TestDay12_2021(t *testing.T) {
	data := `start-A
	start-b
	A-c
	A-b
	b-d
	A-end
	b-end`

	paths := countPaths(buildGraph(strings.TrimSpace(data)))

	if paths != 10 {
		t.Errorf("Bad path count: %v", paths)
	}
}
func TestDay12_2021_e2(t *testing.T) {
	data := `dc-end
	HN-start
	start-kj
	dc-start
	dc-HN
	LN-dc
	HN-end
	kj-sa
	kj-HN
	kj-dc`

	paths := countPaths(buildGraph(strings.TrimSpace(data)))

	if paths != 19 {
		t.Errorf("Bad path count: %v", paths)
	}
}

func TestDay12_2021_e3(t *testing.T) {
	data := `fs-end
	he-DX
	fs-he
	start-DX
	pj-DX
	end-zg
	zg-sl
	zg-pj
	pj-he
	RW-he
	fs-DX
	pj-RW
	zg-RW
	start-pj
	he-WI
	zg-he
	pj-fs
	start-RW`

	paths := countPaths(buildGraph(strings.TrimSpace(data)))

	if paths != 226 {
		t.Errorf("Bad path count: %v", paths)
	}
}
