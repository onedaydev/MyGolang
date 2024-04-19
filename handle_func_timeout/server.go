package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func handleUserAPI(w http.ResponseWriter, r *http.Request) {
	log.Println("I started processing the request")
	time.Sleep(15 * time.Second)

	log.Println(
		"before continuing, i will check if the timeout has already expired",
	)

	if r.Context().Err() != nil {
		log.Printf(
			"Aborting further processing: %v\n",
			r.Context().Err(),
		)
		return
	}

	fmt.Fprint(w, "Hello world!")
	log.Println("I finished processing the request")
}

func main() {
	listenAddr := os.Getenv("LISTEN_ADDR")
	if len(listenAddr) == 0 {
		listenAddr = ":8080"
	}
	timeoutDuration := 14 * time.Second

	userHandler := http.HandlerFunc(handleUserAPI)
	hTimeout := http.TimeoutHandler(
		userHandler,
		timeoutDuration,
		"I ran out of time\n",
	)

	mux := http.NewServeMux()
	mux.Handle("/api/users/", hTimeout)

	log.Fatal(http.ListenAndServe(listenAddr, mux))
}
