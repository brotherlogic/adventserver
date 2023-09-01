package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/brotherlogic/goserver/utils"

	pb "github.com/brotherlogic/adventserver/proto"

	//Needed to pull in gzip encoding init
	_ "google.golang.org/grpc/encoding/gzip"
)

func main() {
	ctx, cancel := utils.ManualContext("adventserver-cli", time.Minute*5)
	defer cancel()

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
				req := &pb.SolveRequest{Day: int32(*day), Part: int32(*part), Year: int32(*year)}
				res, err := client.Solve(ctx, req)
				if err != nil {
					log.Fatalf("Error on Solve:(from %v) %v", err, req)
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
}
