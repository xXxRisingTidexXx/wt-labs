package wt

import (
	"github.com/xXxRisingTidexXx/wt-labs/pkg/jsonrpc"
)

type Propagator struct {

}

func (p *Propagator) ServeJSONRPC(_ jsonrpc.Request) jsonrpc.Response {
	return jsonrpc.WithResult("OK")
}
