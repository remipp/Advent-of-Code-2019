package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
)

var pattern []int = []int{0, 1, 0, -1}

func Abs(x int) int {
        if x < 0 {
                return -x
        }
        return x
}

func fft(input []int, phases int) []int {
	//fmt.Println(phases, ":", input)
	if phases == 100 {
		// done
		return input
	}
	// amount of time to repeat element in pattern = phases+1
	var output []int
	for i := 0; i < len(input); i++ {
		var buf []int
		// p starts at 0
		idx := 0
		// "Skip" the first one
		repeated := 1
		for _, x := range input {
			if repeated == i+1 {
				repeated = 0
				idx++
			}
			if idx == len(pattern) {
				idx = 0
			}
			p := pattern[idx]
			res := x * p
			//fmt.Printf("%d * %d\n", x, p)
			buf = append(buf, res)
			repeated++
		}
		// Calculate sum of buf
		var sum int
		for _, x := range buf {
			sum += x
		}
		sum = Abs(sum) % 10
		//fmt.Println("Sum:", sum)
		output = append(output, sum)
	}

	return fft(output, phases+1)
}

// 0, 1, 0, -1
func main(){
	if len(os.Args) < 2 {
		os.Exit(-1)
	}
	file, _ := os.Open(os.Args[1])
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanRunes)

	var input []int

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err == nil {
			input = append(input, i)
		}
	}
	fmt.Println(fft(input, 0)[:8])
}
