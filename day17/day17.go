package main

import (
	"fmt"
	"strings"

	"github.com/dotMaro/AoC2020/utils"
)

func main() {
	input := utils.ReadFile("day17/input.txt")
	conway := parseConway3D(input)
	conway.steps(6)
	utils.Print("Task 1. Total active cubes after 6 steps: %d", conway.totalActive())
	conway4D := parseConway4D(input)
	conway4D.steps(6)
	utils.Print("Task 2. Total active cubes after 6 steps: %d", conway4D.totalActive())
}

type conway3D struct {
	cubes [][][]bool // z y x
}

func parseConway3D(input string) conway3D {
	cubes := make([][][]bool, 1)
	lines := utils.SplitLine(input)
	cubes[0] = make([][]bool, len(lines))
	for row, line := range lines {
		cubes[0][row] = make([]bool, len(line))
		for col, c := range line {
			cubes[0][row][col] = c == '#'
		}
	}

	return conway3D{cubes: cubes}
}

func (c *conway3D) String() string {
	var b strings.Builder
	for _, layer := range c.cubes {
		for _, row := range layer {
			for _, cube := range row {
				if cube {
					b.WriteRune('#')
				} else {
					b.WriteRune('.')
				}
			}
			b.WriteRune('\n')
		}
		b.WriteRune('\n')
	}
	return b.String()
}

func (c *conway3D) step() {
	copy := c.expandedBlank()

	for z, layer := range copy {
		for y, row := range layer {
			for x := range row {
				copy[z][y][x] = c.shouldBeActive(z-1, y-1, x-1)
			}
		}
	}
	c.cubes = copy
}

func (c *conway3D) steps(n int) {
	for i := 0; i < n; i++ {
		c.step()
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// neighborCount returns the amount of neighbors the coordinate has.
// Supports taking numbers outside of index.
func (c *conway3D) neighborCount(z, y, x int) int {
	count := 0
	for checkZ := max(z-1, 0); checkZ <= min(z+1, len(c.cubes)-1); checkZ++ {
		for checkY := max(y-1, 0); checkY <= min(y+1, len(c.cubes[0])-1); checkY++ {
			for checkX := max(x-1, 0); checkX <= min(x+1, len(c.cubes[0][0])-1); checkX++ {
				// Do not include the cube as its own neighbor.
				if checkZ == z && checkY == y && checkX == x {
					continue
				}
				if c.cubes[checkZ][checkY][checkX] {
					count++
				}
			}
		}
	}
	return count
}

func (c *conway3D) shouldBeActive(z, y, x int) bool {
	currentActive := false
	if z >= 0 && z < len(c.cubes) && y >= 0 && y < len(c.cubes[0]) && x >= 0 && x < len(c.cubes[0][0]) {
		currentActive = c.cubes[z][y][x]
	}
	neighborCount := c.neighborCount(z, y, x)

	if currentActive {
		return neighborCount == 2 || neighborCount == 3
	}
	return neighborCount == 3
}

func (c *conway3D) totalActive() int {
	total := 0
	for _, layer := range c.cubes {
		for _, row := range layer {
			for _, cube := range row {
				if cube {
					total++
				}
			}
		}
	}
	return total
}

// expandedCopy expands the conway dimension by two in each direction and returns it as a copy.
func (c *conway3D) expandedBlank() [][][]bool {
	copy := make([][][]bool, len(c.cubes)+2)
	for z := 0; z < len(c.cubes)+2; z++ {
		copy[z] = make([][]bool, len(c.cubes[0])+2)
		for y := 0; y < len(c.cubes[0])+2; y++ {
			copy[z][y] = make([]bool, len(c.cubes[0][0])+2)
		}
	}
	return copy
}

type conway4D struct {
	cubes [][][][]bool // w z y x
}

func parseConway4D(input string) conway4D {
	cubes := make([][][][]bool, 1)
	lines := utils.SplitLine(input)
	cubes[0] = make([][][]bool, 1)
	cubes[0][0] = make([][]bool, len(lines))
	for row, line := range lines {
		cubes[0][0][row] = make([]bool, len(line))
		for col, c := range line {
			cubes[0][0][row][col] = c == '#'
		}
	}

	return conway4D{cubes: cubes}
}

func (c *conway4D) String() string {
	var b strings.Builder

	for w, wLayer := range c.cubes {
		for z, zLayer := range wLayer {
			b.WriteString(fmt.Sprintf("w=%d, z=%d\n", w, z))
			for _, row := range zLayer {
				for _, cube := range row {
					if cube {
						b.WriteRune('#')
					} else {
						b.WriteRune('.')
					}
				}
				b.WriteRune('\n')
			}
			b.WriteRune('\n')
		}
	}
	return b.String()
}

func (c *conway4D) step() {
	copy := c.expandedBlank()

	for w, wLayer := range copy {
		for z, zLayer := range wLayer {
			for y, row := range zLayer {
				for x := range row {
					copy[w][z][y][x] = c.shouldBeActive(w-1, z-1, y-1, x-1)
				}
			}
		}
	}
	c.cubes = copy
}

func (c *conway4D) steps(n int) {
	for i := 0; i < n; i++ {
		c.step()
	}
}

// neighborCount returns the amount of neighbors the cube at the specified coordinate has.
// Supports taking non-indexable numbers.
func (c *conway4D) neighborCount(w, z, y, x int) int {
	count := 0
	for checkW := max(w-1, 0); checkW <= min(w+1, len(c.cubes)-1); checkW++ {
		for checkZ := max(z-1, 0); checkZ <= min(z+1, len(c.cubes[0])-1); checkZ++ {
			for checkY := max(y-1, 0); checkY <= min(y+1, len(c.cubes[0][0])-1); checkY++ {
				for checkX := max(x-1, 0); checkX <= min(x+1, len(c.cubes[0][0][0])-1); checkX++ {
					// Do not include the cube as its own neighbor.
					if checkW == w && checkZ == z && checkY == y && checkX == x {
						continue
					}
					if c.cubes[checkW][checkZ][checkY][checkX] {
						count++
					}
				}
			}
		}
	}
	return count
}

func (c *conway4D) shouldBeActive(w, z, y, x int) bool {
	currentActive := false
	if w >= 0 && w < len(c.cubes) && z >= 0 && z < len(c.cubes[0]) && y >= 0 && y < len(c.cubes[0][0]) && x >= 0 && x < len(c.cubes[0][0][0]) {
		currentActive = c.cubes[w][z][y][x]
	}
	neighborCount := c.neighborCount(w, z, y, x)

	if currentActive {
		return neighborCount == 2 || neighborCount == 3
	}
	return neighborCount == 3
}

func (c *conway4D) totalActive() int {
	total := 0
	for _, wLayer := range c.cubes {
		for _, zLayer := range wLayer {
			for _, row := range zLayer {
				for _, cube := range row {
					if cube {
						total++
					}
				}
			}
		}
	}
	return total
}

// expandedCopy expands the conway dimension by two in each direction and returns it as a copy.
func (c *conway4D) expandedBlank() [][][][]bool {
	copy := make([][][][]bool, len(c.cubes)+2)
	for w := 0; w < len(c.cubes)+2; w++ {
		copy[w] = make([][][]bool, len(c.cubes[0])+2)
		for z := 0; z < len(c.cubes[0])+2; z++ {
			copy[w][z] = make([][]bool, len(c.cubes[0][0])+2)
			for y := 0; y < len(c.cubes[0][0])+2; y++ {
				copy[w][z][y] = make([]bool, len(c.cubes[0][0][0])+2)
			}
		}
	}
	return copy
}
