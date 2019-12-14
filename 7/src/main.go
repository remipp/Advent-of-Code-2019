package main

import (
	"IntCode"
	"os"
	"bufio"
	"strconv"
	"strings"
	"bytes"
	"fmt"
)

// Perm calls f with each permutation of a.
func Perm(a []int, f func([]int)) {
    perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []int, f func([]int), i int) {
    if i > len(a) {
        f(a)
        return
    }
    perm(a, f, i+1)
    for j := i + 1; j < len(a); j++ {
        a[i], a[j] = a[j], a[i]
        perm(a, f, i+1)
        a[i], a[j] = a[j], a[i]
    }
}

func startComputer(data []int, phase int, input int) int {
	a := IntCode.NewComputer(data)
	var aOut bytes.Buffer
	a.SetStdout(&aOut)
	a.SetStdin(strings.NewReader(strconv.Itoa(phase) + "\n" + strconv.Itoa(input) + "\n"))
	a.Run()
	n, _ := strconv.Atoi(aOut.String())
	return n
}

type result struct {
	phase []int
	output int
}

func main() {
	file, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(file)
	var data []int
	scanner.Scan()
	buf := strings.Split(scanner.Text(), ",")
	for _, x := range buf {
		n, _ := strconv.Atoi(x)
		data = append(data, n)
	}
	var best result
	Perm([]int{0, 1, 2, 3, 4}, func(p []int) {
		fmt.Println("Trying", p)
		var res int
		for _, x := range p {
			res = startComputer(append([]int{}, data...), x, res)
		}
		fmt.Println("Got", res)
		if res > best.output {
			best.output = res
			best.phase = p
		}
	})
	fmt.Println("Best combination", best)
}
