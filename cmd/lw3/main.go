package main

import (
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
		Addr:           ":8484",
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: 1048576,
		Handler: jsonrpc.NewServer(
			map[string]jsonrpc.Method{
				"Greet":   methods.Greet,
				"Square":  methods.Square,
				"StoreIP": methods.StoreIP,
			},
		),
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
