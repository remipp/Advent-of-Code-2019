package main

import (
	"IntCode"
	"os"
	"bufio"
	"strconv"
	"strings"
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

func startComputer(data []int, input <-chan int, output chan<- int, haltChan chan<- bool) {
	a := IntCode.NewComputer(data)
	a.SetStdout(output)
	//a.SetStdin(strings.NewReader(strconv.Itoa(phase) + "\n" + strconv.Itoa(input) + "\n"))
	a.SetStdin(input)
	a.OnHalt = func() {
		haltChan <- true
	}
	go a.Run()
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
	fmt.Println(data)
	var best int
	Perm([]int{5, 6, 7, 8, 9}, func(p []int) {
		ea := make(chan int, 1)
		ab := make(chan int)
		bc := make(chan int)
		cd := make(chan int)
		de := make(chan int)

		haltChan := make(chan bool)
		startComputer(append([]int{}, data...), ea, ab, haltChan)
		startComputer(append([]int{}, data...), ab, bc, haltChan)
		startComputer(append([]int{}, data...), bc, cd, haltChan)
		startComputer(append([]int{}, data...), cd, de, haltChan)
		startComputer(append([]int{}, data...), de, ea, haltChan)

		ea <- p[0]
		ab <- p[1]
		bc <- p[2]
		cd <- p[3]
		de <- p[4]

		ea <- 0

		for i := 0; i < 5; i++ {
			<-haltChan
		}

		val := <-ea
		if val > best {
			best = val
		}
	})
	fmt.Println("BEST", best)
}
