package main

import "time"


func main(){

}



type Context interface{

	// Deadline returns the time when work done on behalf of this context should be canceled.
	Deadline() (deadline time.Time,ok bool)
	
	// Done returns a channel that's closed when work done on behalf of this context should be canceled.
	// Done may return nil if this context can never be canceled.
	Done() <-chan struct{}

	// Err returns a non-nil error value after Done is closed. Err returns Canceled if the context was canceled
	// or DeadlineExceed if the context's deadline passed. No other values for Err are defined.
	Err() error

	// Value returns the value associated with context for key or nil if no value is associated with key.
	Value(key interface{}) interface{}
}

