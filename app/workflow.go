package app

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func ReproWorkflow(ctx workflow.Context, shouldPanic bool) error {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}

	ctx = workflow.WithActivityOptions(ctx, options)

	var a *Activities // Temporal handles calling the registered activity object's methods
	err := workflow.ExecuteActivity(ctx, a.ReproActivity, shouldPanic).Get(ctx, nil)

	return err
}
