package jsonrpc

import (
	"encoding/json"
)

type Request struct {
	method string
	params Params
	id     ID
}

func NewRequest(method string, params Params, id ID) Request {
	return Request{method, params, id}
}

func (r Request) MarshalJSON() ([]byte, error) {
	body := map[string]interface{}{"jsonrpc": Version, "method": r.method}
	if r.params.len() > 0 {
		body["params"] = r.params
	}
	if value, ok := r.id.toValue(); ok {
		body["id"] = value
	}
	return json.Marshal(body)
}
