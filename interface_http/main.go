package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	//fmt.Println(resp)

	/*
		// make is a built in function that
		// takes a type of slice, map or chan
		// and initialises it WITH A SPECIFIC
		// LENGTH and sets it to empty
		bs := make([]byte, 99999)
		resp.Body.Read(bs)
		fmt.Println(string(bs))
	*/

	// do as above but on one line using io.Copy (more succinct)
	// io.Copy(os.Stdout, resp.Body)

	// write our own type that implements the Writer interface
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

// promoting struct logWriter to type Writer interface by
// implementing the Write function
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote", len(bs), "bytes")
	return len(bs), nil
}
