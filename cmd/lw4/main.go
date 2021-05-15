package main

import (
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/xXxRisingTidexXx/wt-labs/pkg/jsonrpc"
	"github.com/xXxRisingTidexXx/wt-labs/pkg/methods"
	"net/http"
	"time"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
	server := &http.Server{
		Addr: ":9292",
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: 1048576,
		Handler: jsonrpc.NewServer(map[string]jsonrpc.Method{"Greet": methods.Greet}),
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
