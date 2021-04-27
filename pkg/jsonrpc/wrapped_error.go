package jsonrpc

type wrappedError struct {
	err error
}

func (e wrappedError) code() int {
	return -32000
}

func (e wrappedError) message() string {
	return "Server error"
}

func (e wrappedError) data() interface{} {
	return e.err
}
