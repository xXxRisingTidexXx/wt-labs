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
	request, _ := jsonrpc.NewRequest("Greet", nil)
	log.Info(client.Call(request))
	request, _ = jsonrpc.NewNotification("Greet", map[string]string{})
	log.Info(client.Call(request))
	request, _ = jsonrpc.NewRequest("Square", nil)
	log.Info(client.Call(request))
	request, _ = jsonrpc.NewRequest("Square", 34)
	log.Info(client.Call(request))
	request, _ = jsonrpc.NewRequest("Square", []float64{1.32})
	log.Info(client.Call(request))
	request, _ = jsonrpc.NewNotification("StoreIP", []string{"12.45.92.102"})
	log.Info(client.Call(request))
}
