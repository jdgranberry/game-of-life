package output

import grid "game-of-life-golang/grid"

const (
	deadCell = "◦"
	liveCell = "▓"
)

func GetDisplayGrid(g *grid.Grid) string {
	s := ""

	for i := 0; i < g.Length; i++ {
		for j := 0; j < g.Length; j++ {
			c := grid.Coord{X: i, Y: j}
			if g.IsAlive(c) {
				s += liveCell
			} else {
				s += deadCell
			}
		}
		s += "\n"
	}
	return s
}
