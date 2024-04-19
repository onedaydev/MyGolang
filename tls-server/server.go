package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world!")
}

func setupHandler(mux *http.ServeMux) {
	mux.HandleFunc("/api", apiHandler)
}

func loggingMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()
			h.ServeHTTP(w, r)
			log.Printf("protocol=%s path=%s method=%s duration=%f",
				r.Proto, r.URL.Path, r.Method, time.Now().Sub(startTime).Seconds(),
			)
		})
}

func main() {
	listenAddr := os.Getenv("LISTEN_ADDR")
	if len(listenAddr) == 0 {
		listenAddr = ":8443"
	}

	tlsCerfFile := os.Getenv("TLS_CERT_FILE_PATH")
	tlsKeyFile := os.Getenv("TLS_KEY_FILE_PATH")

	if len(tlsCerfFile) == 0 || len(tlsKeyFile) == 0 {
		log.Fatal("TLS_CERT_FILE_PATH and TLS_KEY_FILE_PATH must be specified")
	}

	mux := http.NewServeMux()
	setupHandler(mux)
	m := loggingMiddleware(mux)

	log.Fatal(
		http.ListenAndServeTLS(
			listenAddr,
			tlsCerfFile,
			tlsKeyFile,
			m,
		),
	)
}
