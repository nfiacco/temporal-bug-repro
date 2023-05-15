package app

import (
	"context"
	"fmt"
	"time"
)

type Activities struct {
}

type MyStruct struct {
	Field string
}

func (a *Activities) ReproActivity(ctx context.Context, shouldPanic bool) error {

	done := make(chan bool)
	go func() {
		if shouldPanic {
			var s *MyStruct = nil
			fmt.Println(s.Field) // this line panics from nil pointer dereference
		}

		// sleep for 1 minute then write to channel
		<-time.After(time.Minute * 1)
		done <- true
	}()

	<-done

	return nil
}
