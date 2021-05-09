package methods

import (
	"github.com/xXxRisingTidexXx/wt-labs/pkg/jsonrpc"
	"os"
	"time"
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
		file, err := os.OpenFile("jsonrpc.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return jsonrpc.WithError(jsonrpc.NewStructuredError(1001, "Opening failed", err))
		}
		_, err = file.WriteString(time.Now().Format(time.RFC3339) + " " + params[0] + "\n")
		if err != nil {
			_ = file.Close()
			return jsonrpc.WithError(jsonrpc.NewStructuredError(1002, "Writing failed", err))
		}
		if err := file.Close(); err != nil {
			return jsonrpc.WithError(jsonrpc.NewStructuredError(1002, "Closing failed", err))
		}
		return jsonrpc.WithResult("OK")
	},
)
