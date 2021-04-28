package jsonrpc

type clientError struct {
	err error
}

func (e clientError) code() int {
	return -32100
}

func (e clientError) message() string {
	return "Client error"
}

func (e clientError) data() interface{} {
	return e.err
}
