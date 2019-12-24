package main

import (
	"fmt"
	"os"
	"bufio"
	"regexp"
	"strconv"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type coordinates struct {
	x, y, z int
}

type moon struct {
	position coordinates
	velocity coordinates
}

func (m *moon) updateVelocity(m1 moon) {
	if m1.position.x > m.position.x {
		m.velocity.x++
	} else if m1.position.x < m.position.x {
		m.velocity.x--
	}
	if m1.position.y > m.position.y {
		m.velocity.y++
	} else if m1.position.y < m.position.y {
		m.velocity.y--
	}
	if m1.position.z > m.position.z {
		m.velocity.z++
	} else if m1.position.z < m.position.z {
		m.velocity.z--
	}
}

func (m *moon) applyVelocity() {
	m.position.x += m.velocity.x
	m.position.y += m.velocity.y
	m.position.z += m.velocity.z
}

func (m moon) potEnergy() int {
	return Abs(m.position.x) + Abs(m.position.y) + Abs(m.position.z)
}

func (m moon) kinEnergy() int {
	return Abs(m.velocity.x) + Abs(m.velocity.y) + Abs(m.velocity.z)
}

func (m moon) totEnergy() int {
	return m.potEnergy() * m.kinEnergy()
}

func main() {
	if len(os.Args) < 3 {
		os.Exit(-1)
	}
	file, _ := os.Open(os.Args[1])
	defer file.Close()
	scanner := bufio.NewScanner(file)
	re, _ := regexp.Compile("-?[0-9]+")
	var moons []moon
	for (scanner.Scan()) {
		var buf []int
		for _, x := range re.FindAllString(scanner.Text(), -1) {
			n, _ := strconv.Atoi(x)
			buf = append(buf, n)
		}
		moons = append(moons, moon{coordinates{buf[0], buf[1], buf[2]}, coordinates{}})
	}
	var energy int
	steps, _ := strconv.Atoi(os.Args[2])
	for x := 0; x < steps; x++ {
		for i := 0; i < len(moons); i++ {
			for j := 0; j < len(moons); j++ {
				// Moon doesn't change it's own velocity
				if i == j {
					continue
				}
				moons[i].updateVelocity(moons[j])
			}
		}
		for i := 0; i < len(moons); i++ {
			moons[i].applyVelocity()
		}
	}
	for _, moon := range moons {
		energy += moon.totEnergy()
	}
	fmt.Println(moons)
	fmt.Println(energy)
}
