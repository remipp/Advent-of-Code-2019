package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"image"
	"image/png"
	"image/color"
)

const width = 25
const height = 6

type layer struct {
	pixels [height*width]int
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

func (pic picture) render() layer {
	var buf layer
	for i := 0; i < len(buf.pixels); i++ {
		for _, a := range pic.layers {
			if a.pixels[i] != 2 || i == 24 {
				buf.pixels[i] = a.pixels[i]
				break
			}
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
	if len(os.Args) < 3 {
		os.Exit(-1)
	}
	file, _ := os.Open(os.Args[1])
	defer file.Close()
	pic := newPicture(bufio.NewScanner(file))
	r := pic.render()
	fmt.Println(r)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			var c color.Color
			if r.pixels[y*width+x] == 0 {
				c = color.Gray{255}
			} else {
				c = color.Gray{0}
			}
			img.Set(x, y, c)
		}
	}
	out, _ := os.Create(os.Args[2])
	png.Encode(out, img)
}
