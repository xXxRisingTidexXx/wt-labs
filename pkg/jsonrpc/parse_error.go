package jsonrpc

type parseError struct {
	err error
}

func (e parseError) code() int {
	return -32700
}

func (e parseError) message() string {
	return "Parse error"
}

func (e parseError) data() interface{} {
	return e.err
}
