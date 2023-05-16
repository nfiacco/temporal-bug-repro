package app

import (
	"context"
	"fmt"
	"time"

	"gopkg.in/errgo.v2/fmt/errors"
)

type Activities struct {
}

type MyStruct struct {
	Field string
}

func (a *Activities) ReproActivity(ctx context.Context, shouldPanic bool) error {

	doneC := make(chan bool)
	errC := make(chan error)

	go safeCall(func() {
		if shouldPanic {
			var s *MyStruct = nil
			fmt.Println(s.Field) // this line panics from nil pointer dereference
		}

		// sleep for 1 minute then write to channel
		<-time.After(time.Minute * 1)
		doneC <- true
	}, errC)

	select {
	case err := <-errC:
		return err
	case <-doneC:
		return nil
	}
}

// Any new goroutines can crash the whole worker unless we recover from panics
func safeCall(fn func(), errC chan<- error) {
	defer func() {
		if r := recover(); r != nil {
			errC <- errors.Newf("panic: %v", r)
		}
	}()

	fn()
}
