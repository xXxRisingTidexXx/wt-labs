package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	if err := http.ListenAndServe(":8484", http.HandlerFunc(handleRoot)); err != nil {
		log.Fatalf("main: lw1 failed to run the server, %v", err)
	}
}

func handleRoot(writer http.ResponseWriter, _ *http.Request) {
	if _, err := writer.Write([]byte("Hello from LW1!\n")); err != nil {
		log.Errorf("main: lw1 failed to write response, %v", err)
	}
}
