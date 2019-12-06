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

func (p1 path) checkIntersections(p2 path) []point {
	var buf []point
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
					buf = append(buf, point{p1[i].x, p2[k].y})
				}
			} else { // horizontal
				if (p2[k].x != p2[k-1].x) { // also horizontal
					continue
				}
				// intersecting
				if ((p2[k].x >= p1[i].x && p2[k].x <= p1[i-1].x) || (p2[k].x <= p1[i].x && p2[k].x >= p1[i-1].x)) &&
				   ((p2[k].y <= p1[i].y && p2[k-1].y >= p1[i].y) || (p1[k].y >= p1[i].y && p2[k-1].y <= p1[i].y)) {
					buf = append(buf, point{p2[k].x, p1[i].y})
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
	res := p.checkIntersections(toPath(strings.Split(scanner.Text(), ",")))
	var best int = 9999
	for _, a := range res {
		md := int(Abs(a.x)+Abs(a.y))
		if md < best {
			best = md
		}
	}
	fmt.Println(best)
}
