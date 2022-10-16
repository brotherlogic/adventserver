package main

import (
	"log"
	"strconv"
	"strings"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

type props struct {
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func buildProps(data string) []props {
	var ps []props
	for _, l := range strings.Split(data, "\n") {
		bah := strings.Split(l, ":")
		elems := strings.Split(bah[1], ",")
		p := props{}
		for _, elem := range elems {
			bits := strings.Fields(elem)
			value, _ := strconv.ParseInt(bits[1], 10, 32)
			switch bits[0] {
			case "capacity":
				p.capacity = int(value)
			case "durability":
				p.durability = int(value)
			case "flavor":
				p.flavor = int(value)
			case "texture":
				p.texture = int(value)
			case "calories":
				p.calories = int(value)
			}
		}

		ps = append(ps, p)
	}

	return ps
}

func findBestIng(ap []props, maxv int, calories int) int {
	return findBestWith(make([]int, 0), ap, maxv, calories)
}

func findBestWith(sofar []int, ap []props, maxv int, caloriesGoal int) int {
	if len(sofar) == len(ap) {
		capacity := 0
		durability := 0
		flavor := 0
		texture := 0
		calories := 0

		for i := range ap {
			capacity += sofar[i] * ap[i].capacity
			durability += sofar[i] * ap[i].durability
			flavor += sofar[i] * ap[i].flavor
			texture += sofar[i] * ap[i].texture
			calories += sofar[i] * ap[i].calories
		}

		if sofar[0] == 44 {
			log.Printf("%v %v %v %v %+v %+v", capacity, durability, flavor, texture, ap[0], ap[1])
		}

		if capacity < 0 || durability < 0 || flavor < 0 || texture < 0 {
			return 0
		}

		if calories != caloriesGoal && caloriesGoal >= 0 {
			return 0
		}

		return capacity * durability * flavor * texture
	}

	if len(sofar) == len(ap)-1 {
		nv := maxv
		for _, c := range sofar {
			nv -= c
		}

		if nv <= 0 {
			return 0
		}

		sofar = append(sofar, nv)
		return findBestWith(sofar, ap, maxv, caloriesGoal)
	}

	best := 0
	sofar = append(sofar, 0)
	for i := 1; i < 100; i++ {
		sofar[len(sofar)-1] = i
		b := findBestWith(sofar, ap, maxv, caloriesGoal)
		if b > best {
			best = b
		}
	}

	return best
}

func computeBestScore(tp *tracesdk.TracerProvider, ctx context.Context, data string, calories int) int {
	_, span := tp.Tracer("2015-15-2").Start(ctx, "Run")
	defer span.End()

	allProps := buildProps(data)

	best := findBestIng(allProps, 100, calories)

	return best
}

func (s *Server) Solve2015day15part1(ctx context.Context) (*pb.SolveResponse, error) {
	tp, err := tracerProvider("http://toru:14268/api/traces")
	if err != nil {
		log.Fatal(err)
	}
	otel.SetTracerProvider(tp)

	data, err := s.loadFile(ctx, "/media/scratch/advent/2015-15.txt")
	if err != nil {
		return nil, err
	}

	trimmed := strings.TrimSpace(data)

	return &pb.SolveResponse{Answer: int32(computeBestScore(tp, ctx, trimmed, -1))}, nil
}

func tracerProvider(url string) (*tracesdk.TracerProvider, error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("adventserver"),
			attribute.String("environment", "prod"),
			attribute.Int64("ID", 1),
		)),
	)
	return tp, nil
}

func (s *Server) Solve2015day15part2(ctx context.Context) (*pb.SolveResponse, error) {
	tp, err := tracerProvider("http://toru:14268/api/traces")
	if err != nil {
		log.Fatal(err)
	}
	otel.SetTracerProvider(tp)

	newCtx, span := tp.Tracer("2015-15-2").Start(ctx, "Run")
	defer span.End()

	data, err := s.loadFile(newCtx, "/media/scratch/advent/2015-15.txt")
	if err != nil {
		return nil, err
	}

	trimmed := strings.TrimSpace(data)

	return &pb.SolveResponse{Answer: int32(computeBestScore(tp, ctx, trimmed, 500))}, nil
}
