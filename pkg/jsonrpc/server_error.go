package jsonrpc

type serverError struct {
	err error
}

func (e serverError) code() int {
	return -32000
}

func (e serverError) message() string {
	return "Server error"
}

func (e serverError) data() interface{} {
	return e.err
}
