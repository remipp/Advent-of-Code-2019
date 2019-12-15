package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

type planet struct {
	orbiters []*planet
}

var ctr int
func doTheThing(p *planet, worth int) {
	ctr += worth
	for _, o := range p.orbiters {
		doTheThing(o, worth+1)
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
		universe[buf[0]].orbiters = append(universe[buf[0]].orbiters, universe[buf[1]])
	}

	doTheThing(universe["COM"], 0)
	fmt.Println(ctr)
}
