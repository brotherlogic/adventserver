package main

import "testing"

func Test2021Day10Overall(t *testing.T) {
	data := `[({(<(())[]>[[{[]{<()<>>
		[(()[<>])]({[<{<<[]>>(
		{([(<{}[<>[]}>{[]{[(<()>
		(((({<>}<{<{<>}{[]{[]{}
		[[<[([]))<([[{}[[()]]]
		[{[{({}]{}}([{[{{{}}([]
		{<[[]]>}<{[{[{[]{()[[[]
		[<(<(<(<{}))><([]([]()
		<{([([[(<>()){}]>(<<{{
		<{([{{}}[<[[[<>{}]]]>[]]`

	sum := getSum(data)
	if sum != 26397 {
		t.Errorf("Bad Sum %v vs 26397", sum)
	}

	sum2 := getSum2(data)
	if sum2 != 288957 {
		t.Errorf("Bad Sum %v vs 288957", sum2)
	}
}

func Test2021Day10(t *testing.T) {
	cases := []struct {
		in    string
		want  string
		want2 string
	}{
		{"[({(<(())[]>[[{[]{<()<>>", "", "}}]])})]"},
		{"[(()[<>])]({[<{<<[]>>(", "", ")}>]})"},
		{"{([(<{}[<>[]}>{[]{[(<()>", "}", ""},
		{"(((({<>}<{<{<>}{[]{[]{}", "", "}}>}>))))"},
		{"[[<[([]))<([[{}[[()]]]", ")", ""},
		{"[{[{({}]{}}([{[{{{}}([]", "]", ""},
		{"{<[[]]>}<{[{[{[]{()[[[]", "", "]]}}]}]}>"},
		{"[<(<(<(<{}))><([]([]()", ")", ""},
		{"<{([([[(<>()){}]>(<<{{", ">", ""},
		{"<{([{{}}[<[[[<>{}]]]>[]]", "", "])}>"},
	}

	for _, c := range cases {
		got, got2 := getFirstInvalid(c.in)
		if got != c.want || got2 != c.want2 {
			t.Errorf("Spec(%v) == %v, want %v but %v want %v", c.in, got, c.want, got2, c.want2)
		}
	}
}

func Test2021Day10Focus(t *testing.T) {
	data := "[({(<(())[]>[[{[]{<()<>>"
	_, res := getFirstInvalid(data)
	if res != "}}]])})]" {
		t.Errorf("Bad Focus: %v", res)
	}
}
