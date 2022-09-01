package webui

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/joerdav/sebastion"
	"github.com/joerdav/sebastion/webui/templates"
)

const turboType = "text/vnd.turbo-stream.html"

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
)

type WebConfig struct {
	// Workers defines the amount of concurrent workers waiting to process actions.
	// default: 1
	Workers int
}

type WebRunner struct {
	Actions      []sebastion.Action
	Router       http.Handler
	customInputs []WebInputHandler
	jobs         chan startAction
	outputs      map[string]*bytes.Buffer
}

// AppendHandlers allows you to add custom code to retrieve inputs for an action.
func (t *WebRunner) AppendHandlers(h ...WebInputHandler) {
	t.customInputs = append(t.customInputs, h...)
}

func Web(cfg WebConfig, actions ...sebastion.Action) (http.Handler, error) {
	wr := WebRunner{
		Actions: actions,
		jobs:    make(chan startAction),
		outputs: make(map[string]*bytes.Buffer),
	}
	wr.routes()
	err := validateActions(wr.Actions)
	if err != nil {
		return nil, err
	}
	if cfg.Workers < 1 {
		cfg.Workers = 1
	}
	wr.workers(cfg.Workers, wr.jobs)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL)
		wr.Router.ServeHTTP(w, r)
		log.Println("DONE", r.Method, r.URL)
	}), nil
}

func (wr *WebRunner) routes() {
	r := mux.NewRouter()
	r.HandleFunc("/", wr.index)
	r.HandleFunc("/output/{id}/ws", wr.streamOutput)
	r.HandleFunc("/output/{id}", wr.getOutputView).Methods(http.MethodGet)
	r.HandleFunc("/action/{name}", wr.actionForm).Methods(http.MethodGet)
	r.HandleFunc("/action/{name}", wr.runAction).Methods(http.MethodPost)
	wr.Router = r
}

func (wr *WebRunner) index(w http.ResponseWriter, r *http.Request) {
	err := templates.Index(wr.Actions).Render(r.Context(), w)
	if err != nil {
		w.WriteHeader(500)
	}
}
func (wr *WebRunner) streamOutput(w http.ResponseWriter, r *http.Request) {
	log.Println("Received connection request")
	vars := mux.Vars(r)
	outputId := vars["id"]
	if outputId == "" {
		w.WriteHeader(404)
		return
	}
	_, ok := wr.outputs[outputId]
	if !ok {
		w.WriteHeader(404)
		return
	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	defer ws.Close()
	ticker := time.NewTicker(10 * time.Millisecond)
	for range ticker.C {
		o, ok := wr.outputs[outputId]
		if !ok {
			ws.Close()
			return
		}
		html := new(bytes.Buffer)
		err := templates.LogStream(o.String()).
			Render(r.Context(), html)
		if err != nil {
			ws.Close()
			return
		}
		err = ws.WriteMessage(websocket.TextMessage, html.Bytes())
		if err != nil {
			ws.Close()
			return
		}
	}
}

func (wr *WebRunner) getOutputView(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	outputId := vars["id"]
	if outputId == "" {
		w.WriteHeader(404)
		return
	}
	o := wr.outputs[outputId]
	component := templates.Log(outputId, o.String())
	err := component.Render(r.Context(), w)
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

func (wr *WebRunner) runAction(w http.ResponseWriter, r *http.Request) {
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
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	for _, i := range a.Inputs() {
		h := wr.getInputHandler(i)
		v := r.FormValue(i.Name)
		err := h.Set(i, v)
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			return
		}
	}
	outputId := uuid.NewString()
	wr.jobs <- startAction{
		action: a,
		out:    outputId,
	}
	log.Println("Redirecting")
	http.Redirect(w, r, "/output/"+outputId, 303)
}

func getHandler(hs []WebInputHandler, i sebastion.Input) (WebInputHandler, bool) {
	for _, c := range hs {
		if c.CanHandle(i) {
			return c, true
		}
	}
	return nil, false
}

func (wr *WebRunner) getInputHandler(i sebastion.Input) WebInputHandler {
	if h, ok := getHandler(wr.customInputs, i); ok {
		return h
	}
	if h, ok := getHandler(defaultHandlers, i); ok {
		return h
	}
	return stringInput{}
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
