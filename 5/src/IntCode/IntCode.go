package IntCode

import (
	"os"
	"bufio"
	"fmt"
	"strconv"
)

type instruction struct {
	size int
	writeParams []int
	f func(params ...int)
}

func (ins instruction) isWriteParam(i int) bool {
	for _, a := range ins.writeParams {
		if i == a {
			return true
		}
	}
	return false
}

type Computer struct {
	memory []int
	instructions map[int]instruction
	instructionPointer *int
}

func (c Computer) Run() {
	// Parse instruction
	// Instruction number: %100
	for {
		i := c.memory[*c.instructionPointer] % 100
		ins := c.instructions[i]
		// Handle parameter modes
		modes := c.memory[*c.instructionPointer] / 100
		params := c.memory[*c.instructionPointer+1:*c.instructionPointer+ins.size]
		fmt.Println(*c.instructionPointer, "Instruction:", i, "Params:", params, "Modes:", modes)
		for x := 0; x < len(params); x++ {
			m := modes % 10
			modes /= 10
			// Pass write params in immediate mode
			if ins.isWriteParam(x) {
				continue
			}
			switch(m) {
			case 0: // Position Mode
				params[x] = c.memory[params[x]]
			case 1: // Immediate Mode
				continue
			default:
				fmt.Errorf("Unexpected parameter mode: %v", m)
			}
		}
		fmt.Println("With applied modes:", params)
		ins.f(params...)
		*c.instructionPointer += ins.size
	}
}

func NewComputer(memory []int) Computer {
	var c Computer
	c.memory = memory
	c.instructionPointer = new(int)
	c.instructions = make(map[int]instruction)
	c.instructions = map[int]instruction{
		1: {4, []int{2}, func(params ...int) {
			c.memory[params[2]] = params[0] + params[1]
		}},
		2: {4, []int{2}, func(params ...int) {
			c.memory[params[2]] = params[0] * params[1]
		}},
		3: {2, []int{0}, func(params ...int) {
			r := bufio.NewReader(os.Stdin)
			s, err := r.ReadString('\n')
			if err != nil {
				fmt.Errorf("Couldn't read from Stdin: %v\n", err)
			}
			s = s[:len(s)-1]
			n, err := strconv.Atoi(s)
			if err != nil {
				fmt.Errorf("Couldn't convert to int: %v\n", err)
			}
			c.memory[params[0]] = n
		}},
		4: {2, []int{}, func(params ...int) {
			fmt.Println(params[0])
		}},
		5: {3, []int{}, func(params ...int) {
			if params[0] != 0 {
				*c.instructionPointer = params[1]-3
				fmt.Println("Set IP to", *c.instructionPointer+3)
			}
		}},
		6: {3, []int{}, func(params ...int) {
			if params[0] == 0 {
				*c.instructionPointer = params[1]-3
			}
		}},
		7: {4, []int{2}, func(params ...int) {
			if params[0] < params[1] {
				c.memory[params[2]] = 1
			} else {
				c.memory[params[2]] = 0
			}
		}},
		8: {4, []int{2}, func(params ...int) {
			if params[0] == params[1] {
				c.memory[params[2]] = 1
			} else {
				c.memory[params[2]] = 0
			}
		}},
		99: {1, []int{}, func(params ...int) {
			// Stop computer somehow?
			fmt.Println("HALT")
			os.Exit(0)
		}},
	}
	return c
}
