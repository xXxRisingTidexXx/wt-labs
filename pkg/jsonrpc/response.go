package jsonrpc

import (
	"encoding/json"
	"io"
)

type Response struct {
	result []byte
	error  Error
	id     interface{}
}

func WithResult(result interface{}) Response {
	bytes, err := json.Marshal(result)
	if err != nil {
		return Response{error: stringifyError{err}}
	}
	return Response{result: bytes}
}

func WithError(e Error) Response {
	return Response{error: e}
}

func ParseResponse(reader io.Reader) (Response, Error) {
	var (
		body1    map[string]json.RawMessage
		response Response
		version  string
	)
	if err := json.NewDecoder(reader).Decode(&body1); err != nil {
		return response, clientError{err}
	}
	message, ok := body1["jsonrpc"]
	if !ok {
		return response, invalidResponse{"Field \"jsonrpc\" is absent"}
	}
	if err := json.Unmarshal(message, &version); err != nil {
		return response, clientError{err}
	}
	if version != Version {
		return response, invalidResponse{"Field \"jsonrpc\" is invalid"}
	}
	delete(body1, "jsonrpc")
	result, ok1 := body1["result"]
	e, ok2 := body1["error"]
	if !ok1 && !ok2 {
		return response, invalidResponse{"Either \"result\" field or \"error\" one must exist"}
	}
	if ok1 && ok2 {
		return response, invalidResponse{"Fields \"result\" and \"error\" are mutually exclusive"}
	}
	if ok1 {
		response.result = result
		delete(body1, "result")
	} else {
		var body2 map[string]interface{}
		if err := json.Unmarshal(e, &body2); err != nil {
			return response, clientError{err}
		}
		code, ok := body2["code"]
		if !ok {
			return response, invalidResponse{"Field \"code\" is absent"}
		}
		c, ok := code.(float64)
		if !ok {
			return response, invalidResponse{"Field \"code\" is not a number"}
		}
		delete(body2, "code")
		text, ok := body2["message"]
		if !ok {
			return response, invalidResponse{"Field \"message\" is absent"}
		}
		m, ok := text.(string)
		if !ok {
			return response, invalidResponse{"Field \"message\" is not a number"}
		}
		delete(body2, "message")
		response.error = structuredError{int(c), m, body2["data"]}
		delete(body2, "data")
		if len(body2) > 0 {
			return response, invalidResponse{"Error contains extra fields"}
		}
		delete(body1, "error")
	}
	message, ok = body1["id"]
	if !ok {
		return response, invalidResponse{"Field \"id\" is absent"}
	}
	if err := json.Unmarshal(message, &response.id); err != nil {
		return response, clientError{err}
	}
	switch id := response.id.(type) {
	case float64:
		response.id = int64(id)
	case string:
	case nil:
		if !response.HasError() {
			return response, invalidResponse{"Field \"id\" can be null just in a case of error"}
		}
	default:
		return response, invalidResponse{"Field \"id\" is neither number nor string nor null"}
	}
	delete(body1, "id")
	if len(body1) > 0 {
		return response, invalidResponse{"Response contains extra fields"}
	}
	return response, nil
}

func (r Response) HasError() bool {
	return r.error != nil
}

func (r Response) UnmarshalResult(value interface{}) Error {
	if err := json.Unmarshal(r.result, value); err != nil {
		return invalidResult{err}
	}
	return nil
}

func (r Response) MarshalJSON() ([]byte, error) {
	body := map[string]interface{}{"jsonrpc": Version, "id": r.id}
	if r.HasError() {
		e := map[string]interface{}{"code": r.error.code(), "message": r.error.message()}
		if r.error.data() != nil {
			e["data"] = r.error.data()
		}
		body["error"] = e
	} else {
		body["result"] = json.RawMessage(r.result)
	}
	return json.Marshal(body)
}
