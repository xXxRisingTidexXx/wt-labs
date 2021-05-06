package jsonrpc

type invalidParams struct {
	err error
}

func (p invalidParams) code() int {
	return -32602
}

func (p invalidParams) message() string {
	return "Invalid params"
}

func (p invalidParams) data() interface{} {
	return p.err
}
