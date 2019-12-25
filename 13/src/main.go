package main

import (
	"IntCode"
	"bufio"
	"os"
	"strconv"
	"strings"
	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
	"fmt"
	"time"
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
	encoding.Register()
	scr, _ = tcell.NewScreen()
	scr.Init()
	scr.Clear()
	var writebuf []int
	var ctr int
	*c.Write = func(i int) {
		writebuf = append(writebuf, i)
		if len(writebuf) == 3 {
			if writebuf[2] == 2 {
				ctr++
			}
			drawTile(writebuf[0], writebuf[1], writebuf[2])
			writebuf = []int{}
		}
	}
	c.Run()
	fmt.Println(ctr)
}
