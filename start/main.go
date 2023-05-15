package main

import (
	"context"
	"log"
	"temporal-bug-repro/app"

	"go.temporal.io/sdk/client"
)

func main() {
	// Create the client object just once per process
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	option1 := client.StartWorkflowOptions{
		ID:        "wait-workflow",
		TaskQueue: "GoroutinePanicTest",
	}

	// Start the first workflow that waits for 1 min
	c.ExecuteWorkflow(context.Background(), option1, app.ReproWorkflow, false)
	if err != nil {
		log.Fatalln("unable to complete Workflow", err)
	}

	option2 := client.StartWorkflowOptions{
		ID:        "panic-workflow",
		TaskQueue: "GoroutinePanicTest",
	}

	// Start the second workflow that panics
	c.ExecuteWorkflow(context.Background(), option2, app.ReproWorkflow, true)
	if err != nil {
		log.Fatalln("unable to complete Workflow", err)
	}
}
