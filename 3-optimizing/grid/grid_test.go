package grid

import (
	"reflect"
	"testing"
)

func TestGrid_Tick(t *testing.T) {
	glider := []Coord{{1, 2}, {2, 3}, {3, 1}, {3, 2}, {3, 3}}
	gliderNext := []Coord{{2, 1}, {2, 3}, {3, 2}, {3, 3}, {4, 2}}
	cross := []Coord{{0, 0}, {0, 2}, {1, 1}, {2, 0}, {2, 2}}
	crossNext := []Coord{{0, 1}, {1, 0}, {1, 2}, {2, 1}}
	blinker := []Coord{{0, 1}, {1, 1}, {2, 1}}
	blinkerNext := []Coord{{1, 0}, {1, 1}, {1, 2}}

	type fields struct {
		liveCells []Coord
	}
	tests := []struct {
		name   string
		fields fields
		want   []Coord
	}{
		{
			name: "one live cell, no live cells next turn",
			fields: fields{
				liveCells: []Coord{{0, 0}},
			},
			want: []Coord{},
		},
		{
			name: "two live cells, no live cells next turn",
			fields: fields{
				liveCells: []Coord{{0, 0}, {0, 1}},
			},
			want: []Coord{},
		},
		{
			name: "three live cells, two cell next turn",
			fields: fields{
				liveCells: []Coord{{0, 0}, {0, 1}, {0, 2}},
			},
			want: []Coord{{0, 1}, {1, 1}},
		},
		{
			name: "three live cells, four live cell next turn",
			fields: fields{
				liveCells: []Coord{{0, 0}, {0, 1}, {1, 0}},
			},
			want: []Coord{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
		},
		{
			name: "glider pattern, next phase of glider pattern",
			fields: fields{
				liveCells: glider,
			},
			want: gliderNext,
		},
		{
			name: "cross pattern, next phase of cross pattern",
			fields: fields{
				liveCells: cross,
			},
			want: crossNext,
		},
		{
			name: "cross next pattern doesn't change",
			fields: fields{
				liveCells: crossNext,
			},
			want: crossNext,
		},
		{
			name: "blinker -> blinkerNext",
			fields: fields{
				liveCells: blinker,
			},
			want: blinkerNext,
		},
		{
			name: "blinkerNext -> blinker",
			fields: fields{
				liveCells: blinkerNext,
			},
			want: blinker,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Grid{
				Length:    5,
				liveCells: sliceToMap(tt.fields.liveCells),
			}
			g.Tick()
			result := g.liveCells
			if !reflect.DeepEqual(result, sliceToMap(tt.want)) {
				if len(result) == 0 && len(tt.want) == 0 {
					return
				}
				t.Errorf("Tick() = %v, want %v", result, tt.want)
			}
		})
	}
}

func TestGrid_isAlive(t *testing.T) {
	type fields struct {
		length    int
		liveCells []Coord
	}
	type args struct {
		c Coord
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "live cell",
			fields: fields{
				length:    2,
				liveCells: []Coord{{0, 1}},
			},
			args: args{
				c: Coord{0, 1},
			},
			want: true,
		},
		{
			name: "dead cell",
			fields: fields{
				length:    2,
				liveCells: []Coord{{0, 1}, {1, 1}, {100, 100}},
			},
			args: args{
				c: Coord{50, 50},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Grid{
				Length:    tt.fields.length,
				liveCells: sliceToMap(tt.fields.liveCells),
			}
			if got := g.IsAlive(tt.args.c); got != tt.want {
				t.Errorf("IsAlive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrid_liveNeighbors(t *testing.T) {
	type fields struct {
		liveCells []Coord
	}
	type args struct {
		c Coord
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "no live neighbors",
			fields: fields{
				liveCells: []Coord{{0, 1}},
			},
			args: args{
				c: Coord{3, 3},
			},
			want: 0,
		},
		{
			name: "8 live neighbors",
			fields: fields{
				liveCells: []Coord{
					{0, 1}, {0, 2}, {0, 3},
					{1, 1}, {1, 2}, {1, 3},
					{2, 1}, {2, 2}, {2, 3},
				},
			},
			args: args{
				c: Coord{1, 2},
			},
			want: 8,
		},
		{
			name: "2 live neighbors",
			fields: fields{
				liveCells: []Coord{
					{0, 1}, {0, 2}, {0, 3},
					{1, 1}, {1, 2}, {1, 3},
					{2, 1}, {2, 2}, {2, 3},
				},
			},
			args: args{
				c: Coord{3, 1},
			},
			want: 2,
		},
		{
			name: "3 live neighbors",
			fields: fields{
				liveCells: []Coord{
					{0, 1}, {0, 2}, {0, 3},
					{1, 1}, {1, 2}, {1, 3},
					{2, 1}, {2, 2}, {2, 3},
				},
			},
			args: args{
				c: Coord{3, 2},
			},
			want: 3,
		},
		{
			name: "1 live neighbor",
			fields: fields{
				liveCells: []Coord{
					{0, 1}, {0, 2}, {0, 3},
					{1, 1}, {1, 2}, {1, 3},
					{2, 1}, {2, 2}, {2, 3},
				},
			},
			args: args{
				c: Coord{3, 4},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Grid{
				Length:    5,
				liveCells: sliceToMap(tt.fields.liveCells),
			}
			if got := g.liveNeighbors(tt.args.c); got != tt.want {
				t.Errorf("liveNeighbors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrid_willLive(t *testing.T) {
	type fields struct {
		liveCells []Coord
	}
	type args struct {
		c Coord
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "cell is live, no live neighbors",
			fields: fields{
				liveCells: []Coord{{0, 1}},
			},
			args: args{
				c: Coord{0, 1},
			},
			want: false,
		},
		{
			name: "cell is dead, no live neighbors",
			fields: fields{
				liveCells: []Coord{{0, 1}},
			},
			args: args{
				c: Coord{10, 10},
			},
			want: false,
		},
		{
			name: "cell is live, 8 live neighbors",
			fields: fields{
				liveCells: []Coord{
					{0, 1}, {0, 2}, {0, 3},
					{1, 1}, {1, 2}, {1, 3},
					{2, 1}, {2, 2}, {2, 3},
				},
			},
			args: args{
				c: Coord{1, 2},
			},
			want: false,
		},
		{
			name: "cell is dead, 8 live neighbors",
			fields: fields{
				liveCells: []Coord{
					{0, 1}, {0, 2}, {0, 3},
					{1, 1}, {1, 3},
					{2, 1}, {2, 2}, {2, 3},
				},
			},
			args: args{
				c: Coord{1, 2},
			},
			want: false,
		},
		{
			name: "cell is live, 2 live neighbors",
			fields: fields{
				liveCells: []Coord{
					{0, 1}, {0, 2}, {0, 3},
					{1, 1}, {1, 2}, {1, 3},
					{2, 1}, {2, 2}, {2, 3},
					{3, 1},
				},
			},
			args: args{
				c: Coord{3, 1},
			},
			want: true,
		},
		{
			name: "cell is dead, 2 live neighbors",
			fields: fields{
				liveCells: []Coord{
					{0, 1}, {0, 2}, {0, 3},
					{1, 1}, {1, 2}, {1, 3},
					{2, 1}, {2, 2}, {2, 3},
				},
			},
			args: args{
				c: Coord{3, 1},
			},
			want: false,
		},
		{
			name: "cell is live, 3 live neighbors",
			fields: fields{
				liveCells: []Coord{
					{0, 1}, {0, 2}, {0, 3},
					{1, 1}, {1, 2}, {1, 3},
					{2, 1}, {2, 2}, {2, 3},
					{3, 2},
				},
			},
			args: args{
				c: Coord{3, 2},
			},
			want: true,
		},
		{
			name: "cell is dead, 3 live neighbors",
			fields: fields{
				liveCells: []Coord{
					{0, 1}, {0, 2}, {0, 3},
					{1, 1}, {1, 2}, {1, 3},
					{2, 1}, {2, 2}, {2, 3},
				},
			},
			args: args{
				c: Coord{3, 2},
			},
			want: true,
		},
		{
			name: "cell is live, 1 live neighbor",
			fields: fields{
				liveCells: []Coord{
					{0, 1}, {0, 2}, {0, 3},
					{1, 1}, {1, 2}, {1, 3},
					{2, 1}, {2, 2}, {2, 3},
					{3, 4},
				},
			},
			args: args{
				c: Coord{3, 4},
			},
			want: false,
		},
		{
			name: "cell is dead, 1 live neighbor",
			fields: fields{
				liveCells: []Coord{
					{0, 1}, {0, 2}, {0, 3},
					{1, 1}, {1, 2}, {1, 3},
					{2, 1}, {2, 2}, {2, 3},
				},
			},
			args: args{
				c: Coord{3, 4},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Grid{
				Length:    5,
				liveCells: sliceToMap(tt.fields.liveCells),
			}
			if got := g.willLive(tt.args.c); got != tt.want {
				t.Errorf("willLive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewGrid(t *testing.T) {
	expectedSeed := []Coord{{0, 1}}

	type args struct {
		seed   []Coord
		length int
	}
	tests := []struct {
		name string
		args args
		want *Grid
	}{
		{
			name: "seeded grid",
			args: args{
				seed:   expectedSeed,
				length: 10,
			},
			want: &Grid{
				Length:    10,
				liveCells: sliceToMap(expectedSeed),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGrid(sliceToMap(tt.args.seed), tt.args.length); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Utility function to convert a slice of Coords to a map of Coords so I don't have to rewrite the test inputs
func sliceToMap(s []Coord) map[Coord]bool {
	m := make(map[Coord]bool)
	for _, v := range s {
		m[v] = true
	}
	return m
}
