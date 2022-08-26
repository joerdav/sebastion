package ui

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gorilla/mux"
	"github.com/joerdav/sebastion"
	"github.com/joerdav/sebastion/webui/templates"
)

type WebRunner struct {
	Actions      []sebastion.Action
	Router       http.Handler
	customInputs []WebInputHandler
}

// AppendHandlers allows you to add custom code to retrieve inputs for an action.
func (t *WebRunner) AppendHandlers(h ...WebInputHandler) {
	t.customInputs = append(t.customInputs, h...)
}

func Web(actions ...sebastion.Action) (http.Handler, error) {
	wr := WebRunner{Actions: actions}
	wr.routes()
	err := validateActions(wr.Actions)
	if err != nil {
		return nil, err
	}
	return wr.Router, nil
}

func (wr *WebRunner) routes() {
	r := mux.NewRouter()
	r.HandleFunc("/", wr.index)
	r.HandleFunc("/action/{name}", wr.actionForm)
	wr.Router = r
}

func (wr *WebRunner) index(w http.ResponseWriter, r *http.Request) {
	err := templates.Index(wr.Actions).Render(r.Context(), w)
	if err != nil {
		w.WriteHeader(500)
	}
}

func (wr *WebRunner) actionForm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	if name == "" {
		w.WriteHeader(404)
		return
	}
	a, ok := wr.getActionByName(name)
	if !ok {
		w.WriteHeader(404)
		return
	}
	is := wr.getInputComponents(a)
	err := templates.Action(a, is).Render(r.Context(), w)
	if err != nil {
		w.WriteHeader(500)
	}
}

func getHandler(hs []WebInputHandler, i sebastion.Input) (WebInputHandler, bool) {
	for _, c := range hs {
		if c.CanHandle(i) {
			return c, true
		}
	}
	return nil, false
}

func (wr *WebRunner) getInputComponents(a sebastion.Action) []templ.Component {
	var components []templ.Component
	for _, i := range a.Inputs() {
		if h, ok := getHandler(wr.customInputs, i); ok {
			components = append(components, h.Template(i))
			continue
		}
		if h, ok := getHandler(defaultHandlers, i); ok {
			components = append(components, h.Template(i))
			continue
		}
		components = append(components, templates.StringInput(i))
	}
	return components
}

func (wr WebRunner) getActionByName(name string) (sebastion.Action, bool) {
	for _, a := range wr.Actions {
		if name == a.Details().Name {
			return a, true
		}
	}
	return nil, false
}

func validateActions(as []sebastion.Action) error {
	// All slugs must be unique
	slugs := map[string]bool{}
	for _, a := range as {
		if slugs[a.Details().Name] {
			return fmt.Errorf("actions must have unique names")
		}
		slugs[a.Details().Name] = true
	}
	return nil
}
