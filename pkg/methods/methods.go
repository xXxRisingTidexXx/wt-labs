package methods

import (
	"github.com/xXxRisingTidexXx/wt-labs/pkg/jsonrpc"
)

var Greet = jsonrpc.MethodFunc(
	func(_ jsonrpc.Request) jsonrpc.Response {
		return jsonrpc.NewResult("Hello from JSONRPC!")
	},
)

var Square = jsonrpc.MethodFunc(
	func(request jsonrpc.Request) jsonrpc.Response {
		var params []float64
		if e := request.UnmarshalParams(&params); e != nil {
			return jsonrpc.NewError(e)
		}
		return jsonrpc.NewResult(params[0] * params[0])
	},
)
