
package jsonrpc

type invalidResponse struct {
	payload interface{}
}

func (r invalidResponse) code() int {
	return -32102
}

func (r invalidResponse) message() string {
	return "Invalid Response"
}

func (r invalidResponse) data() interface{} {
	return r.payload
}
