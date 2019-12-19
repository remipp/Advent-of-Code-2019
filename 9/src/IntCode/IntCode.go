package IntCode

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type instruction struct {
	size        int
	writeParams []int
	f           func(params ...int)
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
	memory             *[]int
	instructions       map[int]instruction
	instructionPointer *int
	relativeBase       *int
}

func (c *Computer) writeMemory(address int, data int) {
	if address >= len(*c.memory) {
		newBuffer := make([]int, address+1)
		copy(newBuffer, *c.memory)
		*c.memory = newBuffer
	}
	(*c.memory)[address] = data
}
func (c *Computer) readMemory(address int) int {
	if address >= len(*c.memory) {
		newBuffer := make([]int, address+1)
		copy(newBuffer, *c.memory)
		*c.memory = newBuffer
	}
	return (*c.memory)[address]
}

func (c *Computer) readMemorySlice(start int, end int) []int {
	if end >= len(*c.memory) {
		newBuffer := make([]int, end+1)
		copy(newBuffer, *c.memory)
		*c.memory = newBuffer
	}
	return (*c.memory)[start:end]
}

func (c Computer) Run() {
	// Parse instruction
	// Instruction number: %100
	for {
		i := c.readMemory(*c.instructionPointer) % 100
		ins := c.instructions[i]
		// Handle parameter modes
		modes := c.readMemory(*c.instructionPointer) / 100
		params := append([]int{}, c.readMemorySlice(*c.instructionPointer+1, *c.instructionPointer+ins.size)...)
		fmt.Println(*c.instructionPointer, "Instruction:", i, "Base:", *c.relativeBase, "Params:", params, "Modes:", modes)
		for x := 0; x < len(params); x++ {
			m := modes % 10
			modes /= 10
			writeParam := ins.isWriteParam(x)
			switch m {
			case 0: // Position Mode
				if writeParam {
					break
				}
				params[x] = c.readMemory(params[x])
			case 1: // Immediate Mode
				continue
			case 2: // Relative Mode
				if writeParam {
					params[x] = *c.relativeBase + params[x]
				} else {
					params[x] = c.readMemory(*c.relativeBase + params[x])
				}
			default:
				log.Fatalf("Unexpected paramter mode %v", m)
			}
		}
		fmt.Println("With applied modes:", params)
		ins.f(params...)
		*c.instructionPointer += ins.size
	}
}

func NewComputer(memory []int) Computer {
	var c Computer
	c.memory = &memory
	c.instructionPointer = new(int)
	c.relativeBase = new(int)
	c.instructions = make(map[int]instruction)
	c.instructions = map[int]instruction{
		1: {4, []int{2}, func(params ...int) {
			c.writeMemory(params[2], params[0]+params[1])
		}},
		2: {4, []int{2}, func(params ...int) {
			c.writeMemory(params[2], params[0]*params[1])
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
			c.writeMemory(params[0], n)
		}},
		4: {2, []int{}, func(params ...int) {
			fmt.Println(params[0])
		}},
		5: {3, []int{}, func(params ...int) {
			if params[0] != 0 {
				*c.instructionPointer = params[1] - 3
				fmt.Println("Set IP to", *c.instructionPointer+3)
			}
		}},
		6: {3, []int{}, func(params ...int) {
			if params[0] == 0 {
				*c.instructionPointer = params[1] - 3
			}
		}},
		7: {4, []int{2}, func(params ...int) {
			if params[0] < params[1] {
				c.writeMemory(params[2], 1)
			} else {
				c.writeMemory(params[2], 0)
			}
		}},
		8: {4, []int{2}, func(params ...int) {
			if params[0] == params[1] {
				c.writeMemory(params[2], 1)
			} else {
				c.writeMemory(params[2], 0)
			}
		}},
		9: {2, []int{}, func(params ...int) {
			*c.relativeBase += params[0]
		}},
		99: {1, []int{}, func(params ...int) {
			// Stop computer somehow?
			fmt.Println("HALT")
			os.Exit(0)
		}},
	}
	return c
}
