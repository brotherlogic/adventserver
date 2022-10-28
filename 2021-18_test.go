package main

import (
	"testing"
)

func Test2021Day18Parse(t *testing.T) {
	cases := []struct {
		num string
	}{
		{"[1,2]"},
		{"[[1,2],3]"},
		{"[9,[8,7]]"},
		{"[[1,9],[8,5]]"},
		{"[[[[1,2],[3,4]],[[5,6],[7,8]]],9]"},
		{"[[[9,[3,8]],[[0,9],6]],[[[3,7],[4,9]],3]]"},
		{"[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]"},
	}

	for _, tc := range cases {
		if tc.num != printNum(parseNum(tc.num)) {
			t.Fatalf("Cannot process: %v -> %v", tc.num, printNum(parseNum(tc.num)))
		}
	}
}

func Test2021Day18Focus(t *testing.T) {
	parseNum("[[1,9],[8,5]]")
}

func Test2021Day18Split(t *testing.T) {
	cases := []struct {
		start int
		end   string
	}{
		{10, "[5,5]"},
		{11, "[5,6]"},
		{12, "[6,6]"},
		{9, "9"},
	}
	for _, tc := range cases {
		val := printNum(split(&snnum{val: tc.start}))
		if tc.end != val {
			t.Fatalf("Cannot process: %v -> %v vs %v", tc.start, val, tc.end)
		}
	}
}

func Test2021Day18Explode(t *testing.T) {
	cases := []struct {
		start string
		end   string
	}{
		{"[[[[[9,8],1],2],3],4]", "[[[[0,9],2],3],4]"},
		{"[7,[6,[5,[4,[3,2]]]]]", "[7,[6,[5,[7,0]]]]"},
		{"[[6,[5,[4,[3,2]]]],1]", "[[6,[5,[7,0]]],3]"},
		{"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"},
		{"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[7,0]]]]"},
		{"[[[[0,9],2],3],4]", "[[[[0,9],2],3],4]"},
	}

	for _, tc := range cases {
		val := printNum(explode(parseNum(tc.start)))
		if tc.end != val {
			t.Fatalf("Cannot process: %v -> %v vs %v", tc.start, val, tc.end)
		}
	}
}

func Test2021Day18Part1(t *testing.T) {
	data := `[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
	[[[5,[2,8]],4],[5,[[9,9],0]]]
	[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
	[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
	[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
	[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
	[[[[5,4],[7,7]],8],[[8,3],8]]
	[[9,3],[[9,9],[6,[4,9]]]]
	[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
	[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]`

	magnitude := magnitude(runSum(data))
	if magnitude != 4140 {
		t.Errorf("Bad mag: %v vs 4140", magnitude)
	}
}

func Test2021Day18Part2(t *testing.T) {
	data := `[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
	[[[5,[2,8]],4],[5,[[9,9],0]]]
	[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
	[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
	[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
	[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
	[[[[5,4],[7,7]],8],[[8,3],8]]
	[[9,3],[[9,9],[6,[4,9]]]]
	[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
	[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]`

	magnitude := bestSum(data)
	if magnitude != 3993 {
		t.Errorf("Bad mag: %v vs 3993", magnitude)
	}
}

func Test2021Day18Magnitude(t *testing.T) {
	cases := []struct {
		start string
		end   int
	}{
		{"[[1,2],[[3,4],5]]", 143},
		{"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", 1384},
		{"[[[[1,1],[2,2]],[3,3]],[4,4]]", 445},
		{"[[[[3,0],[5,3]],[4,4]],[5,5]]", 791},
		{"[[[[5,0],[7,4]],[5,5]],[6,6]]", 1137},
		{"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", 3488},
	}

	for _, tc := range cases {
		val := magnitude(parseNum(tc.start))
		if tc.end != val {
			t.Fatalf("Cannot process: %v -> %v vs %v", tc.start, val, tc.end)
		}
	}
}

func Test2021Day18Reduce(t *testing.T) {
	nnum := reduceNum(parseNum("[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]"))
	if printNum(nnum) != "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]" {
		t.Errorf("Unable to reduce -> %v", printNum(nnum))
	}
}

func Test2021Day18Add(t *testing.T) {
	nnum := add(parseNum("[[[[4,3],4],4],[7,[[8,4],9]]]"), parseNum("[1,1]"))
	if printNum(nnum) != "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]" {
		t.Errorf("Unable to add -> %v", printNum(nnum))
	}
}

func Test2021Day18RunAddition(t *testing.T) {
	cases := []struct {
		start string
		end   string
	}{
		{`[1,1]
		[2,2]
		[3,3]
		[4,4]`, "[[[[1,1],[2,2]],[3,3]],[4,4]]"},
		{`[1,1]
		[2,2]
		[3,3]
		[4,4]
		[5,5]`, "[[[[3,0],[5,3]],[4,4]],[5,5]]"},
		{`[1,1]
		[2,2]
		[3,3]
		[4,4]
		[5,5]
		[6,6]`, "[[[[5,0],[7,4]],[5,5]],[6,6]]"},
		{`[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]
		[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]
		[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]
		[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]
		[7,[5,[[3,8],[1,4]]]]
		[[2,[2,2]],[8,[8,1]]]
		[2,9]
		[1,[[[9,3],9],[[9,0],[0,7]]]]
		[[[5,[7,4]],7],1]
		[[[[4,2],2],6],[8,7]]`, "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]"},
	}

	for _, tc := range cases {
		val := runSum(tc.start)
		if tc.end != printNum(val) {
			t.Fatalf("Cannot process: %v -> %v vs %v", tc.start, printNum(val), tc.end)
		}
	}
}

func Test2021Day18SingleExplode(t *testing.T) {
	num := parseNum("[[[[[1,1],[2,2]],[3,3]],[4,4]],[5,5]]")
	res := runExplode(num, nil, 0)
	if !res {
		t.Errorf("Bad explode: %v", printNum(num))
	}

	runExplode(num, nil, 0)

	runExplode(num, nil, 0)
}
