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

func equalsX(mns, prev []moon) bool {
	for i := 0; i < len(mns); i++ {
		if mns[i].position.x != prev[i].position.x || mns[i].velocity.x != prev[i].velocity.x {
			return false
		}
	}
	return true
}

func equalsY(mns, prev []moon) bool {
	for i := 0; i < len(mns); i++ {
		if mns[i].position.y != prev[i].position.y || mns[i].velocity.y != prev[i].velocity.y {
			return false
		}
	}
	return true
}

func equalsZ(mns, prev []moon) bool {
	for i := 0; i < len(mns); i++ {
		if mns[i].position.z != prev[i].position.z || mns[i].velocity.z != prev[i].velocity.z {
			return false
		}
	}
	return true
}

func GCD(a, b int) int {
      for b != 0 {
              t := b
              b = a % b
              a = t
      }
      return a
}

func LCM(a, b int, integers ...int) int {
      result := a * b / GCD(a, b)

      for i := 0; i < len(integers); i++ {
              result = LCM(result, integers[i])
      }

      return result
}

func main() {
	if len(os.Args) < 2 {
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
	initialState := append([]moon{}, moons...)
	var ctr int
	x_cycle, y_cycle, z_cycle := -1, -1, -1
	for {
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

		ctr++
		if x_cycle == -1 && equalsX(moons, initialState) {
			fmt.Println("FOUND X in tick", ctr)
			x_cycle = ctr
		}
		if y_cycle == -1 && equalsY(moons, initialState) {
			fmt.Println("FOUND Y in tick", ctr)
			y_cycle = ctr
		}
		if z_cycle == -1 && equalsZ(moons, initialState) {
			fmt.Println("FOUND Z in tick", ctr)
			z_cycle = ctr
		}
		if x_cycle > 0 && y_cycle > 0 && z_cycle > 0 {
			break
		}
	}
	fmt.Println(LCM(x_cycle, y_cycle, z_cycle))
}
