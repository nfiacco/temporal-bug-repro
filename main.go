package main

import (
	"log"

	"temporal-bug-repro/app"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "GoroutinePanicTest", worker.Options{})

	activities := &app.Activities{}

	w.RegisterWorkflow(app.ReproWorkflow)
	w.RegisterActivity(activities.ReproActivity)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
