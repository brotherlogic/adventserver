package main

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

func solveHash(str string, len int) int {
	value := 0
	found := false
	match := ""
	for i := 0; i < len; i++ {
		match += "0"
	}

	for !found {
		value++
		data := []byte(str + strconv.Itoa(value))
		hash := md5.New()
		hash.Write(data)
		md5v := hex.EncodeToString(hash.Sum(nil))
		if strings.HasPrefix(md5v, match) {
			found = true
		}
	}
	return value
}

func (s *Server) Solve2015day4part1(ctx context.Context) (*pb.SolveResponse, error) {
	return &pb.SolveResponse{Answer: int32(solveHash("ckczppom", 5))}, nil
}

func (s *Server) Solve2015day4part2(ctx context.Context) (*pb.SolveResponse, error) {
	return &pb.SolveResponse{Answer: int32(solveHash("ckczppom", 6))}, nil
}
