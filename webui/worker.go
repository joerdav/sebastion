package webui

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"

	"github.com/joerdav/sebastion"
)

func newOutputs() outputs {
	return outputs{make(map[string]io.Reader)}
}

type outputs struct {
	readerMap map[string]io.Reader
}

func (o *outputs) new(outid string) (io.Writer, func()) {
	r, w := io.Pipe()
	o.readerMap[outid] = bufio.NewReader(r)
	return w, func() {
		w.Close()
		delete(o.readerMap, outid)
	}
}

func newWorkerPool(count int, outputs outputs) chan<- startAction {
	c := make(chan startAction)
	for i := 0; i < count; i++ {
		go worker(i, c, outputs)
	}
	return c
}

type startAction struct {
	action sebastion.Action
	out    string
}

func worker(idx int, jobs <-chan startAction, outputs outputs) {
	log.Println("Worker", idx, "waiting for work.")
	for j := range jobs {
		log.Println("Worker", idx, "got a job.")
		bw, close := outputs.new(j.out)
		defer close()
		w := io.MultiWriter(bw, os.Stdout)
		doWork(j, w)
		close()
		log.Println("Worker", idx, "completed a job.")
	}
}

func doWork(job startAction, w io.Writer) {
	ctx := sebastion.NewContext(context.Background())
	ctx.Logger = log.New(w, "", log.LstdFlags)
	ctx.Logger.SetOutput(w)
	ctx.Logger.Println("Processing Job", job.action.Details().Name)
	err := job.action.Run(ctx)
	if err != nil {
		ctx.Logger.Println("Error: ", err)
		return
	}
	ctx.Logger.Println("Done.")
}
