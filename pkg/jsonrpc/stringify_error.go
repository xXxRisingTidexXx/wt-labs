package jsonrpc

type stringifyError struct {
	err error
}

func (e stringifyError) code() int {
	return -32104
}

func (e stringifyError) message() string {
	return "Stringify error"
}

func (e stringifyError) data() interface{} {
	return e.err
}
