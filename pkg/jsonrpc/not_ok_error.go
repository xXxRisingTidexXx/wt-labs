package jsonrpc

type notOKError struct {
	status string
}

func (e notOKError) code() int {
	return -32101
}

func (e notOKError) message() string {
	return "Not OK status"
}

func (e notOKError) data() interface{} {
	return e.status
}
