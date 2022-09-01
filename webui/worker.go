package webui

import (
	"bytes"
	"context"
	"io"
	"log"
	"os"

	"github.com/joerdav/sebastion"
)

type startAction struct {
	action sebastion.Action
	out    string
}

func (wr *WebRunner) workers(count int, jobs <-chan startAction) {
	for i := 0; i < count; i++ {
		go wr.worker(i, jobs)
	}
}

func (wr *WebRunner) worker(idx int, jobs <-chan startAction) {
	log.Println("Worker", idx, "waiting for work.")
	for j := range jobs {
		log.Println("Worker", idx, "got a job.")
		bw, close := wr.newOut(j.out)
		defer close()
		w := io.MultiWriter(bw, os.Stdout)
		ctx := sebastion.NewContext(context.Background())
		ctx.Logger = log.New(w, "", log.LstdFlags)
		ctx.Logger.SetOutput(w)
		ctx.Logger.Println("Processing Job", j.action.Details().Name)
		err := j.action.Run(ctx)
		if err != nil {
			ctx.Logger.Println("Error: ", err)
		}
		ctx.Logger.Println("Done.", err)
		close()
	}
}

func (wr *WebRunner) newOut(outid string) (io.Writer, func()) {
	b := new(bytes.Buffer)
	wr.outputs[outid] = b
	return b, func() {
		delete(wr.outputs, outid)
	}
}
