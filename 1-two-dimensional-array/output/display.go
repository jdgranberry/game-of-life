package output

import (
	"fmt"
	g "game-of-life-golang/grid"
)

const liveChar rune = '▓'
const deadChar rune = '.'

type Display struct {
	Grid    g.Grid
	display [][]rune
}

func NewDisplay(grid g.Grid) Display {
	return Display{
		Grid:    grid,
		display: nil,
	}
}

func (d *Display) Display() {
	d.readyDisplay()
	d.printDisplay()
}

func (d *Display) readyDisplay() {
	grid := d.Grid.GetGrid()
	d.display = make([][]rune, d.Grid.Length)
	for i, _ := range grid {
		line := make([]rune, d.Grid.Length)
		for j, _ := range grid {
			if grid[i][j] {
				line[j] = liveChar
			} else {
				line[j] = deadChar
			}
		}
		d.display[i] = line
	}
}

func (d *Display) printDisplay() {
	// build single string beginning with \r and using \n between lines
	//fmt.Println("\033[0;0H")
	out := ""
	for i, _ := range d.display {
		out += fmt.Sprintf("%6d %s\n", i, string(d.display[i]))
	}
	fmt.Print(out)
}
