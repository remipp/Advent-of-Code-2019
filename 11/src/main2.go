package main

import (
	"IntCode"
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
	"image"
	"image/png"
	col "image/color"
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

var max coordinates
var min coordinates

func (r *robot) step() {
	fmt.Println("facing", r.facing)
	switch r.facing {
	case UP:
		r.position.y++
		if r.position.y > max.y {
			max.y = r.position.y
		}
	case DOWN:
		r.position.y--
		if r.position.y < min.y {
			min.y = r.position.y
		}
	case RIGHT:
		r.position.x++
		if r.position.x > max.x {
			max.x = r.position.x
		}
	case LEFT:
		r.position.x--
		if r.position.x < min.x {
			min.x = r.position.x
		}
	}
}

type direction int

func (d *direction) turn(t direction) {
	switch t {
	case LEFT:
		fmt.Println("Turning left")
		*d--
		if *d == -1 {
			*d = 3
		}
		fmt.Println(*d)
	case RIGHT:
		fmt.Println("Turning right")
		*d++
		if *d == 4 {
			*d = 0
		}
		fmt.Println(*d)
	}
}

const (
	UP direction = 0
	RIGHT direction = 1
	DOWN direction = 2
	LEFT direction = 3
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	if len(os.Args) < 3 {
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
	painter.g.panels[painter.position] = WHITE
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
	fmt.Println(painter.g.panels)
	fmt.Println(len(painter.g.panels))
	fmt.Println(max, min)
	width := Abs(min.x) + max.x
	height := Abs(min.y) + max.y
	fmt.Println("WIDTH", width, "HEIGHT", height)
	img := image.NewRGBA(image.Rect(0, 0, width+1, height+1))

	for a, b := range painter.g.panels {
		var c col.Color
		if b == WHITE {
			c = col.Gray{255}
		} else {
			c = col.Gray{0}
		}
		fmt.Println(a)
		fmt.Println("Writing pixel to", Abs(min.x)+a.x, Abs(min.y)+a.y)
		img.Set(Abs(min.x)+a.x, Abs(min.y)+a.y, c)
	}
	out, _ := os.Create(os.Args[2])
	png.Encode(out, img)
	for a, _ := range painter.g.panels {
		fmt.Println(a)
	}
}
