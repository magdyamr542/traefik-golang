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
	role := flag.String("role", "backend", "Role of the server. (backend,frontend)")

	flag.Parse()
	if *server == "" {
		flag.Usage()
		log.Fatal("server-id is not provided")
	}

	if !(*role == "backend" || *role == "frontend") {
		flag.Usage()
		log.Fatal("role is invalid")
	}

	serverName := fmt.Sprintf("%s-%s", *role, *server)
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Got new %s request at %s\n", r.Method, serverName)
		fmt.Fprintf(w, "Pong from %s\n", serverName)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Up\n")
	})

	log.Printf("Server %s is up on localhost:%d\n", serverName, *port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
		log.Fatalf("Error starting server %q: %v\n", *server, err)
	}
}
