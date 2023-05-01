package grid

import (
	"errors"
	"fmt"
	"log"
)

type Grid struct {
	Length        uint16
	grid          [][]bool
	markedForLife []Coord
}

type Coord struct {
	X uint16
	Y uint16
}

// NewGrid creates a new len X len grid of false bools
func NewGrid(len uint16) *Grid {
	newGrid := make([][]bool, len)
	for i := range newGrid {
		newGrid[i] = make([]bool, len)
	}

	return &Grid{
		Length:        len,
		grid:          newGrid,
		markedForLife: make([]Coord, 0),
	}
}

// Seed seeds the board with live cells
func (g *Grid) Seed(cells []Coord) error {
	for _, c := range cells {
		err := c.validateBounds(g.Length)
		if err != nil {
			return err
		}
	}

	for _, c := range cells {
		g.grid[c.X][c.Y] = true
	}

	return nil
}

// Tick moves the board forward one step
func (g *Grid) Tick() {
	// look at grid and mark live cells for next generation
	for i, _ := range g.grid {
		for j, _ := range g.grid {
			var c = Coord{uint16(i), uint16(j)}
			n, err := g.liveNeighbors(c)
			if err != nil {
				log.Fatalf("Received invalid coordinates, halting. %v", err)
			}
			if i == 2 && j == 3 {
				fmt.Printf("live neighbors for %d, %d: %d\n", i, j, n)
			}
			g.checkForLife(c, n)
		}
	}

	// mark all cells dead, then mark cells live if they are in markedForLife
	for i, _ := range g.grid {
		for j, _ := range g.grid {
			g.grid[i][j] = false
		}
	}
	for _, c := range g.markedForLife {
		g.grid[c.X][c.Y] = true
	}

	// Reset marks
	g.markedForLife = make([]Coord, 0)
}

// GetGrid returns a copy of the grid
func (g *Grid) GetGrid() [][]bool {
	duplicate := make([][]bool, len(g.grid))
	for i := range g.grid {
		duplicate[i] = make([]bool, len(g.grid[i]))
		copy(duplicate[i], g.grid[i])
	}
	return duplicate
}

// CheckForLife checks if a cell should be alive next turn and adds it to markedForLife
func (g *Grid) checkForLife(coord Coord, numLiveNeighbors int) {
	// if cell is live
	if g.grid[coord.X][coord.Y] && numLiveNeighbors >= 2 && numLiveNeighbors <= 3 {
		g.markedForLife = append(g.markedForLife, coord)
	}

	// if cell is deceased
	if !g.grid[coord.X][coord.Y] && numLiveNeighbors == 3 {
		g.markedForLife = append(g.markedForLife, coord)
	}

}

// liveNeighbors returns the number of living cells around a coordinate
func (g *Grid) liveNeighbors(coord Coord) (int, error) {
	var liveNeighbors int
	if err := coord.validateBounds(g.Length); err != nil {
		return 0, err
	}

	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			// if coord value is goes negative, uint16 will wrap around to 65535. safe to assume grids will not be this large
			// so validateBounds will always return false. we could enforce a grid size or change to a signed int and check
			// for negative but meh
			tempX := uint16(int(coord.X) + i)
			tempY := uint16(int(coord.Y) + j)
			tempCoord := Coord{tempX, tempY}

			if coord != tempCoord && g.isCoordLive(tempCoord) {
				liveNeighbors++
			}
		}
	}

	return liveNeighbors, nil
}

// validateBounds validates Coord's X and Y values are within length
func (c *Coord) validateBounds(length uint16) error {
	if c.X >= length || c.Y >= length {
		return errors.New(fmt.Sprintf("Coordinates out of bound for grid of length %d. Coord: %d, %d", length, c.X, c.Y))
	}

	return nil
}

// isCoordLive returns true if the neighboring coord is within the grid and the value is true
func (g *Grid) isCoordLive(t Coord) bool {
	if err := t.validateBounds(g.Length); err != nil {
		return false
	}

	return g.grid[t.X][t.Y] == true
}
