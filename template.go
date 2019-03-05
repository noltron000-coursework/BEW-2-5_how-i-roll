package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	////////////////////////
	// r is request.      //
	// w is writers       //
	//  (& response?)     //
	// ------------------ //
	// need more research //
	// on these meanings. //
	////////////////////////

	fmt.Fprintf(w, "Hello World")
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8080", nil)
}
