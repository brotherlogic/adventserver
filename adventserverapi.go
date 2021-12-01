package main

import (
	"fmt"
	"reflect"

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
	if reflect.ValueOf(s).MethodByName(fmt.Sprintf("Solve%vday%vpart%v", req.GetYear(), req.GetDay(), req.GetPart())).IsZero() {
		return nil, fmt.Errorf("cannot find %v", reflect.ValueOf(s).MethodByName(fmt.Sprintf("Solve%vday%vpart%v", req.GetYear(), req.GetDay(), req.GetPart())))

	}
	val := reflect.ValueOf(s).MethodByName(fmt.Sprintf("Solve%vday%vpart%v", req.GetYear(), req.GetDay(), req.GetPart())).Call([]reflect.Value{reflect.ValueOf(ctx)})
	if val[1].Interface() == nil {
		return val[0].Interface().(*pb.SolveResponse), nil
	}
	return val[0].Interface().(*pb.SolveResponse), val[1].Interface().(error)
}
