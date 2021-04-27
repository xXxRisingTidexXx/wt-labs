package jsonrpc

import (
	"encoding/json"
	"io"
)

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
	switch method := method.(type) {
	case string:
		request.method = method
	default:
		return request, invalidRequest{"Field \"method\" is not string"}
	}
	delete(body, "method")
	params, ok := body["params"]
	if !ok {
		return request, invalidRequest{"Field \"params\" is absent"}
	}
	switch params := params.(type) {
	case []interface{}:
		request.params = positionalParams(params)
	case map[string]interface{}:
		request.params = namedParams(params)
	default:
		return request, invalidRequest{"Field \"params\" is neither array nor object"}
	}
	delete(body, "params")
	if id, ok := body["id"]; ok {
		switch id := id.(type) {
		case int64:
			request.id = numberID(id)
		case string:
			request.id = stringID(id)
		case nil:
			request.id = nullID{}
		default:
			return request, invalidRequest{"Field \"id\" is neither number nor string nor null"}
		}
	} else {
		request.id = notificationID{}
	}
	delete(body, "id")
	if len(body) > 0 {
		return request, invalidRequest{"Request contains extra fields"}
	}
	return request, nil
}
