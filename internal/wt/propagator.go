package wt

import (
	"database/sql"
	log "github.com/sirupsen/logrus"
	"github.com/xXxRisingTidexXx/wt-labs/internal/config"
	"github.com/xXxRisingTidexXx/wt-labs/pkg/jsonrpc"
	"net"
	"net/http"
	"time"
)

type Propagator struct {
	method  string
	source  string
	db      *sql.DB
	clients map[string]*jsonrpc.Client
}

func NewPropagator(method, source string, targets []string, db *sql.DB) *Propagator {
	clients := make(map[string]*jsonrpc.Client, len(targets))
	for _, target := range targets {
		clients[target] = jsonrpc.NewClient(
			&http.Client{Timeout: 5 * time.Second},
			"http://wt-app-"+target+"/",
		)
	}
	return &Propagator{method, source, db, clients}
}

func (p *Propagator) ServeJSONRPC(request jsonrpc.Request) jsonrpc.Response {
	var ips []net.IP
	if e := request.UnmarshalParams(&ips); e != nil {
		return jsonrpc.WithError(e)
	}
	if len(ips) != 1 {
		return jsonrpc.WithError(jsonrpc.NewInvalidParams("IP number must equal 1"))
	}
	log.Info(ips[0])
	go p.propagateIP(ips[0])
	return jsonrpc.WithResult(p.source)
}

func (p *Propagator) propagateIP(ip net.IP) {
	targets, err := p.storeIP(ip)
	if err != nil {
		jsonrpc.LogError(newIPStoringError(err))
	} else {
		for _, target := range targets {
			go p.sendIP(target, ip)
		}
	}
}

func (p *Propagator) storeIP(ip net.IP) ([]string, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return nil, err
	}
	rows, err := tx.Query("select node from propagations where ip = $1", ip.String())
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}
	nodes := make(config.Set, 0)
	for rows.Next() {
		var node string
		if err := rows.Scan(&node); err != nil {
			_ = rows.Close()
			_ = tx.Rollback()
			return nil, err
		}
		nodes.Add(node)
	}
	if err := rows.Err(); err != nil {
		_ = rows.Close()
		_ = tx.Rollback()
		return nil, err
	}
	if err := rows.Close(); err != nil {
		_ = tx.Rollback()
		return nil, err
	}
	if nodes.Has(p.source) {
		if err := tx.Commit(); err != nil {
			return nil, err
		}
		return nil, nil
	}
	_, err = tx.Exec(
		"insert into propagations(node, ip) values ($1, $2)",
		p.source,
		ip.String(),
	)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	targets := make([]string, 0)
	for target := range p.clients {
		if !nodes.Has(target) {
			targets = append(targets, target)
		}
	}
	return targets, nil
}

func (p *Propagator) sendIP(target string, ip net.IP) {
	if e := p.sendIPWithError(target, ip); e != nil {
		jsonrpc.LogError(e)
	}
}

func (p *Propagator) sendIPWithError(target string, ip net.IP) jsonrpc.Error {
	request, e := jsonrpc.NewRequest(p.method, []net.IP{ip})
	if e != nil {
		return e
	}
	response := p.clients[target].Call(request)
	if e := response.Error(); e != nil {
		return e
	}
	var node string
	if e := response.UnmarshalResult(&node); e != nil {
		return e
	}
	if target != node {
		return jsonrpc.NewStructuredError(
			1002,
			"Node mismatch",
			map[string]string{"actual": node, "expected": target},
		)
	}
	return nil
}
