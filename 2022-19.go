package main

import (
	"fmt"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"golang.org/x/net/context"
)

type blueprint struct {
	ore, clay       int
	obsOre, obsClay int
	geoOre, geoObs  int
}

func (b blueprint) maxOre() int {
	return max(b.ore, max(b.clay, max(b.obsOre, b.geoOre)))
}

type robotNode struct {
	ore, clay, obs, geo                     int
	oreRobot, clayRobot, obsRobot, geoRobot int
	minutes                                 int
	built                                   string
	builtTimes                              string
}

func (r robotNode) getRep() string {
	return fmt.Sprintf("%v-%v-%v-%v;%v-%v-%v-%v;%v", r.ore, r.clay, r.obs, r.geo, r.oreRobot, r.clayRobot, r.obsRobot, r.geoRobot, r.minutes)
}

func (r robotNode) getBuilt() string {
	return r.built
}

func buildBluePrint(data string) blueprint {
	fields := strings.Fields(data)
	return blueprint{
		ore:     getInt32(fields[6]),
		clay:    getInt32(fields[12]),
		obsOre:  getInt32(fields[18]),
		obsClay: getInt32(fields[21]),
		geoOre:  getInt32(fields[27]),
		geoObs:  getInt32(fields[30])}
}

func low(n int) int {
	if n < 0 {
		return 1
	}
	return n
}

func ceil(a, b int) int {
	if a%b == 0 {
		return a / b
	}

	return a/b + 1
}

func buildNexts(bp blueprint, rn robotNode) []robotNode {

	var next []robotNode

	// Build an Ore robot
	if rn.minutes-bp.ore/rn.oreRobot > 0 {
		timeTaken := (low(bp.ore-rn.ore) / rn.oreRobot) + 1
		next = append(next, robotNode{
			ore:       rn.ore + rn.oreRobot*timeTaken - bp.ore,
			clay:      rn.clay + rn.clayRobot*timeTaken,
			obs:       rn.obs + rn.obsRobot*timeTaken,
			geo:       rn.geo + rn.geoRobot*timeTaken,
			oreRobot:  rn.oreRobot + 1,
			clayRobot: rn.clayRobot,
			obsRobot:  rn.obsRobot,
			geoRobot:  rn.geoRobot,
			minutes:   rn.minutes - timeTaken,
			built:     rn.built + "ORE-",
		})

	}

	// Build a clay robot
	if rn.minutes-bp.clay/rn.oreRobot > 0 {
		timeTaken := (low(bp.clay-rn.ore) / rn.oreRobot) + 1
		next = append(next, robotNode{
			ore:       rn.ore + rn.oreRobot*timeTaken - bp.clay,
			clay:      rn.clay + rn.clayRobot*timeTaken,
			obs:       rn.obs + rn.obsRobot*timeTaken,
			geo:       rn.geo + rn.geoRobot*timeTaken,
			oreRobot:  rn.oreRobot,
			clayRobot: rn.clayRobot + 1,
			obsRobot:  rn.obsRobot,
			geoRobot:  rn.geoRobot,
			minutes:   rn.minutes - timeTaken,
			built:     rn.built + "CLAY-",
		})
	}

	// Build an obsidian robot
	if rn.oreRobot > 0 && rn.clayRobot > 0 {
		timeTakenObs := max(ceil(low(bp.obsOre-rn.ore), rn.oreRobot), ceil(low(bp.obsClay-rn.clay), rn.clayRobot)) + 1
		if timeTakenObs < rn.minutes {
			next = append(next, robotNode{
				ore:       rn.ore + rn.oreRobot*timeTakenObs - bp.obsOre,
				clay:      rn.clay + rn.clayRobot*timeTakenObs - bp.obsClay,
				obs:       rn.obs + rn.obsRobot*timeTakenObs,
				geo:       rn.geo + rn.geoRobot*timeTakenObs,
				oreRobot:  rn.oreRobot,
				clayRobot: rn.clayRobot,
				obsRobot:  rn.obsRobot + 1,
				geoRobot:  rn.geoRobot,
				minutes:   rn.minutes - timeTakenObs,
				built:     rn.built + "OBS-",
			})
		}

	}

	// Build an geode robot
	if rn.oreRobot > 0 && rn.obsRobot > 0 {
		timeTakenGeo := max(ceil(low(bp.geoOre-rn.ore), rn.oreRobot), ceil(low(bp.geoObs-rn.obs), rn.obsRobot)) + 1
		if timeTakenGeo < rn.minutes {
			next = append(next, robotNode{
				ore:       rn.ore + rn.oreRobot*timeTakenGeo - bp.geoOre,
				clay:      rn.clay + rn.clayRobot*timeTakenGeo,
				obs:       rn.obs + rn.obsRobot*timeTakenGeo - bp.geoObs,
				geo:       rn.geo + rn.geoRobot*timeTakenGeo,
				oreRobot:  rn.oreRobot,
				clayRobot: rn.clayRobot,
				obsRobot:  rn.obsRobot,
				geoRobot:  rn.geoRobot + 1,
				minutes:   rn.minutes - timeTakenGeo,
				built:     rn.built + "GEO-",
			})
		}

	}

	// Run out the clock if we need to
	if len(next) == 0 {
		next = append(next, robotNode{
			geo:   rn.geo + rn.geoRobot*(rn.minutes),
			built: rn.built,
		})
	}

	return next
}

func buildNextsLong(bp blueprint, rn robotNode, bests map[string]int) []robotNode {

	var next []robotNode

	// Build an geode robot
	if rn.ore >= bp.geoOre && rn.obs >= bp.geoObs {
		node := robotNode{
			ore:        rn.ore + rn.oreRobot - bp.geoOre,
			clay:       rn.clay + rn.clayRobot,
			obs:        rn.obs + rn.obsRobot - bp.geoObs,
			geo:        rn.geo + rn.geoRobot,
			oreRobot:   rn.oreRobot,
			clayRobot:  rn.clayRobot,
			obsRobot:   rn.obsRobot,
			geoRobot:   rn.geoRobot + 1,
			minutes:    rn.minutes - 1,
			built:      rn.built + "GEO-",
			builtTimes: rn.builtTimes + fmt.Sprintf("%v:%v-", "GEO", rn.minutes-1),
		}
		if val, ok := bests[node.getBuilt()]; !ok || node.minutes > val {
			next = append(next, node)
		}
	}

	// Build an obsidian robot
	if rn.ore >= bp.obsOre && rn.clay >= bp.obsClay {
		node := robotNode{
			ore:        rn.ore + rn.oreRobot - bp.obsOre,
			clay:       rn.clay + rn.clayRobot - bp.obsClay,
			obs:        rn.obs + rn.obsRobot,
			geo:        rn.geo + rn.geoRobot,
			oreRobot:   rn.oreRobot,
			clayRobot:  rn.clayRobot,
			obsRobot:   rn.obsRobot + 1,
			geoRobot:   rn.geoRobot,
			minutes:    rn.minutes - 1,
			built:      rn.built + "OBS-",
			builtTimes: rn.builtTimes + fmt.Sprintf("%v:%v-", "OBS", rn.minutes-1),
		}
		if val, ok := bests[node.getBuilt()]; !ok || node.minutes > val {
			next = append(next, node)
		}
	}

	// Build a clay robot
	if rn.ore >= bp.clay && rn.clayRobot < bp.obsClay {
		node := robotNode{
			ore:        rn.ore + rn.oreRobot - bp.clay,
			clay:       rn.clay + rn.clayRobot,
			obs:        rn.obs + rn.obsRobot,
			geo:        rn.geo + rn.geoRobot,
			oreRobot:   rn.oreRobot,
			clayRobot:  rn.clayRobot + 1,
			obsRobot:   rn.obsRobot,
			geoRobot:   rn.geoRobot,
			minutes:    rn.minutes - 1,
			built:      rn.built + "CLAY-",
			builtTimes: rn.builtTimes + fmt.Sprintf("%v:%v-", "CLAY", rn.minutes-1),
		}
		if val, ok := bests[node.getBuilt()]; !ok || node.minutes > val {
			next = append(next, node)
		}
	}

	// Build an Ore robot
	if rn.ore >= bp.ore && rn.oreRobot < bp.maxOre() {
		node := robotNode{
			ore:        rn.ore + rn.oreRobot - bp.ore,
			clay:       rn.clay + rn.clayRobot,
			obs:        rn.obs + rn.obsRobot,
			geo:        rn.geo + rn.geoRobot,
			oreRobot:   rn.oreRobot + 1,
			clayRobot:  rn.clayRobot,
			obsRobot:   rn.obsRobot,
			geoRobot:   rn.geoRobot,
			minutes:    rn.minutes - 1,
			built:      rn.built + "ORE-",
			builtTimes: rn.builtTimes + fmt.Sprintf("%v:%v-", "ORE", rn.minutes-1),
		}
		if val, ok := bests[node.getBuilt()]; !ok || node.minutes > val {
			next = append(next, node)
		}
	}

	// Do nothing
	next = append(next, robotNode{
		ore:        rn.ore + rn.oreRobot,
		clay:       rn.clay + rn.clayRobot,
		obs:        rn.obs + rn.obsRobot,
		geo:        rn.geo + rn.geoRobot,
		oreRobot:   rn.oreRobot,
		clayRobot:  rn.clayRobot,
		obsRobot:   rn.obsRobot,
		geoRobot:   rn.geoRobot,
		minutes:    rn.minutes - 1,
		built:      rn.built,
		builtTimes: rn.builtTimes,
	})

	return next
}

func runNode(bp blueprint, rn robotNode) int {

	queue := []robotNode{rn}
	bestRes := 0
	seen := make(map[string]bool)
	bests := make(map[string]int)

	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]

		if head.minutes == 0 {
			if head.geo > bestRes {
				bestRes = head.geo
			}
			continue
		}

		nexts := buildNextsLong(bp, head, bests)
		for _, n := range nexts {
			if _, ok := seen[n.getRep()]; !ok {
				seen[n.getRep()] = true
				queue = append(queue, n)
				if val, ok := bests[n.getBuilt()]; !ok || n.minutes > val {
					bests[n.getBuilt()] = n.minutes
				}
			}
		}
	}

	return bestRes
}

func getBestBlue(data string) int {
	bestVal := 0

	for i, line := range strings.Split(strings.TrimSpace(data), "\n") {
		bp := buildBluePrint(strings.TrimSpace(line))
		best := runNode(bp, robotNode{ore: 1, oreRobot: 1, minutes: 23})
		bestVal += (i + 1) * best
	}

	return bestVal
}

func (s *Server) Solve2022day19part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2022-19.txt")
	if err != nil {
		return nil, err
	}

	return &pb.SolveResponse{Answer: int32(getBestBlue(data))}, nil
}
