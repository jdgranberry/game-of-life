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

	g := grid.NewGrid(map[grid.Coord]bool{
		{1, 2}: true,
		{2, 3}: true,
		{3, 1}: true,
		{3, 2}: true,
		{3, 3}: true}, gridSize)

	for i := 0; i < ticks; i++ {
		fmt.Println(output.GetDisplayGrid(g))
		g.Tick()
		time.Sleep(sleepMs * time.Millisecond)
	}

}
