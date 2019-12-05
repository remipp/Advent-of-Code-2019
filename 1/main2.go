package main
import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var res int
	for scanner.Scan() {
		x, _ := strconv.Atoi(scanner.Text())
		//res += int(x/3)-2
		var buf int
		for buf = x/3-2; buf > 0; buf = buf/3-2 {
			res += buf
		}
	}
	fmt.Println(res)
}
