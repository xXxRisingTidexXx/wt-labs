package main

import (
	"database/sql"
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
	c, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	db, err := sql.Open("postgres", c.DSN)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		_ = db.Close()
		log.Fatal(err)
	}
	method := "PropagateIP"
	server := &http.Server{
		Addr:           ":80",
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: 1048576,
		Handler: jsonrpc.NewServer(
			map[string]jsonrpc.Method{
				method: wt.NewPropagator(method, c.Node, c.Graph[c.Node].ToSlice(), db),
			},
		),
	}
	if err := server.ListenAndServe(); err != nil {
		_ = db.Close()
		log.Fatal(err)
	}
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
}
