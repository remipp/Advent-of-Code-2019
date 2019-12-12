package main
import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type path []point

func toPath(wire []string) path {
	//var current path
	current := make([]point, 0, len(wire))
	var buf point
	for _, a := range wire {
		x, _ := strconv.Atoi(a[1:])
		switch(a[0]) {
		case 'U':
			buf.y += x
		case 'D':
			buf.y -= x
		case 'R':
			buf.x += x
		case 'L':
			buf.x -= x
		}
		current = append(current, buf)
	}
	return path(current)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (p1 path) checkIntersections(p2 path) []int {
	//var buf []point
	var buf []int
	for i := 1; i < len(p1); i++ {
		for k := 1; k < len(p2); k++ {
			// vertical
			if p1[i].x == p1[i-1].x {
				// also vertical
				if p2[k].x == p2[k-1].x {
					continue
				}
				// intersecting
				if ((p2[k].y >= p1[i].y && p2[k].y <= p1[i-1].y) || (p2[k].y <= p1[i].y && p2[k].y >= p1[i-1].y)) &&
				   ((p2[k].x <= p1[i].x && p2[k-1].x >= p1[i].x) || (p2[k].x >= p1[i].x && p2[k-1].x <= p1[i].x)) {
					//buf = append(buf, point{p1[i].x, p2[k].y})
					// p1 steps until intersection at p1[i].x and p2[k].y
					var steps int
					// x < i: don't include last one, we will add p2's y here
					steps += Abs(p1[0].x) + Abs(p1[0].y)
					for x := 1; x < i; x++ {
						steps += Abs(p1[x].x - p1[x-1].x) + Abs(p1[x].y - p1[x-1].y)
					}
					steps += Abs(p2[k].y - p1[i-1].y)

					// same for p2
					steps += Abs(p2[0].x) + Abs(p2[0].y)
					for x := 1; x < k; x++ {
						steps += Abs(p2[x].x - p2[x-1].x) + Abs(p2[x].y - p2[x-1].y)
					}
					steps += Abs(p1[i].x - p2[k-1].x)
					buf = append(buf, steps)
				}
			} else { // horizontal
				if (p2[k].x != p2[k-1].x) { // also horizontal
					continue
				}
				// intersecting
				if ((p2[k].x >= p1[i].x && p2[k].x <= p1[i-1].x) || (p2[k].x <= p1[i].x && p2[k].x >= p1[i-1].x)) &&
				   ((p2[k].y <= p1[i].y && p2[k-1].y >= p1[i].y) || (p1[k].y >= p1[i].y && p2[k-1].y <= p1[i].y)) {
					// Like before but other way around(p1 horizontal, p2 vertical)
					//buf = append(buf, point{p2[k].x, p1[i].y})
					var steps int
					steps += Abs(p2[0].x) + Abs(p2[0].y)
					for x := 1; x < k; x++ {
						steps += Abs(p2[x].x - p2[x-1].x) + Abs(p2[x].y - p2[x-1].y)
					}
					steps += Abs(p1[i].y - p2[k-1].y)

					// same for p1
					steps += Abs(p1[0].x) + Abs(p1[0].y)
					for x := 1; x < i; x++ {
						steps += Abs(p1[x].x - p1[x-1].x) + Abs(p1[x].y - p1[x-1].y)
					}
					steps += Abs(p2[k].x - p1[i-1].x)
					buf = append(buf, steps)
				}
			}
		}
	}
	return buf
}

func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	p := toPath(strings.Split(scanner.Text(), ","))
	scanner.Scan()
	fmt.Println(p)
	res := p.checkIntersections(toPath(strings.Split(scanner.Text(), ",")))
	fmt.Println(res)
	var best int = 9999
	for _, a := range res {
		if a < best {
			best = a
		}
	}
	fmt.Println(best)
}
