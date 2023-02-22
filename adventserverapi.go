package main

import (
	"fmt"
	"reflect"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"golang.org/x/net/context"

	pbaoc "github.com/brotherlogic/adventofcode/proto"
	pb "github.com/brotherlogic/adventserver/proto"
)

var (
	runtime = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "adventserver_runtime",
		Help: "Runtime",
	}, []string{"year", "day"})
)

func (s *Server) Solve(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {

	if req.GetYear() == 2017 && req.GetDay() == 13 && req.GetPart() == 2 {
		conn, err := s.FDialServer(ctx, "adventofcode")
		if err != nil {
			return nil, err
		}
		client := pbaoc.NewAdventServerServiceClient(conn)
		resp, err := client.Solve(ctx, &pbaoc.SolveRequest{Year: req.GetYear(), Day: req.GetDay(), Part: req.GetPart()})
		if err != nil {
			return nil, err
		}
		return &pb.SolveResponse{Answer: resp.GetAnswer()}, nil
	}

	if !reflect.ValueOf(s).MethodByName(fmt.Sprintf("Solve%vday%vpart%v", req.GetYear(), req.GetDay(), req.GetPart())).IsValid() {
		return nil, fmt.Errorf("cannot find %v", reflect.ValueOf(s).MethodByName(fmt.Sprintf("Solve%vday%vpart%v", req.GetYear(), req.GetDay(), req.GetPart())))

	}
	val := reflect.ValueOf(s).MethodByName(fmt.Sprintf("Solve%vday%vpart%v", req.GetYear(), req.GetDay(), req.GetPart())).Call([]reflect.Value{reflect.ValueOf(ctx)})
	if val[1].Interface() == nil {
		return val[0].Interface().(*pb.SolveResponse), nil
	}
	return val[0].Interface().(*pb.SolveResponse), val[1].Interface().(error)
}
