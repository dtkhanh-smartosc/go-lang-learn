package main

import (
	"fmt"
	"net/http"
	"time"
)

/*
	HTTP servers are useful for demonstrating the usage of context.
	Context for controlling cancellation. A Context carries deadlines,
 	cancellation signals, and other request-scoped values
  	across API boundaries and goroutines.
*/

func contextFucntion(w http.ResponseWriter, req *http.Request) {
	fmt.Println("------contextFucntion-------")
	// A context.Context is created for each request by the net/http machinery,
	//  and is available with the Context() method.

	ctx := req.Context()
	fmt.Println("server: learn handler started")
	defer fmt.Println("server: learn handler ended")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():

		err := ctx.Err()
		fmt.Println("server", err)
		internalErr := http.StatusInternalServerError
		http.Error(w, err.Error(), internalErr)
	}
}
