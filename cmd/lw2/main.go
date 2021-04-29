package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/xXxRisingTidexXx/wt-labs/pkg/jsonrpc"
	"net/http"
	"time"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
	client := jsonrpc.NewClient(&http.Client{Timeout: 5 * time.Second}, "http://localhost:8484")
	log.Info(
		client.Call(jsonrpc.NewRequest("Pizda", jsonrpc.NewEmpty(), jsonrpc.NewInt(181023))),
	)
	log.Info(
		client.Call(jsonrpc.NewRequest("Pizda", jsonrpc.NewEmpty(), jsonrpc.NewNotification())),
	)
	log.Info(
		client.Call(jsonrpc.NewRequest("Greet", jsonrpc.NewEmpty(), jsonrpc.NewString("abfwefds"))),
	)
}
