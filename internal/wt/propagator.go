package wt

import (
	"database/sql"
	"github.com/xXxRisingTidexXx/wt-labs/pkg/jsonrpc"
	"net"
	"net/http"
	"time"
)

type Propagator struct {
	source  string
	targets []string
	db      *sql.DB
	clients map[string]*jsonrpc.Client
}

func NewPropagator(source string, targets []string, db *sql.DB) *Propagator {
	clients := make(map[string]*jsonrpc.Client, len(targets))
	for _, target := range targets {
		clients[target] = jsonrpc.NewClient(
			&http.Client{Timeout: 5 * time.Second},
			"http://wt-app-"+target+"/",
		)
	}
	return &Propagator{source, targets, db, clients}
}

func (p *Propagator) ServeJSONRPC(request jsonrpc.Request) jsonrpc.Response {
	var ips []net.IP
	if e := request.UnmarshalParams(&ips); e != nil {
		return jsonrpc.WithError(e)
	}
	if len(ips) != 1 {
		return jsonrpc.WithError(jsonrpc.NewInvalidParams("IP number must equal 1"))
	}
	return jsonrpc.WithResult(p.source)
}
