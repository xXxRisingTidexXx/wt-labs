package methods

import (
	"github.com/xXxRisingTidexXx/wt-labs/pkg/jsonrpc"
)

var Greet = jsonrpc.MethodFunc(
	func(_ jsonrpc.Request) jsonrpc.Response {
		return jsonrpc.WithResult("Hello from JSONRPC!")
	},
)

var Square = jsonrpc.MethodFunc(
	func(request jsonrpc.Request) jsonrpc.Response {
		var params []float64
		if e := request.UnmarshalParams(&params); e != nil {
			return jsonrpc.WithError(e)
		}
		if len(params) != 1 {
			return jsonrpc.WithError(jsonrpc.NewInvalidParams("Param number must equal 1"))
		}
		return jsonrpc.WithResult(params[0] * params[0])
	},
)
