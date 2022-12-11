package main

import "testing"

func Test2022_11_1_Build(t *testing.T) {
	data := `Monkey 0:
	Starting items: 79, 98
	Operation: new = old * 19
	Test: divisible by 23
	  If true: throw to monkey 2
	  If false: throw to monkey 3`

	monkeys := buildMonkeys(data)

	if len(monkeys) != 1 {
		t.Errorf("Too many monkeys: %v", len(monkeys))
	}

	if monkeys[0].number != 0 {
		t.Errorf("Wrong number: %v", monkeys[0].number)
	}

	if len(monkeys[0].items) != 2 || monkeys[0].items[0] != 79 {
		t.Errorf("Not read items: %v", monkeys[0].items)
	}
	if monkeys[0].adjustment != 19 ||
		monkeys[0].falseMonkey != 3 ||
		monkeys[0].number != 0 ||
		monkeys[0].operation != "*" ||
		monkeys[0].test != 23 ||
		monkeys[0].trueMonkey != 2 {
		t.Errorf("Bad monkey: %+v", monkeys[0])
	}
}

func Test2022_11_1(t *testing.T) {
	data := `Monkey 0:
	Starting items: 79, 98
	Operation: new = old * 19
	Test: divisible by 23
	  If true: throw to monkey 2
	  If false: throw to monkey 3
  
  Monkey 1:
	Starting items: 54, 65, 75, 74
	Operation: new = old + 6
	Test: divisible by 19
	  If true: throw to monkey 2
	  If false: throw to monkey 0
  
  Monkey 2:
	Starting items: 79, 60, 97
	Operation: new = old * old
	Test: divisible by 13
	  If true: throw to monkey 1
	  If false: throw to monkey 3
  
  Monkey 3:
	Starting items: 74
	Operation: new = old + 3
	Test: divisible by 17
	  If true: throw to monkey 0
	  If false: throw to monkey 1`

	monkeys := buildMonkeys(data)

	for i := 0; i < 20; i++ {
		runMonkeys(monkeys)
	}

	values := getMonkeyTimes(monkeys)
	if values[0]*values[1] != 10605 {
		t.Errorf("Bad monkey run: %v", values[0]*values[1])
	}
}