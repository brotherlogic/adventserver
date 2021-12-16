package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func appendZeros(str string) string {
	switch len(str) {
	case 1:
		return "000" + str
	case 2:
		return "00" + str
	case 3:
		return "0" + str
	}
	return str
}

func convertHex(hex string) string {
	bstr := ""
	for _, c := range hex {
		val, _ := strconv.ParseInt(string(c), 16, 32)
		bstr += appendZeros(fmt.Sprintf("%b", val))
	}
	return bstr
}

type packet struct {
	version  int
	pid      int
	value    int64
	subcodes []packet
}

func convert(bin string) int64 {
	val, _ := strconv.ParseInt(bin, 2, 32)
	return val
}

func sumVersion(code []packet) int {
	sv := 0
	for _, c := range code {
		sv += int(c.version) + sumVersion(c.subcodes)
	}
	return sv
}

func parseCode(bin string, maxlen int) ([]packet, int) {
	log.Printf("PARSE %v", bin)
	var codes []packet

	pointer := 0
	for pointer < len(bin) {
		log.Printf("WORK  %v", bin[pointer:])
		check := false
		for _, c := range bin[pointer:] {
			if string(c) == "1" {
				check = true
			}
		}

		if !check {
			return codes, 0
		}

		code := packet{}
		code.version = int(convert(bin[pointer : pointer+3]))
		code.pid = int(convert(bin[pointer+3 : pointer+6]))
		log.Printf("READ: %+v", code)

		switch code.pid {
		case 4:
			count := 0
			fnum := ""
			sstr := ""
			for !strings.HasPrefix(fnum, "0") {
				fnum = bin[pointer+6+count*5 : pointer+6+count*5+5]

				sstr += fnum[1:]
				count++
			}
			conv, _ := strconv.ParseInt(sstr, 2, 64)
			code.value = conv
			pointer = pointer + 6 + count*5
		default:
			typeid := string(bin[pointer+6])
			log.Printf("READ TYPE: %v", typeid)
			switch typeid {
			case "0":
				plen, _ := (strconv.ParseInt(bin[pointer+6+1:pointer+6+16], 2, 32))
				log.Printf("PLEN: %v -> %v", plen, bin[pointer+6+1:pointer+6+16])
				code.subcodes, _ = parseCode(bin[pointer+6+16:pointer+6+16+int(plen)], -1)
				pointer += 6 + 16 + int(plen)
			case "1":
				maxnum, _ := strconv.ParseInt(bin[pointer+6+1:pointer+6+12], 2, 32)
				log.Printf("SUBS: %v", maxnum)
				sc, len := parseCode(bin[pointer+6+12:], int(maxnum))
				code.subcodes = sc
				pointer += 6 + 12 + len
			default:
				log.Fatalf("BAD Type %v", typeid)
			}
		}

		codes = append(codes, code)
		if maxlen > 0 && len(codes) >= maxlen {
			return codes, pointer
		}
	}

	return codes, 0
}

func (s *Server) Solve2021day16part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2021-16.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	pc, _ := parseCode(convertHex(trimmed), -1)
	count := sumVersion(pc)
	return &pb.SolveResponse{Answer: int32(count)}, nil
}
