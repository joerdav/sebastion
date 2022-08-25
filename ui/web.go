package ui

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/joerdav/sebastion"
	"github.com/joerdav/sebastion/ui/templates"
)

type WebRunner struct {
	Actions []sebastion.Action
	Mux     *http.ServeMux
}

func Web(actions ...sebastion.Action) (http.Handler, error) {
	wr := WebRunner{Actions: actions, Mux: new(http.ServeMux)}
	wr.routes()
	err := validateActions(wr.Actions)
	if err != nil {
		return nil, err
	}
	return wr.Mux, nil
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

func (wr WebRunner) getActionBySlug(slug string) (sebastion.Action, bool) {
	for _, a := range wr.Actions {
		n, _ := a.Details()
		if slug == actionNameToSlug(n) {
			return a, true
		}
	}
	return nil, false
}

func actionNameToSlug(name string) string {
	return url.PathEscape(name)
}

func validateActions(as []sebastion.Action) error {
	slugs := map[string]bool{}
	for _, a := range as {
		n, _ := a.Details()
		s := actionNameToSlug(n)
		if slugs[s] {
			return fmt.Errorf("actions must have unique names")
		}
		slugs[s] = true
	}
	return nil
}
