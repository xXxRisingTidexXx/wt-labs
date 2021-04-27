package jsonrpc

type invalidRequest struct {
	payload interface{}
}

func (r invalidRequest) code() int {
	return -32600
}

func (r invalidRequest) message() string {
	return "Invalid Request"
}

func (r invalidRequest) data() interface{} {
	return r.payload
}
