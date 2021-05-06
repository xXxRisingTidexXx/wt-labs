package jsonrpc

import (
	"encoding/json"
	"github.com/lithammer/shortuuid"
	"io"
)

type Request struct {
	method string
	params []byte
	id     interface{}
	hasID  bool
}

func NewRequest(method string, params interface{}) (Request, Error) {
	return newRequest(method, shortuuid.New(), true, params)
}

func newRequest(
	method string,
	id interface{},
	hasID bool,
	params interface{},
) (Request, Error) {
	request := Request{method: method, id: id, hasID: hasID}
	if params != nil {
		var err error
		if request.params, err = json.Marshal(params); err != nil {
			return request, parseError{err}
		}
	}
	return request, nil
}

func NewNotification(method string, params interface{}) (Request, Error) {
	return newRequest(method, nil, false, params)
}

func ParseRequest(reader io.Reader) (Request, Error) {
	var (
		body    map[string]json.RawMessage
		request Request
		version string
	)
	if err := json.NewDecoder(reader).Decode(&body); err != nil {
		return request, parseError{err}
	}
	message, ok := body["jsonrpc"]
	if !ok {
		return request, invalidRequest{"Field \"jsonrpc\" is absent"}
	}
	if err := json.Unmarshal(message, &version); err != nil {
		return request, parseError{err}
	}
	if version != Version {
		return request, invalidRequest{"Field \"jsonrpc\" is invalid"}
	}
	delete(body, "jsonrpc")
	message, ok = body["method"]
	if !ok {
		return request, invalidRequest{"Field \"method\" is absent"}
	}
	if err := json.Unmarshal(message, &request.method); err != nil {
		return request, parseError{err}
	}
	delete(body, "method")
	if request.params, ok = body["params"]; ok {
		if len(request.params) < 2 || request.params[0] != '[' && request.params[0] != '{' {
			return request, invalidRequest{"Field \"params\" is neither array nor object"}
		}
		delete(body, "params")
	}
	if message, request.hasID = body["id"]; request.hasID {
		if err := json.Unmarshal(message, &request.id); err != nil {
			return request, parseError{err}
		}
		switch id := request.id.(type) {
		case float64:
			request.id = int64(id)
		case string, nil:
		default:
			return request, invalidRequest{"Field \"id\" is neither number nor string nor null"}
		}
		delete(body, "id")
	}
	if len(body) > 0 {
		return request, invalidRequest{"Request contains extra fields"}
	}
	return request, nil
}

func (r Request) IsNotification() bool {
	return !r.hasID
}

func (r Request) UnmarshalParams(value interface{}) Error {
	if err := json.Unmarshal(r.params, value); err != nil {
		return parseError{err}
	}
	return nil
}

func (r Request) MarshalJSON() ([]byte, error) {
	body := map[string]interface{}{"jsonrpc": Version, "method": r.method}
	if r.params != nil {
		body["params"] = json.RawMessage(r.params)
	}
	if r.hasID {
		body["id"] = r.id
	}
	return json.Marshal(body)
}
