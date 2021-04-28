package jsonrpc

type MethodFunc func(Request) Response

func (f MethodFunc) ServeJSONRPC(request Request) Response {
	return f(request)
}
