package main

import (
	"log"
	"math"
	"strings"

	pb "github.com/brotherlogic/adventserver/proto"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"golang.org/x/net/context"
)

var (
	searches = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "adventserver_day19_searches",
		Help: "The number of server requests",
	})
	msize = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "adventserver_day19_msize",
		Help: "The number of server requests",
	})
)

func buildMaps(data string) (map[string][]string, string) {
	done := false
	mapper := make(map[string][]string)
	key := ""
	for _, line := range strings.Split(data, "\n") {
		if len(strings.TrimSpace(line)) == 0 {
			done = true
		} else {
			if !done {
				elems := strings.Fields(line)
				mapper[elems[0]] = append(mapper[elems[0]], elems[2])
			} else {
				key = line
			}
		}
	}

	return mapper, strings.TrimSpace(key)
}

func buildString(parts []string, pos int, from, adj string) string {
	nstr := ""
	for i := 0; i < pos; i++ {
		nstr += parts[i] + from
	}
	nstr += parts[pos] + adj
	for i := pos + 1; i < len(parts)-1; i++ {
		nstr += parts[i] + from
	}
	nstr += parts[len(parts)-1]
	return nstr
}

func translate(key string, trans map[string][]string) map[string]bool {
	result := make(map[string]bool)

	for adj, tos := range trans {
		for _, to := range tos {
			parts := strings.Split(key, adj)
			for i := 0; i < len(parts)-1; i++ {
				nstr := buildString(parts, i, adj, to)
				result[nstr] = true
			}
		}
	}

	//log.Printf("RES %v", result)
	return result
}

func treeMolecules(data string) int {
	trans, key := buildMaps(data)

	flipped := make(map[string]string)
	for k, v := range trans {
		for _, vv := range v {
			flipped[vv] = k
		}
	}

	//res := runBackwards(key, "e", flipped, 0)
	seen := make(map[string]tracker)
	seen[key] = tracker{cost: 0, processed: false}
	res := runSearch(seen, "e", flipped)

	return res
}

type tracker struct {
	cost      int
	processed bool
}

func runSearch(seen map[string]tracker, goal string, trans map[string]string) int {
	searches.Inc()
	msize.Set(float64(len(seen)))
	if val, ok := seen[goal]; ok {
		return val.cost
	}

	best := ""
	bv := math.MaxInt
	for key, val := range seen {
		if val.cost < bv && !val.processed {
			best = key
			bv = val.cost
		}
	}

	delete(seen, best)

	for key, val := range trans {
		indices := getIndices(key, best)
		for _, index := range indices {
			nstr := best[:index] + val + best[index+len(key):]
			if strings.Count(nstr, "e") == 0 || (strings.Count(nstr, "e") == 1 && len(nstr) == 1) {
				if _, ok := seen[nstr]; !ok {
					seen[nstr] = tracker{cost: bv + 1, processed: false}
				}
			}
		}
	}

	return runSearch(seen, goal, trans)
}

func runBackwards(current, goal string, trans map[string]string, count int) int {
	log.Printf("%v", current)
	searches.Inc()
	if current == goal {
		return count
	}

	if strings.Contains(current, "e") {
		return math.MaxInt
	}

	log.Printf("%v", current)

	best := math.MaxInt
	for key, res := range trans {
		indices := getIndices(key, current)
		for _, index := range indices {
			nstr := current[:index] + res + current[index+len(key):]
			nval := runBackwards(nstr, goal, trans, count+1)

			if nval < best {
				best = nval
			}
		}
	}

	return best
}

func getIndices(key string, lon string) []int {
	var indices []int
	for i := 0; i <= len(lon)-len(key); i++ {
		if lon[i:i+len(key)] == key {
			indices = append(indices, i)
		}
	}

	return indices
}

func runMTree(sofar string, key string, trans map[string][]string, count int, seen map[string]bool) (int, map[string]bool) {
	if _, ok := seen[sofar]; ok {
		return math.MaxInt, seen
	}
	//log.Printf("SEARCH %v", sofar)

	seen[sofar] = true

	searches.Inc()
	if len(sofar) > len(key) {
		return math.MaxInt, seen
	}

	if sofar == key {
		return count, seen
	}

	best := math.MaxInt
	for tkey, transl := range trans {
		for _, val := range getIndices(tkey, sofar) {
			for _, tval := range transl {
				nstring := sofar[:val] + tval + sofar[val+len(tkey):]
				nval, nseen := runMTree(nstring, key, trans, count+1, seen)
				seen = nseen
				if nval < best {
					best = nval
				}
			}
		}
	}

	return best, seen
}

func getMolecules(data string) int {
	trans, key := buildMaps(data)

	res := translate(key, trans)

	return len(res)
}

func (s *Server) Solve2015day19part1(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2015-19.txt")
	if err != nil {
		return nil, err
	}

	trimmed := strings.TrimSpace(data)

	return &pb.SolveResponse{Answer: int32(getMolecules(trimmed))}, nil
}

func (s *Server) Solve2015day19part2(ctx context.Context) (*pb.SolveResponse, error) {
	data, err := s.loadFile(ctx, "/media/scratch/advent/2015-19.txt")
	if err != nil {
		return nil, err
	}

	trimmed := strings.TrimSpace(data)

	return &pb.SolveResponse{Answer: int32(treeMolecules(trimmed))}, nil
}
