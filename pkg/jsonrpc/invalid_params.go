package jsonrpc

import (
	"fmt"
)

type invalidParams struct {
	err error
}

func NewInvalidParams(message string) Error {
	return invalidParams{fmt.Errorf(message)}
}

func (p invalidParams) code() int {
	return -32602
}

func (p invalidParams) message() string {
	return "Invalid params"
}

func (p invalidParams) data() interface{} {
	return p.err.Error()
}
