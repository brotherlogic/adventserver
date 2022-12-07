package main

import "testing"

func Test2022_7_1_Main(t *testing.T) {
	data := `$ cd /
	$ ls
	dir a
	14848514 b.txt
	8504156 c.dat
	dir d
	$ cd a
	$ ls
	dir e
	29116 f
	2557 g
	62596 h.lst
	$ cd e
	$ ls
	584 i
	$ cd ..
	$ cd ..
	$ cd d
	$ ls
	4060174 j
	8033020 d.log
	5626152 d.ext
	7214296 k`

	list := buildDirs(data)

	sumv := list.dirSum(100000)

	if sumv != 95437 {
		t.Errorf("Bad sum of dirs: %v (95437)", sumv)
	}
}

func Test2022_7_2_Main(t *testing.T) {
	data := `$ cd /
	$ ls
	dir a
	14848514 b.txt
	8504156 c.dat
	dir d
	$ cd a
	$ ls
	dir e
	29116 f
	2557 g
	62596 h.lst
	$ cd e
	$ ls
	584 i
	$ cd ..
	$ cd ..
	$ cd d
	$ ls
	4060174 j
	8033020 d.log
	5626152 d.ext
	7214296 k`

	list := buildDirs(data)

	sumv := list.remove(70000000, 30000000)

	if sumv != 24933642 {
		t.Errorf("Bad sum of dirs: %v (24933642)", sumv)
	}
}
