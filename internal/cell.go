package internal

import (
	"math/rand"
)

type Cell struct {
	pixels         [][]Pixel
	neighbors      [8]*Cell
	alive          bool
	wasAlive       bool
	timesInVirtual int
}

func (c *Cell) Init(pixels [][]Pixel, neighbors [8]*Cell) {
	c.alive = rand.Intn(100) < 10
	c.wasAlive = c.alive

	c.pixels = pixels
	c.neighbors = neighbors
}

func (c *Cell) Update(maxTimesInVirtual int) {
	if c.timesInVirtual >= maxTimesInVirtual {
		c.alive = false
		return
	}

	alive := 0
	for _, cell := range c.neighbors {
		if cell != nil && cell.wasAlive {
			alive++
		}
	}
	if c.alive {
		c.alive = alive == 2 || alive == 3
	} else {
		c.alive = alive == 3
	}
}

func (c *Cell) Draw() {
	for _, pixels := range c.pixels {
		for _, px := range pixels {
			if c.alive {
				px.White()
			} else {
				px.Black()
			}
		}
	}
	c.wasAlive = c.alive
}

func (c *Cell) InVirtual() {
	c.timesInVirtual += 1
}

func (c *Cell) OutVirtual() {
	c.timesInVirtual = 0
}
