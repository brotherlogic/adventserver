package main

import (
	"log"
	"math"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func getSnaf(in int64) string {
	switch in {
	case 0:
		return "0"
	case 1:
		return "1"
	case 2:
		return "2"
	case -1:
		return "-"
	case -2:
		return "="
	}

	return "WHATWHATWHAT"
}

func getVal(in rune) int64 {
	switch in {
	case '2':
		return 2
	case '1':
		return 1
	case '0':
		return 0
	case '-':
		return -1
	case '=':
		return -2
	}

	log.Fatalf("BAD RUNE: %v", in)
	return -100
}

func snafu(in string) int64 {
	val := int64(0)
	for i := 0; i < len(in); i++ {
		mult := int64(math.Pow(5, float64(i)))
		if mult == 0 {
			mult = 1
		}
		vval := getVal(rune(in[len(in)-1-i]))
		val += mult * vval
	}
	return val
}

func pow64(a, b int64) int64 {
	return int64(math.Pow(float64(a), float64(b)))
}

func rep(str string, c int) string {
	ret := ""

	for i := 0; i < c; i++ {
		ret += str
	}

	return ret
}

func rsnafu(in int64) string {
	if in == 0 {
		return "0"
	}
	if in == 1 {
		return "1"
	}
	if in == 2 {
		return "2"
	}
	if in == 3 {
		return "1="
	}

	rbound := int64(5 - 2)
	gbound := int64(0)
	count := (0)

	for in > rbound {
		count++
		rbound = pow64(int64(5), int64(count))
		gbound = 2 * pow64(int64(5), int64(count-1))
		for i := count - 1; i >= 0; i-- {
			rbound -= 2 * pow64(int64(5), int64(i))
			if i > 0 {
				gbound -= 2 * pow64(int64(5), int64((i-1)))
			}
		}
	}

	// Value is now 1 or 2
	if in < gbound {
		lower := pow64(5, int64(count-1))
		if in > lower {
			rval := rsnafu(in - lower)
			return "1" + rep("0", count-len(rval)-1) + rval
		} else {
			rval := reverseSnaf(rsnafu(lower - in))
			return "1" + rep("0", count-len(rval)-1) + rval
		}
	}
	lower := 2 * pow64(5, int64(count-1))

	if in > lower {
		rval := rsnafu(in - lower)
		return "2" + rep("0", count-len(rval)-1) + rval
	}
	rval := reverseSnaf(rsnafu(lower - in))
	return "2" + rep("0", count-len(rval)-1) + rval
}

func reverseSnaf(in string) string {
	out := ""
	for _, char := range in {
		switch char {
		case '2':
			out += "="
		case '1':
			out += "-"
		case '0':
			out += "0"
		case '-':
			out += "1"
		case '=':
			out += "2"
		}
	}

	return out
}

func computeSnafuSum(data string) string {
	sumv := int64(0)
	for _, line := range strings.Split(strings.TrimSpace(data), "\n") {
		sumv += snafu(strings.TrimSpace(line))
	}

	return rsnafu(sumv)
}

func (s *Server) Solve2022day25part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-25.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{StringAnswer: computeSnafuSum(data)}, nil
}
