package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/brotherlogic/goserver/utils"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"

	pb "github.com/brotherlogic/adventserver/proto"

	//Needed to pull in gzip encoding init
	_ "google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/resolver"
)

func init() {
	resolver.Register(&utils.DiscoveryClientResolverBuilder{})
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
			semconv.ServiceNameKey.String("adventserver-cli"),
			attribute.String("environment", "prod"),
			attribute.Int64("ID", 1),
		)),
	)
	return tp, nil
}

func main() {
	tp, err := tracerProvider("http://toru:14268/api/traces")
	if err != nil {
		log.Fatal(err)
	}
	otel.SetTracerProvider(tp)
	nctx, cancel := utils.ManualContext("adventserver-cli", time.Minute*5)
	defer cancel()

	ctx, span := otel.Tracer("adventserver-cli").Start(nctx, "CLI")
	defer span.End()

	conn, err := utils.LFDialServer(ctx, "adventserver")
	if err != nil {
		log.Fatalf("Unable to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewAdventServerServiceClient(conn)

	switch os.Args[1] {
	case "solve":
		t := time.Now()
		addFlags := flag.NewFlagSet("AddRecords", flag.ExitOnError)
		var year = addFlags.Int("year", -1, "Id of the record to add")
		var day = addFlags.Int("day", 0, "Cost of the record")
		var part = addFlags.Int("part", 0, "Goal folder for the record")

		if err := addFlags.Parse(os.Args[2:]); err == nil {
			if *year > 0 && *day > 0 && *part > 0 {
				res, err := client.Solve(ctx, &pb.SolveRequest{Day: int32(*day), Part: int32(*part), Year: int32(*year)})
				if err != nil {
					log.Fatalf("Error on Solve: %v", err)
				}

				fmt.Printf("Solved in %v\n", time.Since(t))
				if res.GetStringAnswer() != "" {
					fmt.Printf("%v", res.GetStringAnswer())
				} else {
					fmt.Printf("Solved: %v\n", res)
				}
			}
		}

	}

	fmt.Printf("Shutdown: %v\n", err)

}
