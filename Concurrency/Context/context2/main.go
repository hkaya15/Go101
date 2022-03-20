package main

import "context"

func main() {
	context.Background() // It returns empty Context
	context.TODO()       // It returns empty Context. It intended purpose is to serve as a placeholder for when you don't know which Context to utilize, or to be provided.
	//context.WithCancel(parent Context) (ctx Context,cancel CancelFunc) // returns a new Context that closes its done channel when the returned cancel function is called
	//context.WithDeadline(parent Context, deadline time.Time) (Context,CancelFunc) // returns a new Context that closes its done channel when the machine's clock advances past the given given deadline
	//context.WithTimeout(parent Context,timeout time.Duration) (Context, CancelFunc) // returns a new Context that closes its done channel after the given timeout duration
}
