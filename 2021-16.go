package main

import (
	"fmt"
	"log"
	"math/big"
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
	value    *big.Int
	subcodes []packet
}

func convert(bin string) *big.Int {
	val := new(big.Int)
	val.SetString(bin, 2)
	return val
}

func sumVersion(code []packet) int {
	sv := 0
	for _, c := range code {
		sv += int(c.version) + sumVersion(c.subcodes)
	}
	return sv
}

func computeCode(code packet) *big.Int {
	log.Printf("RUNNING %v", code.pid)
	switch code.pid {
	case 0:
		sumv := big.NewInt(0)
		for _, sc := range code.subcodes {
			sumv.Add(sumv, computeCode(sc))
		}
		return sumv
	case 1:
		log.Printf("PROD: %v", computeCode(code.subcodes[0]))
		prodv := big.NewInt(1)
		for _, sc := range code.subcodes {
			log.Printf("PROD %v", computeCode(sc))
			prodv.Mul(prodv, computeCode(sc))
		}
		return prodv
	case 2:
		maxv := computeCode(code.subcodes[0])
		for _, sc := range code.subcodes {
			val := computeCode(sc)
			if maxv.Cmp(val) > 0 {
				maxv = val
			}
		}
		return maxv
	case 3:
		maxv := big.NewInt(0)
		for _, sc := range code.subcodes {
			val := computeCode(sc)
			log.Printf("MAX: %v", val)
			if maxv.Cmp(val) < 0 {
				maxv = val
			}
		}
		log.Printf("RET %v", maxv)
		return maxv
	case 4:
		return code.value
	case 5:
		if computeCode(code.subcodes[0]).Cmp(computeCode(code.subcodes[1])) > 0 {
			return big.NewInt(1)
		}
		return big.NewInt(0)
	case 6:
		if computeCode(code.subcodes[0]).Cmp(computeCode(code.subcodes[1])) < 0 {
			return big.NewInt(1)
		}
		return big.NewInt(0)
	case 7:
		if computeCode(code.subcodes[0]).Cmp(computeCode(code.subcodes[1])) == 0 {
			return big.NewInt(1)
		}
		return big.NewInt(0)
	default:
		log.Fatalf("Cannot process: %+v", code.pid)
	}

	return big.NewInt(0)
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
		code.version = int(convert(bin[pointer : pointer+3]).Int64())
		code.pid = int(convert(bin[pointer+3 : pointer+6]).Int64())
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
			code.value = new(big.Int)
			code.value.SetString(sstr, 2)
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

func (s *Server) print(code packet) {
	for _, sc := range code.subcodes {
		s.print(sc)
	}
}

func (s *Server) Solve2021day16part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2021-16.txt")
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(data)

	pc, _ := parseCode(convertHex(trimmed), -1)

	s.print(pc[0])
	count := computeCode(pc[0])
	return &pb.SolveResponse{StringAnswer: count.String()}, nil
}
