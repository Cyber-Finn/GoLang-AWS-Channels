package main

import (
	"fmt"
	"single_external_Channel_Listener"
)

/*
	this is essentially an example of Step 3 in the Readme.md
*/

func main() {
	// create a channel of type error
	errCh := make(chan error)
	fmt.Println("adding entry")

	// launch a goroutine
	//	every channel needs a goroutine that sends data to it - you cant just place data on the channel
	go func() {
		errCh <- fmt.Errorf("This is error 1")
		errCh <- fmt.Errorf("This is error 2")
		errCh <- fmt.Errorf("This is error 3")
	}()

	fmt.Println("Processing: sending to error-channel's listener..")
	fmt.Println(single_external_Channel_Listener.ReadFromChannel(errCh))
}
