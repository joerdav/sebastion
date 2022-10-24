package webui

import (
	"context"
	"log"

	"github.com/joerdav/sebastion"
)

func newWorkerPool(count int, logs LogStore) chan<- startJob {
	c := make(chan startJob)
	for i := 0; i < count; i++ {
		go worker(i, c, logs)
	}
	return c
}

type startJob struct {
	action sebastion.Action
	out    string
}

func worker(idx int, jobs <-chan startJob, logs LogStore) {
	log.Println("Worker", idx, "waiting for work.")
	for j := range jobs {
		log.Println("Worker", idx, "got a job.")
		doWork(j, logs)
		log.Println("Worker", idx, "completed a job.")
	}
}

func doWork(job startJob, logs LogStore) {
	ctx := sebastion.NewContext(context.Background())
	logger, close, err := logs.CreateLogger(job.out)
	if err != nil {
		log.Printf("failed to initialize logger: %v\n", err)
		return
	}
	defer close()
	ctx.Logger = logger
	ctx.Logger.Println("Processing Job", job.action.Details().Name)
	err = job.action.Run(ctx)
	if err != nil {
		ctx.Logger.Println("Error: ", err)
		return
	}
	ctx.Logger.Println("Done.")
}
