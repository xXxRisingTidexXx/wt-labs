package methods

import (
	"github.com/xXxRisingTidexXx/wt-labs/pkg/jsonrpc"
)

var Greet = jsonrpc.MethodFunc(
	func(_ jsonrpc.Request) jsonrpc.Response {
		return jsonrpc.NewResult("Hello from JSONRPC!")
	},
)