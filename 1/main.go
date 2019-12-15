package main
import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)


func main() {
	if len(os.Args) < 2 {
		os.Exit(-1)
	}
	file, _ := os.Open(os.Args[1])
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var res int
	for scanner.Scan() {
		x, _ := strconv.Atoi(scanner.Text())
		res += int(x/3)-2
	}
	fmt.Println(res)
}
