package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	endpoint := flag.String("endpoint", "", "Endpoint of the server (required and needs to start with 'https')")
	flag.Parse()
	if !strings.HasPrefix(*endpoint, "https") {
		log.Fatalf("Endpoint of the server needs to start with https\n")
	}

	log.Printf("Will make a GET request to %q\n", *endpoint)

	// Load client certificate and private key
	cert, err := tls.LoadX509KeyPair("tls/client.crt", "tls/client.key")
	if err != nil {
		panic(err)
	}

	// Load CA certificate
	caCert, err := os.ReadFile("tls/ca.crt")
	if err != nil {
		panic(err)
	}

	// Create a certificate pool and add the CA certificate to it
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create a TLS configuration
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}

	// Create an HTTP client with the TLS configuration
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	// Make an HTTPS request to the server
	resp, err := client.Get(*endpoint)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Print the response body
	fmt.Println(string(body))
}
