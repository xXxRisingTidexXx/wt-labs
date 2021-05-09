package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	if err := http.ListenAndServe(":8484", http.HandlerFunc(handleRoot)); err != nil {
		log.Fatal(err)
	}
}

func handleRoot(writer http.ResponseWriter, _ *http.Request) {
	if _, err := writer.Write([]byte("Hello from LW1!\n")); err != nil {
		log.Error(err)
	}
}
