package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Go!"))
}

func main() {
	p := flag.Uint("p", 8080, "The port on which the server will listen")
	flag.Parse()

	// receives a signal from the OS if the process should be ended
	stop := make(chan os.Signal, 1)
	signal.Notify(stop)

	// setup the default http server
	http.HandleFunc("/hello", helloHandler)
	h := &http.Server{Addr: fmt.Sprintf(":%d", *p)}

	// start the web server in a separate routine
	go func() {
		fmt.Println("starting the server...")
		log.Fatal(h.ListenAndServe())
	}()

	// block here until we receive a stop signal from the os
	<-stop

	fmt.Println("shutting down gracefully...")
	h.Shutdown(context.Background())
}
