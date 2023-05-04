package grid

type Coord struct {
	X int
	Y int
}

type Grid struct {
	Length    int
	liveCells []Coord
}

func NewGrid(seed []Coord, length int) *Grid {
	return &Grid{Length: length, liveCells: seed}
}

func (g *Grid) liveNeighbors(c Coord) int {
	var result int

	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if g.IsAlive(Coord{c.X + i, c.Y + j}) {
				if i == 0 && j == 0 {
					continue
				}
				result++
			}
		}
	}

	return result
}

func (g *Grid) IsAlive(c Coord) bool {
	for _, l := range g.liveCells {
		if l == c {
			return true
		}
	}
	return false
}

func (g *Grid) willLive(c Coord) bool {
	neighbors := g.liveNeighbors(c)
	if (neighbors == 2 || neighbors == 3) && g.IsAlive(c) {
		return true
	}
	if neighbors == 3 && !g.IsAlive(c) {
		return true
	}
	return false
}

func (g *Grid) Tick() {
	var nextLiveCoords []Coord

	for i := 0; i < g.Length; i++ {
		for j := 0; j < g.Length; j++ {
			if g.willLive(Coord{i, j}) {
				nextLiveCoords = append(nextLiveCoords, Coord{i, j})
			}
		}
	}

	g.liveCells = nextLiveCoords
}
