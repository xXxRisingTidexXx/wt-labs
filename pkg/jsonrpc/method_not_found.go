package jsonrpc

type methodNotFound struct {
	name string
}

func (m methodNotFound) code() int {
	return -32601
}

func (m methodNotFound) message() string {
	return "Method not found"
}

func (m methodNotFound) data() interface{} {
	return m.name
}
