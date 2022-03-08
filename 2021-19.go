package main

import (
	"log"
	"strconv"
	"strings"
)

type pcoord struct {
	x    int64
	y    int64
	z    int64
	xloc int
	yloc int
	zloc int
	xt   bool
	yt   bool
	zt   bool
}

type overlap struct {
	dist int64
	i1   []int
	i2   []int
}

type match struct {
	n1  string
	n2  string
	rel *pcoord
}

func convertCoords(data string) []*pcoord {
	var coords []*pcoord
	for _, line := range strings.Split(strings.TrimSpace(data), "\n") {
		elems := strings.Split(strings.TrimSpace(line), ",")
		x, _ := strconv.Atoi(elems[0])
		y, _ := strconv.Atoi(elems[1])
		z, _ := strconv.Atoi(elems[2])
		coords = append(coords, &pcoord{x: int64(x), y: int64(y), z: int64(z)})
	}
	return coords
}

func computeDist(c1, c2 *pcoord) int64 {
	return (c1.x-c2.x)*(c1.x-c2.x) + (c1.y-c2.y)*(c1.y-c2.y) + (c1.z-c2.z)*(c1.z-c2.z)
}

func findLocation(c1, c2 []*pcoord) *pcoord {
	ret := &pcoord{}

	log.Printf("%v %v %v", c1[0].x-c2[0].x, c1[0].x-c2[0].y, c1[0].x-c2[0].z)
	log.Printf("%v %v %v", c1[1].x-c2[1].x, c1[1].x-c2[1].y, c1[1].x-c2[1].z)
	log.Printf("%v %v %v", c1[0].x+c2[0].x, c1[0].x+c2[0].y, c1[0].x+c2[0].z)
	log.Printf("%v %v %v", c1[1].x+c2[1].x, c1[1].x+c2[1].y, c1[1].x+c2[1].z)

	if c1[0].x-c2[0].x == c1[1].x-c2[1].x {
		ret.x = c1[0].x - c2[0].x
		ret.xloc = 0
		ret.xt = false
	}

	if c1[0].x+c2[0].x == c1[1].x+c2[1].x {
		ret.x = c1[0].x + c2[0].x
		ret.xloc = 0
		ret.xt = true
	}

	if c1[0].x-c2[0].y == c1[1].x-c2[1].y {
		ret.x = c1[0].x - c2[0].y
		ret.xloc = 1
		ret.xt = false
	}

	if c1[0].x+c2[0].y == c1[1].x+c2[1].y {
		ret.x = c1[0].x + c2[0].y
		ret.xloc = 1
		ret.xt = true
	}

	if c1[0].x-c2[0].z == c1[1].x-c2[1].z {
		ret.x = c1[0].x - c2[0].z
		ret.xloc = 2
		ret.xt = false
	}

	if c1[0].x+c2[0].z == c1[1].x+c2[1].z {
		ret.x = c1[0].z + c2[0].z
		ret.xloc = 2
		ret.xt = true
	}

	log.Printf("====")
	log.Printf("%v %v %v", c1[0].y-c2[0].x, c1[0].y-c2[0].y, c1[0].y-c2[0].z)
	log.Printf("%v %v %v", c1[1].y-c2[1].x, c1[1].y-c2[1].y, c1[1].y-c2[1].z)
	log.Printf("%v %v %v", c1[0].y+c2[0].x, c1[0].y+c2[0].y, c1[0].y+c2[0].z)
	log.Printf("%v %v %v", c1[1].y+c2[1].x, c1[1].y+c2[1].y, c1[1].y+c2[1].z)

	if c1[0].y-c2[0].y == c1[1].y-c2[1].y {
		ret.y = c1[0].y - c2[0].y
		ret.yloc = 1
		ret.yt = false
	}

	if c1[0].y+c2[0].y == c1[1].y+c2[1].y {
		ret.y = c1[0].y + c2[0].y
		ret.yloc = 1
		ret.yt = true
	}

	if c1[0].y-c2[0].x == c1[1].y-c2[1].x {
		ret.y = c1[0].y - c2[0].x
		ret.yloc = 0
		ret.yt = false
	}

	if c1[0].y+c2[0].x == c1[1].y+c2[1].x {
		ret.y = c1[0].y + c2[0].x
		ret.yloc = 0
		ret.yt = true
	}

	if c1[0].y-c2[0].z == c1[1].y-c2[1].z {
		ret.y = c1[0].y - c2[0].z
		ret.yloc = 2
		ret.yt = false
	}

	if c1[0].y+c2[0].z == c1[1].y+c2[1].z {
		ret.y = c1[0].y + c2[0].z
		ret.yloc = 2
		ret.yt = true
	}

	log.Printf("====")
	log.Printf("%v %v %v", c1[0].z-c2[0].x, c1[0].z-c2[0].y, c1[0].z-c2[0].z)
	log.Printf("%v %v %v", c1[1].z-c2[1].x, c1[1].z-c2[1].y, c1[1].z-c2[1].z)
	log.Printf("%v %v %v", c1[0].z+c2[0].x, c1[0].z+c2[0].y, c1[0].z+c2[0].z)
	log.Printf("%v %v %v", c1[1].z+c2[1].x, c1[1].z+c2[1].y, c1[1].z+c2[1].z)

	if c1[0].z-c2[0].z == c1[1].z-c2[1].z {
		ret.z = c1[0].z - c2[0].z
		ret.zloc = 2
		ret.zt = false
	}

	if c1[0].z+c2[0].z == c1[1].z+c2[1].z {
		ret.z = c1[0].z + c2[0].z
		ret.zloc = 2
		ret.zt = true
	}

	if c1[0].z-c2[0].x == c1[1].z-c2[1].x {
		ret.z = c1[0].z - c2[0].x
		ret.zloc = 0
		ret.zt = false
	}

	if c1[0].z+c2[0].x == c1[1].z+c2[1].x {
		ret.z = c1[0].z + c2[0].x
		ret.zloc = 0
		ret.zt = true
	}
	if c1[0].z-c2[0].y == c1[1].z-c2[1].y {
		ret.z = c1[0].z - c2[0].y
		ret.zloc = 1
		ret.zt = false
	}

	if c1[0].z+c2[0].y == c1[1].z+c2[1].y {
		ret.z = c1[0].z + c2[0].y
		ret.zloc = 1
		ret.zt = true
	}

	return ret
}

func buildOverlaps(name1, name2 string, c1, c2 []*pcoord) map[int]int {
	log.Printf("Computing overlap between %v and %v", name1, name2)
	var overlaps []*overlap
	for i := 0; i < len(c1); i++ {
		for j := i + 1; j < len(c1); j++ {
			dist := computeDist(c1[i], c1[j])
			log.Printf("DIST1: %v, %v; %v, %v -> %v", name1, name2, i, j, dist)

			found := false
			for _, overlap := range overlaps {
				if overlap.dist == dist {
					found = true
					overlap.i1 = []int{i, j}
				}
			}
			if !found {
				overlaps = append(overlaps, &overlap{dist: dist, i1: []int{i, j}})
			}
		}
	}

	for i := 0; i < len(c2); i++ {
		for j := i + 1; j < len(c2); j++ {
			dist := computeDist(c2[i], c2[j])
			if j == 25 && name2 == "--- scanner 4 ---" {
				log.Printf("DIST2: %v, %v; %v, %v -> %v", name1, name2, i, j, dist)
			}

			for _, overlap := range overlaps {
				if overlap.dist == dist {
					overlap.i2 = []int{i, j}
				}
			}
		}
	}

	overmap := make(map[int][]int)
	supermap := make(map[int]int)
	for _, overlap := range overlaps {
		log.Printf("OVERLAP %v and %v -> %v", name1, name2, overlap)
		if len(overlap.i1) > 0 && len(overlap.i2) > 0 {

			for _, elem1 := range overlap.i1 {
				if val, ok := overmap[elem1]; !ok {
					overmap[elem1] = overlap.i2
				} else {
					for _, value := range overlap.i2 {
						for _, bounce := range val {
							if value == bounce {
								supermap[elem1] = value
							}
						}
					}
				}
			}
		}
	}

	log.Printf("SMAP: (%v,%v) %v -> %v", name1, name2, supermap, len(supermap))
	return supermap
}

/*	var d1 []*pcoord
	var d2 []*pcoord
	for key, val := range supermap {
		d1 = append(d1, c1[key])
		d2 = append(d2, c2[val])
	}

	if len(d1) > 0 {
		relative := findLocation(d1, d2)
		log.Printf("MATCH %v -> %v = %+v", name1, name2, relative)

		return &match{n1: name1, n2: name2, rel: relative}
	}
	return nil
}*/

func resolveCoord(pc *pcoord, name string, matches []*match) *pcoord {
	if name == "--- scanner 0 ---" {
		return pc
	}
	for _, m := range matches {
		if m.n1 == "--- scanner 0 ---" && m.n2 == name {
			log.Printf("CONV %v %+v -> %+v", pc, name, m.rel)
			var x, y, z int64
			if m.rel.xloc == 0 {
				if m.rel.xt {
					x = m.rel.x + pc.x
				} else {
					x = m.rel.x - pc.x
				}
			}
			if m.rel.xloc == 1 {
				if m.rel.xt {
					x = m.rel.x - pc.y
				} else {
					x = m.rel.x + pc.y
				}
			}
			if m.rel.xloc == 2 {
				if m.rel.xt {
					x = m.rel.x - pc.z
				} else {
					x = m.rel.x + pc.z
				}
			}

			if m.rel.yloc == 0 {
				if m.rel.yt {
					y = m.rel.y + pc.x
				} else {
					y = m.rel.y - pc.x
				}
			}
			if m.rel.yloc == 1 {
				if m.rel.yt {
					y = m.rel.y - pc.y
				} else {
					y = m.rel.y + pc.y
				}
			}
			if m.rel.yloc == 2 {
				if m.rel.yt {
					y = m.rel.y - pc.z
				} else {
					y = m.rel.y + pc.z
				}
			}

			if m.rel.zloc == 0 {
				if m.rel.zt {
					z = m.rel.z - pc.x
				} else {
					z = m.rel.z + pc.x
				}
			}
			if m.rel.zloc == 1 {
				if m.rel.zt {
					z = m.rel.z - pc.y
				} else {
					z = m.rel.z + pc.y
				}
			}
			if m.rel.zloc == 2 {
				if m.rel.zt {
					z = m.rel.z - pc.z
				} else {
					z = m.rel.z + pc.z
				}
			}

			log.Printf("CONV TO %+v %+v", &pcoord{x: x, y: y, z: z}, m.rel)
			return &pcoord{x: x, y: y, z: z}
		}
	}
	return &pcoord{}
}

type sumap struct {
	name1 string
	name2 string
	smap  map[int]int
}

func countOverlap(data string) int {
	smap := make(map[string][]*pcoord)
	var keys []string

	currscan := ""
	var coords []*pcoord
	for _, line := range strings.Split(strings.TrimSpace(data), "\n") {
		if strings.HasPrefix(strings.TrimSpace(line), "--") {
			if len(currscan) > 0 {
				smap[currscan] = append(smap[currscan], coords...)
				keys = append(keys, currscan)
			}
			currscan = strings.TrimSpace(line)
			coords = make([]*pcoord, 0)
		} else {
			if len(strings.TrimSpace(line)) > 0 {
				coords = append(coords, convertCoords(line)...)
			}
		}
	}
	smap[currscan] = append(smap[currscan], coords...)
	keys = append(keys, currscan)

	//var matches []*match
	var ssmap []*sumap
	for i := 0; i < len(keys); i++ {
		for j := i + 1; j < len(keys); j++ {
			match := buildOverlaps(keys[i], keys[j], smap[keys[i]], smap[keys[j]])
			if match != nil {
				ssmap = append(ssmap, &sumap{name1: keys[i], name2: keys[j], smap: match})
			}
		}
	}

	var uniques []*pcoord
	for _, pcoord := range smap["--- scanner 0 ---"] {
		uniques = append(uniques, pcoord)
	}
	for i := 0; i < len(keys); i++ {
		for j := 0; j < len(smap[keys[i]]); j++ {
			found := false
			if keys[i] != "--- scanner 0 ---" {
				for _, ssm := range ssmap {
					if ssm.name2 == keys[i] {
						for _, val := range ssm.smap {
							if val == j {
								found = true
							}
						}
					}
				}
			}
			if !found {
				uniques = append(uniques, smap[keys[i]][j])
			}
		}
	}

	return len(uniques)

	/*	matches = buildMatches(keys, matches)

		for _, m := range matches {
			log.Printf("FINAL %+v -> %+v", m, m.rel)
		}

		var superlist []*pcoord
		for name, coords := range smap {
			for _, coord := range coords {
				rcoord := resolveCoord(coord, name, matches)
				log.Printf("RESOLVE %v -> %v", coord, rcoord)
				found := false
				for _, c := range superlist {
					if c.x == rcoord.x && c.y == rcoord.y && c.z == rcoord.z {
						found = true
					}
				}
				if !found {
					superlist = append(superlist, rcoord)
				}
			}
		}

		for _, elem := range superlist {
			log.Printf("SUPERLIST %+v", elem)
		}

		return len(superlist)*/
}

func buildMatches(names []string, matches []*match) []*match {
	name1 := names[0]
	for _, name2 := range names[1:] {
		found := false
		for _, m := range matches {
			if m.n1 == name1 && m.n2 == name2 {
				found = true
			}
		}

		if !found {
			log.Printf("MISSING %v -> %v", name1, name2)
			for _, interim := range names {
				var foundLeft *match
				var foundRight *match
				for _, m := range matches {
					if m.n1 == name1 && m.n2 == interim {
						foundLeft = m
					}
					if m.n1 == interim && m.n2 == name2 {
						foundRight = m
					}
				}

				if foundLeft != nil && foundRight != nil {
					matches = append(matches, resolveMatch(foundLeft, foundRight))
					break

				}
			}
		}
	}

	return matches
}

func resolveMatch(m1 *match, m2 *match) *match {
	log.Printf("AHA %+v -> %+v", m1.rel, m2.rel)

	match := &match{n1: m1.n1, n2: m2.n2, rel: &pcoord{}}

	if m2.rel.xloc == 0 {
		if m1.rel.xt != m2.rel.xt {
			match.rel.x = m1.rel.x - m2.rel.x
		} else {
			match.rel.x = m1.rel.x + m2.rel.x
		}
	}
	if m2.rel.xloc == 1 {
		if m1.rel.xt != m2.rel.yt {
			match.rel.x = m1.rel.x - m2.rel.y
		} else {
			match.rel.x = m1.rel.x + m2.rel.y
		}
	}
	if m2.rel.xloc == 2 {
		if m1.rel.xt != m2.rel.zt {
			match.rel.x = m1.rel.x - m2.rel.z
		} else {
			match.rel.x = m1.rel.x + m2.rel.z
		}
	}

	if m2.rel.yloc == 0 {
		if m1.rel.yt != m2.rel.yt {
			match.rel.y = m1.rel.y - m2.rel.x
		} else {
			match.rel.y = m1.rel.y + m2.rel.x
		}
	}

	if m2.rel.yloc == 1 {
		if m1.rel.yt != m2.rel.yt {
			match.rel.y = m1.rel.y - m2.rel.y
		} else {
			match.rel.y = m1.rel.y + m2.rel.y
		}
	}

	if m2.rel.yloc == 2 {
		if m1.rel.yt != m2.rel.yt {
			match.rel.y = m1.rel.y + m2.rel.z
		} else {
			match.rel.y = m1.rel.y - m2.rel.z
		}
	}

	if m2.rel.zloc == 0 {
		if m1.rel.zt {
			match.rel.z = m1.rel.z - m2.rel.x
		} else {
			match.rel.z = m1.rel.z + m2.rel.x
		}
	}

	if m2.rel.zloc == 1 {
		if m1.rel.zt {
			match.rel.z = m1.rel.z - m2.rel.y
		} else {
			match.rel.z = m1.rel.z + m2.rel.y
		}
	}

	if m2.rel.zloc == 2 {
		if m1.rel.zt {
			match.rel.z = m1.rel.z - m2.rel.z
		} else {
			match.rel.y = m1.rel.z + m2.rel.z
		}
	}

	if m1.rel.xloc == 0 {
		match.rel.xloc = m2.rel.xloc
	} else if m1.rel.xloc == 1 {
		match.rel.xloc = m2.rel.yloc
	} else {
		match.rel.xloc = m2.rel.zloc
	}

	if m1.rel.yloc == 1 {
		match.rel.yloc = m2.rel.yloc
	} else if m1.rel.xloc == 0 {
		match.rel.yloc = m2.rel.xloc
	} else {
		match.rel.xloc = m2.rel.zloc
	}

	if m1.rel.zloc == 2 {
		match.rel.zloc = m2.rel.zloc
	} else if m1.rel.zloc == 1 {
		match.rel.zloc = m2.rel.yloc
	} else {
		match.rel.zloc = m2.rel.xloc
	}

	match.rel.zt = true

	log.Printf("AHA2 %+v -> %+v", match, match.rel)
	return match
}
