package main

import (
	"IntCode"
	"os"
	"bufio"
	"strconv"
	"strings"
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
	c.Run()
}
