package main

import (
	"IntCode"
	"bufio"
	"os"
	"strconv"
	"strings"
	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
	"time"
	"fmt"
)

var tile = map[int]rune{
	0: ' ',
	1: '█',
	2: '▒',
	3: '▁',
	4: '●',
}

var scr tcell.Screen

func drawTile(x, y, id int) {
	scr.SetContent(x, y, tile[id], []rune(""), tcell.StyleDefault)
	scr.Show()
	time.Sleep(time.Millisecond * 1)
}

const (
	score_x = 100
	score_y = 10
)

var sc int

func drawScore(score int) {
	x, y := score_x, score_y
	for _, c := range strconv.Itoa(score) {
		scr.SetContent(x, y, c, []rune(""), tcell.StyleDefault)
		x++
	}
}

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
	data[0] = 2
	c := IntCode.NewComputer(data)
	encoding.Register()
	scr, _ = tcell.NewScreen()
	scr.Init()
	scr.Clear()
	var writebuf []int
	var ballX, paddleX int
	*c.Write = func(i int) {
		writebuf = append(writebuf, i)
		if len(writebuf) == 3 {
			if writebuf[0] == -1 && writebuf[1] == 0 {
				drawScore(writebuf[2])
				sc = writebuf[2]
			} else {
				if writebuf[2] == 3 {
					paddleX = writebuf[0]
				}
				if writebuf[2] == 4 {
					ballX = writebuf[0]
				}
				drawTile(writebuf[0], writebuf[1], writebuf[2])
			}
			writebuf = []int{}
		}
	}
	*c.Read = func() int {
		//ev := scr.PollEvent();
		//switch ev := ev.(type) {
		//case *tcell.EventKey:
			//switch ev.Key() {
			//case tcell.KeyRune:
				//switch ev.Rune() {
					//case 'h':
						//return -1
					//case 'l':
						//return 1
				//}
			//}
		//}
		if ballX > paddleX {
			return 1
		} else if ballX < paddleX {
			return -1
		}
		return 0
	}
	c.Run()
	fmt.Println(sc)
}
