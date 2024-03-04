package single_external_Channel_Listener

import "fmt"

/*
	this is essentially an example of Step 1-2 in the Readme.md
*/

// creating a function that takes an error channel as an input
func ReadFromChannel(c <-chan error) (resp string) {
	resp = "Success: processed the data on the error channel"
	//defer return;
	fmt.Println(resp)
	// loop until the channel is closed
	for x := range c {
		// print the received value
		fmt.Println(x)
	}
	return resp
}

func main() {

}
