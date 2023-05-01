package main

import (
	"game-of-life-golang/grid"
	o "game-of-life-golang/output"
	"log"
	"time"
)

const (
	// Length/Width of square grid
	gridSize = 75
	ticks    = 500
	sleepMs  = 200
)

func main() {
	g := grid.NewGrid(gridSize)
	// glider pattern
	seed := []grid.Coord{{0, 1}, {1, 2}, {2, 0}, {2, 1}, {2, 2}}
	err := g.Seed(seed)
	if err != nil {
		log.Fatalf("error seeding grid: %v", err)
	}

	d := o.NewDisplay(*g)
	d.Display()

	for i := 0; i < ticks; i++ {
		d.Grid.Tick()
		d.Display()
		time.Sleep(sleepMs * time.Millisecond)
	}
}
