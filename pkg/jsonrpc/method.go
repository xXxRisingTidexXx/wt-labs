package jsonrpc

type Method interface {
	ServeJSONRPC(Request) Response
}
