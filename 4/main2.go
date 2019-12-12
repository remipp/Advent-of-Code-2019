package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isValid(num int) bool {
	a := strconv.Itoa(num)
	counts := make(map[byte]int)
	counts[a[0]]++
	for i := 1; i < len(a); i++ {
		if a[i] < a[i-1] {
			return false
		}
		counts[a[i]]++
	}
	for _, c := range counts {
		if c == 2 {
			return true
		}
	}
	return false
}

func main() {
	input := strings.Split(os.Args[1], "-")
	start, _ := strconv.Atoi(input[0])
	end, _ := strconv.Atoi(input[1])
	var ctr int
	for i := start; i <= end; i++ {
		if (isValid(i)) {
			ctr++
		}
	}
	fmt.Println(ctr)
}
