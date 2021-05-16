package wt

import (
	"github.com/xXxRisingTidexXx/wt-labs/pkg/jsonrpc"
)

func newIPStoringError(err error) jsonrpc.Error {
	return jsonrpc.NewStructuredError(1001, "IP storing error", err)
}
