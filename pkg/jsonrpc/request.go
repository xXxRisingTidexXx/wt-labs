package jsonrpc

import (
	"encoding/json"
	"io"
)

type Request struct {
	method string
	params Params
	id     ID
}

func NewRequest(method string, params Params, id ID) Request {
	return Request{method, params, id}
}

func ParseRequest(reader io.Reader) (Request, Error) {
	var (
		body    map[string]interface{}
		request Request
	)
	if err := json.NewDecoder(reader).Decode(&body); err != nil {
		return request, parseError{err}
	}
	if version, ok := body["jsonrpc"]; !ok || version != Version {
		return request, invalidRequest{"Field \"jsonrpc\" is either absent or invalid"}
	}
	delete(body, "jsonrpc")
	method, ok := body["method"]
	if !ok {
		return request, invalidRequest{"Field \"method\" is absent"}
	}
	if request.method, ok = method.(string); !ok {
		return request, invalidRequest{"Field \"method\" is not a string"}
	}
	delete(body, "method")
	if params, ok := body["params"]; ok {
		switch params := params.(type) {
		case []interface{}:
			request.params = positionalParams(params)
		case map[string]interface{}:
			request.params = namedParams(params)
		default:
			return request, invalidRequest{"Field \"params\" is neither array nor object"}
		}
		delete(body, "params")
	} else {
		request.params = emptyParams{}
	}
	if id, ok := body["id"]; ok {
		switch id := id.(type) {
		case float64:
			request.id = numberID(id)
		case string:
			request.id = stringID(id)
		case nil:
			request.id = nullID{}
		default:
			return request, invalidRequest{"Field \"id\" is neither number nor string nor null"}
		}
		delete(body, "id")
	} else {
		request.id = notificationID{}
	}
	if len(body) > 0 {
		return request, invalidRequest{"Request contains extra fields"}
	}
	return request, nil
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
