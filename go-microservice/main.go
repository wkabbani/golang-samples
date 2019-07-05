package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/wkabbani/go-samples/go-microservice/home"
	"github.com/wkabbani/go-samples/go-microservice/server"
)

// config vars
var (
	CertFile   string
	KeyFile    string
	ServerAddr string
)

func main() {

	// https://peter.bourgon.org/go-best-practices-2016/#configuration
	readConfig()

	// setup custom logger
	logger := log.New(os.Stdout, "go-microservice", log.LstdFlags|log.Lshortfile)

	// setup database connection
	// db, err := sqlx.Open("", "")
	// if err != nil {
	// 	logger.Fatalln(err)
	// }
	// err := db.Ping()
	// if err != nil {
	// 	logger.Fatalln(err)
	// }

	// setup the handler
	mux := http.NewServeMux()

	// pass home pkg the dependencies it needs
	h := home.NewHomeHanders(logger)
	// to have more control over the package, pass mux to setup routes
	h.SetupRoutes(mux)

	// setup server
	srv := server.New(ServerAddr, mux)

	// run the server and handle error
	logger.Println("running with the following config")
	logger.Printf("Cert path: %s", CertFile)
	logger.Printf("Key path: %s", KeyFile)
	logger.Printf("Address: %s", ServerAddr)

	logger.Println("starting server...")
	err := srv.ListenAndServeTLS(CertFile, KeyFile)
	if err != nil {
		logger.Fatalf("server failed to start: %v", err)
	}
}

func readConfig() {
	flag.StringVar(&CertFile, "c", "server-certs/server.pem", "Server cert file path")
	flag.StringVar(&KeyFile, "k", "server-certs/server-key.pem", "Server key file path")
	flag.StringVar(&ServerAddr, "a", "0.0.0.0:8080", "HTTP service address.")

	flag.Parse()
}
