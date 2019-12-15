package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

type layer struct {
	pixels [6*25]int
}

func (lay layer) countZeros() int {
	var count int
	for _, a := range lay.pixels {
		if a == 0 {
			count++
		}
	}
	return count
}

type picture struct {
	layers []layer
}

func (pic picture) fewestZeroPixels() layer {
	var buf layer
	var best int = 9999999
	for _, a := range pic.layers {
		if x := a.countZeros(); x < best {
			best = x
			buf = a
		}
	}
	return buf
}

func newPicture(scanner *bufio.Scanner) picture {
	var pic picture
	scanner.Split(bufio.ScanRunes)
	for {
		var buf layer
		for i := 0; i < len(buf.pixels); i++ {
			if !scanner.Scan() {
				return pic
			}
			buf.pixels[i], _ = strconv.Atoi(scanner.Text())
		}
		pic.layers = append(pic.layers, buf)
	}
}

func main(){
	if len(os.Args) < 2 {
		os.Exit(-1)
	}
	file, _ := os.Open(os.Args[1])
	defer file.Close()
	pic := newPicture(bufio.NewScanner(file))
	lay := pic.fewestZeroPixels()
	var ones int
	var twos int
	for _, a := range lay.pixels {
		if a == 1 {
			ones++
		} else if a == 2 {
			twos++
		}
	}
	fmt.Println(ones*twos)
}
