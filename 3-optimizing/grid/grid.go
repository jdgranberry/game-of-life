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

func (g *Grid) IsAlive(c Coord) bool {
	return g.liveCells[c]
}

func (g *Grid) liveNeighbors(c Coord) int {
	var result int
	neighbors := g.allNeighbors(c)

	for _, n := range neighbors {
		if g.IsAlive(n) {
			result++
		}
	}

	return result
}

func (g *Grid) allNeighbors(c Coord) []Coord {
	var result []Coord

	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			newX := c.X + i
			newY := c.Y + j
			if !(newX < 0 || newX >= g.Length || newY < 0 || newY >= g.Length || (i == 0 && j == 0)) {
				result = append(result, Coord{newX, newY})
			}
		}
	}

	return result
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
	coordsToCheck := make(map[Coord]bool)
	for k, _ := range g.liveCells {
		coordsToCheck[k] = true
		neighbors := g.allNeighbors(k)
		for _, n := range neighbors {
			coordsToCheck[n] = true
		}
	}

	// TODO in a sparse map, we can optimize this by only checking cells neighboring live cells
	for k, _ := range coordsToCheck {
		if g.willLive(k) {
			nextLiveCoords[k] = true
		}
	}

	g.liveCells = nextLiveCoords
}
