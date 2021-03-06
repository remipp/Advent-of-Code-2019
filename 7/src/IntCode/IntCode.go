package IntCode

import (
	"fmt"
	"log"
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
	stdin *<-chan int
	stdout *chan<- int
	halted *bool
	OnHalt func()
}

func (c Computer) SetStdout(w chan<- int) {
	*c.stdout = w
}

func (c Computer) SetStdin(r <-chan int) {
	*c.stdin = r
}

func (c Computer) Run() {
// Parse instruction
	// Instruction number: %100
	for !*c.halted {
		i := c.memory[*c.instructionPointer] % 100
		ins := c.instructions[i]
		// Handle parameter modes
		modes := c.memory[*c.instructionPointer] / 100
		params := append([]int{}, c.memory[*c.instructionPointer+1:*c.instructionPointer+ins.size]...)
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
				log.Fatalf("Unexpected paramter mode %v", m)
			}
		}
		fmt.Println("With applied modes:", params)
		ins.f(params...)
		*c.instructionPointer += ins.size
	}
	c.OnHalt()
}

func NewComputer(memory []int) Computer {
	var c Computer
	c.memory = memory
	c.instructionPointer = new(int)
	c.halted = new(bool)
	c.stdin = new(<-chan int)
	c.stdout = new(chan<- int)
	c.instructions = make(map[int]instruction)
	c.instructions = map[int]instruction{
		1: {4, []int{2}, func(params ...int) {
			c.memory[params[2]] = params[0] + params[1]
		}},
		2: {4, []int{2}, func(params ...int) {
			c.memory[params[2]] = params[0] * params[1]
		}},
		3: {2, []int{0}, func(params ...int) {
			n := <-(*c.stdin)
			c.memory[params[0]] = n
		}},
		4: {2, []int{}, func(params ...int) {
			(*c.stdout) <- params[0]
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
			fmt.Println("HALT")
			*c.halted = true
		}},
	}
	return c
}
