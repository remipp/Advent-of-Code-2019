package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
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
	for i := 0; i < 100; i++ {
		for k := 0; k < 100; k++ {
			buf := append([]int{}, data...)
			buf[1] = i
			buf[2] = k
			for i := 0; i < len(buf); i += 4 {
				switch(buf[i]) {
				case 1:
					buf[buf[i+3]] = buf[buf[i+1]] + buf[buf[i+2]]
				case 2:
					buf[buf[i+3]] = buf[buf[i+1]] * buf[buf[i+2]]
				case 99:
					break
				}
			}
			if buf[0] == 19690720 {
				fmt.Println(100 * i + k)
				return
			}
		}
	}
}
