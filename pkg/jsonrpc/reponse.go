package jsonrpc

import (
	"encoding/json"
	"io"
)

type Response struct {
	result interface{}
	error  Error
	id     interface{}
}

func NewResult(result interface{}) Response {
	return Response{result: result}
}

func NewError(e Error) Response {
	return Response{error: e}
}

func ParseResponse(reader io.Reader) (Response, Error) {
	var (
		body1    map[string]interface{}
		response Response
	)
	if err := json.NewDecoder(reader).Decode(&body1); err != nil {
		return response, clientError{err}
	}
	if version, ok := body1["jsonrpc"]; !ok || version != Version {
		return response, invalidResponse{"Field \"jsonrpc\" is either absent or invalid"}
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
		body2, ok := e.(map[string]interface{})
		if !ok {
			return response, invalidResponse{"Field \"error\" is not an object"}
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
		message, ok := body2["message"]
		if !ok {
			return response, invalidResponse{"Field \"message\" is absent"}
		}
		m, ok := message.(string)
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
	//id, ok := body1["id"]
	//if !ok {
	//	return response, invalidResponse{"Field \"id\" is absent"}
	//}
	//switch id := id.(type) {
	//case float64:
	//	response.id = numberID(id)
	//case string:
	//	response.id = stringID(id)
	//case nil:
	//	if !response.HasError() {
	//		return response, invalidResponse{"Field \"id\" can be null just in a case of error"}
	//	}
	//	response.id = nullID{}
	//default:
	//	return response, invalidResponse{"Field \"id\" is neither number nor string nor null"}
	//}
	delete(body1, "id")
	if len(body1) > 0 {
		return response, invalidResponse{"Response contains extra fields"}
	}
	return response, nil
}

func (r Response) HasError() bool {
	return r.error != nil
}

func (r Response) MarshalJSON() ([]byte, error) {
	body := map[string]interface{}{"jsonrpc": Version}
	if r.HasError() {
		e := map[string]interface{}{"code": r.error.code(), "message": r.error.message()}
		if r.error.data() != nil {
			e["data"] = r.error.data()
		}
		body["error"] = e
	} else {
		body["result"] = r.result
	}
	//if id, ok := r.id.toValue(); ok {
	//	body["id"] = id
	//}
	return json.Marshal(body)
}
