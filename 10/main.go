package main

import (
	"fmt"
	"os"
	"bufio"
)

type asteroid struct {
	x, y int
}

func (a1 asteroid) slopeTo(a2 asteroid) slope {
	return slope{a2.x-a1.x, a2.y-a1.y}
}

type slope struct {
	x, y int
}

func (s1 slope) equals(s2 slope) bool {
	if (s1.x > 0 && s2.x < 0 || s1.x < 0 && s2.x > 0) || (s1.y > 0 && s2.y < 0 || s1.y < 0 && s2.y > 0) {
		return false
	}
	s1.x *= s2.y
	s2.x *= s1.y
	return s1.x == s2.x
}

type result struct {
	a asteroid
	inSight int
}

func main() {
	if len(os.Args) < 2 {
		os.Exit(-1)
	}
	file, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(file)

	var asteroids []asteroid

	var y int
	for scanner.Scan() {
		var x int
		for _, c := range scanner.Text() {
			if c == '#' {
				asteroids = append(asteroids, asteroid{x, y})
			}
			x++
		}
		y++
	}

	var best result
	for i := 0; i < len(asteroids); i++ {
		var knownSlopes []slope
		for j := 0; j < len(asteroids); j++ {
			s := asteroids[i].slopeTo(asteroids[j])
			unique := true
			for _, a := range knownSlopes {
				if a.equals(s) {
					unique = false
					break
				}
			}
			if unique {
				knownSlopes = append(knownSlopes, s)
			}
		}
		if len(knownSlopes) > best.inSight {
			best.a = asteroids[i]
			best.inSight = len(knownSlopes)
		}
	}
	fmt.Println(best)
}
