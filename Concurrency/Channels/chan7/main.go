package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)

	//unb := make(chan int)
	buf := make(chan int, 2)

	go func() {

		//defer close(unb)
		defer close(buf)
		//defer fmt.Fprintln(&stdoutBuff, "Done")
		// for i := 0; i < 4; i++ {
		// 	fmt.Fprintf(&stdoutBuff, "Value unbuffered added %d\n", i)
		// 	unb <- i
		// }
		for i := 0; i < 9; i++ {
			//fmt.Fprintf(&stdoutBuff, "Value buffered added %d\n", i)
			buf <- i
		}

	}()

	// for res := range unb {
	// 	fmt.Fprintf(&stdoutBuff,"Unbuffered : %v\n", res)
	// }

	for res1 := range buf {
		fmt.Fprintf(&stdoutBuff, "Buffered : %v\n", res1)
	}

}
