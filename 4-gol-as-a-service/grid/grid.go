package grid

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Coord struct {
	X int
	Y int
}

type Dimensions struct {
	Length int `json:"length"`
	Height int `json:"height"`
}

type LiveCellsResponse struct {
	PreviousState []Coord    `json:"previous_state"`
	NextState     []Coord    `json:"next_state"`
	Dimensions    Dimensions `json:"dimensions"`
}

type gridResource struct{}

func (gr *gridResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	return r
}
