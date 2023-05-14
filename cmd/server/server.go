package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := flag.Int("port", 1234, "Port to listen to")
	server := flag.String("server-id", "", "Id of this server (required)")

	flag.Parse()
	if *server == "" {
		flag.Usage()
		log.Fatal("server-id is not provided")
	}

	http.HandleFunc("/backend/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Pong from %q\n", *server)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Up\n")
	})

	log.Printf("Server %q is up on localhost:%d\n", *server, *port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
		log.Fatalf("Error starting server %q: %v\n", *server, err)
	}
}
