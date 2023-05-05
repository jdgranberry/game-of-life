package grid

type Coord struct {
	X int
	Y int
}

type Grid struct {
	Length    int
	liveCells map[Coord]bool
}

func NewGrid(seed map[Coord]bool, length int) *Grid {
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
	return g.liveCells[c]
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
	var nextLiveCoords = make(map[Coord]bool)

	for i := 0; i < g.Length; i++ {
		for j := 0; j < g.Length; j++ {
			if g.willLive(Coord{i, j}) {
				nextLiveCoords[Coord{i, j}] = true
			}
		}
	}

	g.liveCells = nextLiveCoords
}
