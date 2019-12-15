package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(-1)
	}
	file, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(file)
	scanner.Split(func (data []byte, atEOF bool) (int, []byte, error) {
		if !atEOF {
			return 0, nil, nil
		}
		var buf []byte
		var i int
		var x byte
		for i, x = range data {
			if x == ',' {
				break
			}
			buf = append(buf, x)
		}
		return i+1, buf, nil
	})
	var data []int
	for scanner.Scan() {
		x, _ := strconv.Atoi(scanner.Text())
		data = append(data, x)
	}
	data[1] = 12
	data[2] = 2
	for i := 0; i < len(data); i += 4 {
		fmt.Println(data[i])
		switch(data[i]) {
		case 1:
			data[data[i+3]] = data[data[i+1]] + data[data[i+2]]
		case 2:
			data[data[i+3]] = data[data[i+1]] * data[data[i+2]]
		case 99:
			break
		}
	}
	fmt.Println(data[00])
}
