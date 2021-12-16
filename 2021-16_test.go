package main

import (
	"log"
	"testing"
)

func Test2021Day16ConvertHex(t *testing.T) {
	str := "D2FE28"
	res := convertHex(str)
	if res != "110100101111111000101000" {
		t.Errorf("Bad conversion  %v", res)
		t.Errorf("Bad conversion 110100101111111000101000")
	}

	elem, _ := parseCode(res, -1)
	if len(elem) == 0 {
		t.Fatalf("No code")
	}
	if elem[0].pid != 4 || elem[0].version != 6 || elem[0].value != 2021 {
		t.Errorf("Bad code: %+v", elem[0])
	}

	elem, _ = parseCode("110100010100101001000100100", -1)
	if elem[0].pid != 4 {
		t.Errorf("Bad elem: %+v", elem[0])
	}
}

func Test2021Day16ConvertHarder(t *testing.T) {
	str := "38006F45291200"
	res, _ := parseCode(convertHex(str), -1)

	if len(res) != 1 {
		t.Errorf("bad res : %v", res)
	}

	if len(res[0].subcodes) != 2 {
		t.Errorf("Bad subcodes: %+v", res[0])
	}

	res, _ = parseCode(convertHex("EE00D40C823060"), -1)
}

func Test2021Day16(t *testing.T) {

	cases := []struct {
		in   string
		want int
	}{
		{"8A004A801A8002F478", 16},
		{"620080001611562C8802118E34", 12},
		{"C0015000016115A2E0802F182340", 23},
		{"A0016C880162017C3686B18A3D4780", 31},
	}

	for _, scase := range cases {
		log.Printf("Running  %v", scase.in)
		pc, _ := parseCode(convertHex(scase.in), -1)
		count := sumVersion(pc)
		if count != scase.want {
			t.Errorf("Bad comp: %v vs %v (%v)", count, scase.want, scase.in)
		}
	}
}
