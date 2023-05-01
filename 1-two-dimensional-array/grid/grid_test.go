package grid

import (
	"reflect"
	"testing"
)

func TestGrid_Seed(t *testing.T) {
	//unusedCoords := util.Pointer(make([]Coord, 0))

	type args struct {
		cells []Coord
	}
	tests := []struct {
		name    string
		args    args
		want    Grid
		wantErr bool
	}{
		{
			name: "Seed() works on 2x2 grid",
			args: args{
				cells: []Coord{
					{0, 0},
					{1, 1},
				},
			},
			want: Grid{
				grid: [][]bool{
					{true, false},
					{false, true},
				},
			},
			wantErr: false,
		},
		{
			name: "Seed() throws out of bounds error and does not modify grid",
			args: args{
				cells: []Coord{
					{0, 0},
					{31337, 31337},
				},
			},
			want: Grid{
				grid: [][]bool{
					{false, false},
					{false, false},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGrid(2)
			err := g.Seed(tt.args.cells)
			if (err != nil) != tt.wantErr {
				t.Errorf("Seed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got := g.grid
			if !reflect.DeepEqual(got, tt.want.grid) {
				t.Errorf("Seed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCoord_validateBounds(t *testing.T) {
	type fields struct {
		x uint16
		y uint16
	}
	type args struct {
		length uint16
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "validateBounds() returns no error when coord is within bounds",
			fields: fields{
				x: 0,
				y: 1,
			},
			args: args{
				length: 2,
			},
		},
		{
			name: "validateBounds() returns error when coord.X is out of bounds",
			fields: fields{
				x: 100,
				y: 1,
			},
			args: args{
				length: 2,
			},
			wantErr: true,
		},
		{
			name: "validateBounds() returns error when coord.Y is out of bounds",
			fields: fields{
				x: 1,
				y: 100,
			},
			args: args{
				length: 2,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Coord{
				X: tt.fields.x,
				Y: tt.fields.y,
			}
			if err := c.validateBounds(tt.args.length); (err != nil) != tt.wantErr {
				t.Errorf("validateBounds() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGrid_GetGrid(t *testing.T) {
	type fields struct {
		Length        uint16
		grid          [][]bool
		markedForLife []Coord
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "GetGrid() returns a copy of the grid",
			fields: fields{
				grid: [][]bool{
					{true, false},
					{false, true},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Grid{
				Length:        2,
				grid:          tt.fields.grid,
				markedForLife: []Coord{},
			}
			result := g.GetGrid()
			if &result == &tt.fields.grid {
				t.Errorf("GetGrid() pointer values should not match, got %v, original %v", result, tt.fields.grid)
			}
			if !reflect.DeepEqual(result, tt.fields.grid) {
				t.Errorf("GetGrid() = %v, want %v", result, tt.fields.grid)
			}
		})
	}
}

func TestGrid_Tick(t *testing.T) {
	type fields struct {
		Length        uint16
		grid          [][]bool
		markedForLife []Coord
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Grid{
				Length:        tt.fields.Length,
				grid:          tt.fields.grid,
				markedForLife: tt.fields.markedForLife,
			}
			g.Tick()
		})
	}
}

func TestGrid_checkForLife(t *testing.T) {
	var testCoord = Coord{0, 0}
	type fields struct {
		grid [][]bool
	}
	type args struct {
		coord            Coord
		numLiveNeighbors int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []Coord
	}{
		{
			name:   "checkForLife() marks a living cell for life when it has exactly 2 living neighbors",
			fields: fields{grid: [][]bool{{true}}},
			args:   args{coord: testCoord, numLiveNeighbors: 2},
			want:   []Coord{testCoord},
		},
		{
			name:   "checkForLife() marks a living cell for life when it has exactly 3 living neighbors",
			fields: fields{grid: [][]bool{{true}}},
			args:   args{coord: testCoord, numLiveNeighbors: 3},
			want:   []Coord{testCoord},
		},
		{
			name:   "checkForLife() does not mark a living cell for life when it has exactly 4 living neighbors",
			fields: fields{grid: [][]bool{{true}}},
			args:   args{coord: testCoord, numLiveNeighbors: 4},
			want:   []Coord{},
		},
		{
			name:   "checkForLife() does not mark a living cell for life when it has exactly 1 living neighbors",
			fields: fields{grid: [][]bool{{true}}},
			args:   args{coord: testCoord, numLiveNeighbors: 1},
			want:   []Coord{},
		},
		{
			name:   "checkForLife() does not mark a dead cell for life when it has exactly 1 living neighbors",
			fields: fields{grid: [][]bool{{false}}},
			args:   args{coord: testCoord, numLiveNeighbors: 1},
			want:   []Coord{},
		},
		{
			name:   "checkForLife() does not mark a dead cell for life when it has exactly 4 living neighbors",
			fields: fields{grid: [][]bool{{true}}},
			args:   args{coord: testCoord, numLiveNeighbors: 4},
			want:   []Coord{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGrid(1)
			g.grid = tt.fields.grid
			g.checkForLife(testCoord, tt.args.numLiveNeighbors)
			if !reflect.DeepEqual(g.markedForLife, tt.want) {
				t.Errorf("Seed() = %v, want %v", g.markedForLife, tt.want)
			}
		})
	}
}

func TestGrid_liveNeighbors(t *testing.T) {
	type fields struct {
		grid [][]bool
	}
	type args struct {
		coord Coord
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "liveNeighbors() returns 0 when there are no live cells",
			fields: fields{
				grid: [][]bool{
					{false, false, false},
					{false, false, false},
					{false, false, false},
				},
			},
			args: args{
				coord: Coord{1, 1},
			},
			want: 0,
		},
		{
			name: "liveNeighbors() returns 3 when there are 3 live neighbors around 1,1",
			fields: fields{
				grid: [][]bool{
					{true, false, false},
					{false, false, true},
					{false, true, false},
				},
			},
			args: args{
				coord: Coord{1, 1},
			},
			want: 3,
		},
		{
			name: "liveNeighbors() returns 4 when there are 4 live neighbors around 1,1 and 1,1 is alive",
			fields: fields{
				grid: [][]bool{
					{true, true, false},
					{false, true, true},
					{false, true, false},
				},
			},
			args: args{
				coord: Coord{1, 1},
			},
			want: 4,
		},
		{
			name: "liveNeighbors() returns 3 when there are 4 live neighbors around 0,2",
			fields: fields{
				grid: [][]bool{
					{true, true, true},
					{false, true, true},
					{false, false, false},
				},
			},
			args: args{
				coord: Coord{0, 2},
			},
			want: 3,
		},
		{
			name: "liveNeighbors() returns 2 when there are 4 live neighbors around 2,2",
			fields: fields{
				grid: [][]bool{
					{false, false, false},
					{false, true, true},
					{false, false, false},
				},
			},
			args: args{
				coord: Coord{2, 2},
			},
			want: 2,
		},
		{
			name: "liveNeighbors() returns 0 when there are 0 live neighbors around 0,2 and 0,2 is alive",
			fields: fields{
				grid: [][]bool{
					{false, false, false},
					{false, false, false},
					{true, false, false},
				},
			},
			args: args{
				coord: Coord{0, 2},
			},
			want: 0,
		},
		{
			name: "liveNeighbors() returns 0 when there are 0 live neighbors around 0,0 and 0,0 is alive",
			fields: fields{
				grid: [][]bool{
					{true, false, false},
					{false, false, false},
					{false, false, false},
				},
			},
			args: args{
				coord: Coord{0, 0},
			},
			want: 0,
		},
		{
			name: "liveNeighbors() returns error when Coord is out of bounds",
			fields: fields{
				grid: [][]bool{
					{false, false, false},
					{false, false, false},
					{false, false, false},
				},
			},
			args: args{
				coord: Coord{4, 4},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Grid{
				Length:        3,
				grid:          tt.fields.grid,
				markedForLife: []Coord{},
			}
			got, err := g.liveNeighbors(tt.args.coord)
			if (err != nil) != tt.wantErr {
				t.Errorf("liveNeighbors() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("liveNeighbors() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrid_isCoordLive(t *testing.T) {
	type fields struct {
		grid [][]bool
	}
	type args struct {
		t Coord
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "return false if coord.X is out of bounds",
			fields: fields{
				grid: [][]bool{
					{true, true, true},
					{true, true, true},
					{true, true, true},
				},
			},
			args: args{t: Coord{3, 0}},
			want: false,
		},
		{
			name: "return false is coord.Y is out of bounds",
			fields: fields{
				grid: [][]bool{
					{true, true, true},
					{true, true, true},
					{true, true, true},
				},
			},
			args: args{t: Coord{3, 0}},
			want: false,
		},
		{
			name: "return false is coord 2,0 is not live",
			fields: fields{
				grid: [][]bool{
					{true, true, true},
					{true, true, true},
					{false, true, true},
				},
			},
			args: args{t: Coord{2, 0}},
			want: false,
		},
		{
			name: "return false if coord 0,2 is not live",
			fields: fields{
				grid: [][]bool{
					{true, true, false},
					{true, true, true},
					{true, true, true},
				},
			},
			args: args{t: Coord{0, 2}},
			want: false,
		},
		{
			name: "return true is coord 2,0 is live",
			fields: fields{
				grid: [][]bool{
					{false, false, false},
					{false, false, false},
					{true, false, false},
				},
			},
			args: args{t: Coord{2, 0}},
			want: true,
		},
		{
			name: "return true if coord 0,2 is live",
			fields: fields{
				grid: [][]bool{
					{false, false, true},
					{false, false, false},
					{false, false, false},
				},
			},
			args: args{t: Coord{0, 2}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Grid{
				Length:        3,
				grid:          tt.fields.grid,
				markedForLife: []Coord{},
			}
			if got := g.isCoordLive(tt.args.t); got != tt.want {
				t.Errorf("isCoordLive() = %v, want %v", got, tt.want)
			}
		})
	}
}
