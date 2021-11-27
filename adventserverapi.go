package main

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"golang.org/x/net/context"

	pb "github.com/brotherlogic/adventserver/proto"
)

var (
	runtime = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "adventserver_runtime",
		Help: "Runtime",
	}, []string{"year", "day"})
)

func (s *Server) Solve(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	if req.GetYear() == 2015 && req.GetDay() == 1 && req.GetPart() == 1 {
		ans, err := s.solve2015day1part1(ctx)
		if err != nil {
			return nil, err
		}
		return &pb.SolveResponse{Answer: ans}, nil
	} else if req.GetYear() == 2015 && req.GetDay() == 1 && req.GetPart() == 2 {
		ans, err := s.solve2015day1part2(ctx)
		if err != nil {
			return nil, err
		}
		return &pb.SolveResponse{Answer: ans}, nil
	} else if req.GetYear() == 2016 && req.GetDay() == 1 && req.GetPart() == 1 {
		ans, err := s.solve2016day1part1(ctx)
		if err != nil {
			return nil, err
		}
		return &pb.SolveResponse{Answer: ans}, nil
	} else if req.GetYear() == 2016 && req.GetDay() == 1 && req.GetPart() == 2 {
		ans, err := s.solve2016day1part2(ctx)
		if err != nil {
			return nil, err
		}
		return &pb.SolveResponse{Answer: ans}, nil
	}
	return nil, fmt.Errorf("Not implemented yet")
}
