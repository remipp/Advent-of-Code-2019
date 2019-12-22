package main

import (
	"IntCode"
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
)

type color int
const (
	BLACK color = 0
	WHITE color = 1
)

type coordinates struct {
	x, y int
}

type grid struct {
	panels map[coordinates]color
}

type robot struct {
	position coordinates
	facing direction
	g grid
}

func (r *robot) step() {
	fmt.Println("facing", r.facing)
	switch r.facing {
	case UP:
		r.position.y++
	case DOWN:
		r.position.y--
	case RIGHT:
		r.position.x++
	case LEFT:
		r.position.x--
	}
}

type direction int

func (d *direction) turn(t direction) {
	switch t {
	case LEFT:
		*d--
		if *d == -1 {
			*d = 3
		}
	case RIGHT:
		*d++
		if *d == 4 {
			*d = 0
		}
	}
}

const (
	UP direction = 0
	RIGHT direction = 1
	DOWN direction = 2
	LEFT direction = 3
)

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
	var painter robot
	painter.g.panels = make(map[coordinates]color)
	*c.Read = func() int {
		fmt.Println("READ: CALLED")
		return int(painter.g.panels[painter.position])
	}
	var colorProvided bool
	*c.Write = func(i int){
		fmt.Println("WRITE: CALLED")
		if colorProvided {
			switch i {
			case 0:
				painter.facing.turn(LEFT)
			case 1:
				painter.facing.turn(RIGHT)
			}
			colorProvided = false
			painter.step()
		} else {
			fmt.Println("position", painter.position)
			fmt.Println("before", painter.g.panels[painter.position])
			painter.g.panels[painter.position] = color(i)
			fmt.Println("after", painter.g.panels[painter.position])
			colorProvided = true
		}
	}
	c.Run()
	fmt.Println(len(painter.g.panels))
}
