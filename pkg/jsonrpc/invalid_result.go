package jsonrpc

type invalidResult struct {
	err error
}

func (r invalidResult) code() int {
	return -32105
}

func (r invalidResult) message() string {
	return "Invalid result"
}

func (r invalidResult) data() interface{} {
	return r.err
}
