package main

import (
	"IntCode"
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
	"github.com/gdamore/tcell"
	//"github.com/gdamore/tcell/encoding"
	"time"
)

var scr tcell.Screen

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func draw(x int, y int, c rune) {
	scr.SetContent(80+x, 20+y, rune(c), []rune(""), tcell.StyleDefault)
	scr.Show()
	time.Sleep(10*time.Millisecond)
}

type coordinate struct {
	x, y int
}

func (c coordinate) apply(m int) coordinate {
	switch m {
	case 1:
		c.y++
	case 2:
		c.y--
	case 3:
		c.x--
	case 4:
		c.x++
	}
	return c
}

type node struct {
	position coordinate // will be node identifier here
	neighbours []*node
	// Needed for BFS
	previous *node
	visited bool
}

type maze struct {
	nodes []node
	walls map[coordinate]bool // Stored so the Write method has a way to check for previously tried coordinates
}

func (m maze) createCopy() *maze {
	var buf maze
	buf.nodes = m.nodes
	for i := 0; i < len(buf.nodes); i++ {
		// Make pointers to neighbours point to nodes in new object
		var newNeighbours []*node
		for _, nb := range buf.nodes[i].neighbours {
			if nb == nil {
				continue
			}
			newNeighbours = append(newNeighbours, buf.getNode(nb.position))
		}
		buf.nodes[i].neighbours = newNeighbours
		// Previous pointer doesn't matter
		buf.nodes[i].visited = false
	}
	return &buf
}

func (m maze) String() string {
	var buf string
	for _, node := range m.nodes {
		buf = fmt.Sprint(buf, node.position, ":")
		for _, neighbour := range node.neighbours {
			if neighbour == nil {
				continue
			}
			buf = fmt.Sprint(buf, " ", neighbour.position)
		}
		buf = fmt.Sprintln(buf)
	}
	return buf
}

func (m maze) contains(c coordinate) bool {
	for _, n := range m.nodes {
		if n.position == c {
			return true
		}
	}
	if _, ok := m.walls[c]; ok {
		return true
	}
	return false
}

func (m maze) getNode(c coordinate) *node {
	for i := 0; i < len(m.nodes); i++ {
		if m.nodes[i].position == c {
			return &m.nodes[i]
		}
	}
	return nil
}

func (m maze) bfs(start *node, dest *node) []*node {
	var q []*node
	q = []*node{start}
	start.visited = true
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		if v.position == dest.position {
			var path []*node
			for {
				path = append(path, v)
				v = v.previous
				if v == start {
					break
				}
			}
			return path
		}
		for _, n := range v.neighbours {
			if n == nil || n.visited {
				continue
			}
			n.visited = true
			n.previous = v
			q = append(q, n)
		}
	}
	return nil
}

func (m *maze) reset() {
	for i := 0; i < len(m.nodes); i++ {
		m.nodes[i].visited = false
		m.nodes[i].previous = nil
	}
}

type move struct {
	directionsTried []int
	lastDirection int // The move that got us to this position
}

var curPos coordinate
var m maze

var directions = map[int][]int {
	1: []int{1, 3, 4, 2},
	2: []int{2, 3, 4, 1},
	3: []int{3, 1, 2, 4},
	4: []int{4, 1, 2, 3},
}

// Try inverse last (backtracking at the end)
func main() {
	if len(os.Args) < 2 {
		os.Exit(-1)
	}
	file, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(file)
	var data []int
	scanner.Scan()
	buf := strings.Split(scanner.Text(), ",")
	for _, x := range buf {
		n, _ := strconv.Atoi(x)
		data = append(data, n)
	}
	c := IntCode.NewComputer(data)


	m.walls = make(map[coordinate]bool)

	startPos := coordinate{0, 0}
	curPos := startPos
	lastPos := startPos

	moves := make(map[coordinate]move)
	moves[startPos] = move{
		[]int{},
		1,
	}

	//encoding.Register()
	scr, _ = tcell.NewScreen()
	//scr.Init()
	//scr.Clear()

	var generator coordinate

	*c.Write = func(i int) {
		// Response: Status of last movement command
		switch i {
		case 0:
			m.walls[curPos] = true
			// Go back to last Pos
			curPos = lastPos
		case 2:
			generator = curPos
			fallthrough
		case 1:
			if n := m.getNode(curPos); n == nil {
				var buf node
				buf.position = curPos
				buf.neighbours = append(buf.neighbours, m.getNode(lastPos))
				m.nodes = append(m.nodes, buf)
			} else {
				if l := m.getNode(lastPos); l != nil {
					n.neighbours = append(n.neighbours, l)
				}
			}
		}
	}
	*c.Read = func() int {
		// Tell robot next move
		i := len(moves[curPos].directionsTried)
		if i == 4 {
			c.Halt()
			return 1
		}
		blah := directions[moves[curPos].lastDirection]
		nextMove := blah[i]
		nextPos := curPos.apply(nextMove)
		buf := moves[curPos]
		buf.directionsTried = append(moves[curPos].directionsTried, nextMove)
		moves[curPos] = buf
		if _, ok := moves[nextPos]; !ok {
			moves[nextPos] = move{
				[]int{},
				nextMove,
			}
		}
		lastPos = curPos
		curPos = nextPos
		return nextMove
	}
	c.Run()
	// Find path
	var max int
	for i := 0; i < len(m.nodes); i++ {
		cpy := m.createCopy()
		if generator == m.nodes[i].position { // BFS crashes if start and destination are equal
			continue
		}
		val := len(cpy.bfs(&cpy.nodes[i], cpy.getNode(generator)))
		if val > max {
			max = val
		}
	}
	fmt.Println(max)
}
