package main

import "testing"

func Test2022_22_1_Main(t *testing.T) {
	data = `        ...#
        .#..
        #...
        ....
...#.......#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.

10R5L5R10L4R5L5`

	res := runFunnyMaze(data)
	if res != 6032 {
		t.Errorf("Bad result: %v (6032)", res)
	}
}
