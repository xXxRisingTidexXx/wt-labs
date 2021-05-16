package main

import (
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/xXxRisingTidexXx/wt-labs/internal/config"
	"github.com/xXxRisingTidexXx/wt-labs/internal/wt"
	"github.com/xXxRisingTidexXx/wt-labs/pkg/jsonrpc"
	"net/http"
	"time"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
	_, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	server := &http.Server{
		Addr:           ":80",
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: 1048576,
		Handler: jsonrpc.NewServer(
			map[string]jsonrpc.Method{"PropagateIP": &wt.Propagator{}},
		),
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}