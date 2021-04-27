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

func (r Request) IsNotification() bool {
	_, ok := r.id.toValue()
	return !ok
}

func (r Request) MarshalJSON() ([]byte, error) {
	body := map[string]interface{}{"jsonrpc": Version, "method": r.method}
	if r.params.len() > 0 {
		body["params"] = r.params
	}
	if id, ok := r.id.toValue(); ok {
		body["id"] = id
	}
	return json.Marshal(body)
}
