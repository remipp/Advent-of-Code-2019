package main

import (
	"fmt"
	"os"
	"bufio"
	"sort"
	"math"
)

type asteroid struct {
	x, y int
}

func (a1 asteroid) slopeTo(a2 asteroid) slope {
	return slope{a2.x-a1.x, a2.y-a1.y}
}

func (a1 asteroid) distanceTo(a2 asteroid) int {
	return Abs(a1.x - a2.x) + Abs(a1.y - a2.y)
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

func Abs(x int) int {
        if x < 0 {
                return -x
        }
        return x
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
	victims := make(map[slope]*[]asteroid)
	for _, a := range asteroids {
		if best.a == a {
			continue
		}
		s := best.a.slopeTo(a)
		var found bool
		for vS, v := range victims {
			if s.equals(vS) {
				// Insert in the right place
				for i, x := range *v {
					fmt.Println("ASTEROIDS", x, a)
					fmt.Println("DISTANCES", best.a.distanceTo(x), best.a.distanceTo(a))
					if best.a.distanceTo(x) > best.a.distanceTo(a) {
						buf := append([]asteroid{}, (*v)[:i]...)
						buf = append(buf, a)
						buf = append(buf, (*v)[i:]...)
						*v = buf
						found = true
						break
					}
				}
				if !found {
					*v = append(*v, a)
					found = true
					break
				}
			}
		}
		if !found {
			victims[s] = &[]asteroid{a}
		}
	}
	for s, v := range victims {
		fmt.Println("Slope", s)
		for _, a := range *v {
			fmt.Println(a)
		}
	}

	// Access the map in specific order -> Sort
	var buf []slope
	for a := range victims {
		buf = append(buf, a)
	}

	fmt.Println(buf)
	sort.Slice(buf, func(i, j int) bool {
		a1 := math.Atan2(float64(buf[i].y), float64(buf[i].x))
		a2 := math.Atan2(float64(buf[j].y), float64(buf[j].x))
		if a1 >= -math.Pi/2 && a1 <= 0 {
			if a2 >= -math.Pi/2 && a2 <= 0 {
				return a1 < a2
			}
			return true
		}
		if a2 >= -math.Pi/2 && a2 <= 0 {
			return false
		}
		if a1 <= math.Pi && a1 >= 0 {
			if a2 <= math.Pi && a2 >= 0 {
				return a1 < a2
			}
		}
		if a2 <= math.Pi && a2 >= 0 {
			return a1 > a2
		}
		if a1 >= -math.Pi && a1 <= 0 {
			if a2 >= -math.Pi && a2 <= 0 {
				return a1 < a2
			}
		}
		if a2 >= -math.Pi && a2 <= 0 {
			return a1 > a2
		}
		return a1 > a2
	})
	fmt.Println(buf)
	for _, a := range buf {
		fmt.Println(math.Atan2(float64(a.y), float64(a.x)), a, victims[a])
	}

	var i int = 1
	var found bool
	for !found{
		for _, a := range buf {
			fmt.Println(*victims[a])
			if len(*victims[a]) > 0 {
				fmt.Println(".")
				if i == 200 {
					fmt.Println("result", (*victims[a])[0])
					found = true
					break
				}
				fmt.Println("before", *victims[a])
				*victims[a] = (*victims[a])[1:]
				fmt.Println("after", *victims[a])
				i++
			}
		}
		fmt.Println("NEW ITERATION")
	}
}
