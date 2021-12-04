package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func getHash(key string) string {
	hash := md5.New()
	hash.Write([]byte(key))
	return hex.EncodeToString(hash.Sum(nil))
}

func computeHashPass(key string) string {
	password := ""
	for i := 0; i < math.MaxInt32; i++ {
		if strings.HasPrefix(getHash(fmt.Sprintf("%v%v", key, i)), "00000") {
			password += string(getHash(fmt.Sprintf("%v%v", key, i))[5])
		}

		if len(password) == 8 {
			return password
		}
	}
	return ""
}

func computeHashPass2(key string) string {
	password := "00000000"
	done := make(map[int]bool)
	for i := 0; i < math.MaxInt32; i++ {
		if strings.HasPrefix(getHash(fmt.Sprintf("%v%v", key, i)), "00000") {
			val := string(getHash(fmt.Sprintf("%v%v", key, i)))
			index, err := strconv.Atoi(string(val[5]))
			if err == nil && index < 8 && !done[index] {
				password = password[:index] + string(val[6]) + password[index+1:]
				done[index] = true
			}
		}

		if len(done) == 8 {
			return password
		}
	}
	return ""
}

func (s *Server) Solve2016day5part1(ctx context.Context) (*pb.SolveResponse, error) {
	return &pb.SolveResponse{StringAnswer: computeHashPass("ojvtpuvg")}, nil
}

func (s *Server) Solve2016day5part2(ctx context.Context) (*pb.SolveResponse, error) {
	return &pb.SolveResponse{StringAnswer: computeHashPass2("ojvtpuvg")}, nil
}
