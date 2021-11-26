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
	return nil, fmt.Errorf("Not implemented yet")
}
