package wt

import (
	"github.com/xXxRisingTidexXx/wt-labs/pkg/jsonrpc"
)

func newMessageStoringError(err error) jsonrpc.Error {
	return jsonrpc.NewStructuredError(1001, "Message storing error", err)
}
