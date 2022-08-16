package ui

import (
	"net/http"

	"github.com/joerdav/sebastion"
	"github.com/joerdav/sebastion/ui/templates"
)

type WebRunner struct {
	Actions []sebastion.Action
	Mux     *http.ServeMux
}

func Web(actions ...sebastion.Action) http.Handler {
	wr := WebRunner{Actions: actions, Mux: new(http.ServeMux)}
	wr.routes()
	return wr.Mux
}

func (wr *WebRunner) routes() {
	wr.Mux.HandleFunc("/", wr.index)
}

func (wr *WebRunner) index(w http.ResponseWriter, r *http.Request) {
	err := templates.Index(wr.Actions).Render(r.Context(), w)
	if err != nil {
		w.WriteHeader(500)
	}
}
