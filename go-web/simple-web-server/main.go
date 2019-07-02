package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Go!"))
}

func main() {
	// read the port from cli arguments
	port := flag.Uint("p", 8080, "Port on which the server will listen")
	flag.Parse()

	// request multiplexer
	mux := http.NewServeMux()

	// handle request with a predefined function
	mux.HandleFunc("/hello", helloHandler)

	// handle request with an anonymous function
	mux.HandleFunc("/anonymous", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("generated with an anonymous handler"))
	})

	// serve static files
	fs := http.FileServer(http.Dir(path.Join("client", ".")))
	mux.Handle("/client/", http.StripPrefix("/client", fs))

	// start the server
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), mux)

	// print the error and exit
	fmt.Println(err)
	os.Exit(1)
}
