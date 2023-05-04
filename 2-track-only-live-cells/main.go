package main

import (
	"fmt"
	"game-of-life-golang/grid"
	"game-of-life-golang/output"
	"time"
)

const (
	gridSize = 75
	ticks    = 500
	sleepMs  = 200
)

func main() {
	fmt.Println("hello world!")

	g := grid.NewGrid([]grid.Coord{{1, 2}, {2, 3}, {3, 1}, {3, 2}, {3, 3}}, gridSize)

	for i := 0; i < ticks; i++ {
		fmt.Println(output.GetDisplayGrid(g))
		g.Tick()
		time.Sleep(sleepMs * time.Millisecond)
	}

}
