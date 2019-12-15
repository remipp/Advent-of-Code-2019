package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"math"
)

type planet struct {
	orbitees []*planet
}


func doTheThing(p1, p2 *planet) int {
	var best int = math.MaxInt32
	doTheThingHelper(p1, p2, 0, &best)
	return best
}

func doTheThingHelper(p1, p2 *planet, ctr int, best *int) {
	for _, o := range p1.orbitees {
		if x, ok := transfersRequired(p2, o); ok {
			if x+ctr < *best {
				*best = x+ctr
			}
		}
		doTheThingHelper(o, p2, ctr+1, best)
	}
}

func transfersRequired(p, destination *planet) (int, bool) {
	var best int = math.MaxInt32
	var found bool
	transfersRequiredHelper(p, destination, 0, &best, &found)
	return best, found
}

func transfersRequiredHelper(p, destination *planet, steps int, best *int, found *bool) {
	for _, o := range p.orbitees {
		if o == destination {
			*found = true
			if steps < *best {
				*best = steps
			}
		}
		transfersRequiredHelper(o, destination, steps+1, best, found)
	}
}

func main() {
	if len(os.Args) < 2 {
		os.Exit(-1)
	}
	file, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(file)

	universe := make(map[string]*planet)
	for scanner.Scan() {
		buf := strings.Split(scanner.Text(), ")")
		if _, ok := universe[buf[0]]; !ok {
			universe[buf[0]] = new(planet)
		}
		if _, ok := universe[buf[1]]; !ok {
			universe[buf[1]] = new(planet)
		}
		universe[buf[1]].orbitees = append(universe[buf[1]].orbitees, universe[buf[0]])
	}
	fmt.Println(doTheThing(universe["YOU"], universe["SAN"]))
}
