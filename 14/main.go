package main

import (
	"os"
	"fmt"
	"bufio"
	"regexp"
	"strconv"
	"math"
)

type chemical struct {
	quantity int
	name string
}

type reaction struct {
	result chemical
	ingredients []chemical
}

var oreNeeded int
var reactions map[string]reaction
var leftOver map[string]int

func doTheThing(s string, numNeeded int) {
	if numNeeded == 0 {
		return
	}
	if s == "ORE" {
		oreNeeded += numNeeded
		return
	}

	timesToMake := int(math.Ceil(float64(numNeeded)/float64(reactions[s].result.quantity)))
	leftOver[s] += timesToMake*reactions[s].result.quantity-numNeeded

	for leftOver[s] >= reactions[s].result.quantity {
		timesToMake--
		leftOver[s] -= reactions[s].result.quantity
	}

	for _, ingredient := range reactions[s].ingredients {
		doTheThing(ingredient.name, ingredient.quantity*timesToMake)
	}
}

func main(){
	if len(os.Args) < 2 {
		os.Exit(-1)
	}
	file, _ := os.Open(os.Args[1])
	defer file.Close()
	scanner := bufio.NewScanner(file)
	re, _ := regexp.Compile("[0-9]+ [^ ,]*")
	quantity, _ := regexp.Compile("[0-9]+")
	name, _ := regexp.Compile("[A-z]+")

	reactions = make(map[string]reaction)
	leftOver = make(map[string]int)
	for scanner.Scan() {
		buf := re.FindAllString(scanner.Text(), -1)
		var res chemical
		res.name = name.FindString(buf[len(buf)-1])
		res.quantity, _ = strconv.Atoi(quantity.FindString(buf[len(buf)-1]))
		for i := 0; i < len(buf)-1; i++ {
			var ingredient chemical
			ingredient.name = name.FindString(buf[i])
			ingredient.quantity, _ = strconv.Atoi(quantity.FindString(buf[i]))
			reactions[res.name] = reaction{res, append(reactions[res.name].ingredients, ingredient)}
		}
	}
	doTheThing("FUEL", 1)
	fmt.Println(oreNeeded)
}
