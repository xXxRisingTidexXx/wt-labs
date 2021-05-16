package wt

import (
	"database/sql"
	"github.com/xXxRisingTidexXx/wt-labs/pkg/jsonrpc"
)

type Propagator struct {
	source  string
	targets []string
	db      *sql.DB
}

func NewPropagator(source string, targets []string, db *sql.DB) *Propagator {
	return &Propagator{source, targets, db}
}

func (p *Propagator) ServeJSONRPC(_ jsonrpc.Request) jsonrpc.Response {
	return jsonrpc.WithResult(p.source)
}
