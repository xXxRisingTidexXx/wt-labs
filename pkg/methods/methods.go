package methods

import (
	"github.com/xXxRisingTidexXx/wt-labs/pkg/jsonrpc"
	"os"
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

var StoreIP = jsonrpc.MethodFunc(
	func(request jsonrpc.Request) jsonrpc.Response {
		var params []string
		if e := request.UnmarshalParams(&params); e != nil {
			return jsonrpc.WithError(e)
		}
		if len(params) != 1 {
			return jsonrpc.WithError(jsonrpc.NewInvalidParams("Param number must equal 1"))
		}
		file, err := os.OpenFile("registry.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return jsonrpc.WithError(jsonrpc.NewStructuredError(1000, "Append failed", err))
		}
		if _, err := file.WriteString(params[0] + "\n"); err != nil {
			return jsonrpc.WithError(jsonrpc.NewStructuredError())
		}
		return jsonrpc.WithResult("OK")
	},
)
