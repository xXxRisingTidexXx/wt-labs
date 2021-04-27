package jsonrpc

import (
	"encoding/json"
)

type Response struct {
	shouldReturn bool
	result       interface{}
	error        Error
	id           ID
}

func (r Response) HasError() bool {
	return r.error != nil
}

func (r Response) MarshalJSON() ([]byte, error) {
	id, _ := r.id.toValue()
	body := map[string]interface{}{"jsonrpc": Version, "id": id}
	if r.HasError() {
		e := map[string]interface{}{"code": r.error.code(), "message": r.error.message()}
		if r.error.data() != nil {
			e["data"] = r.error.data()
		}
		body["error"] = e
	} else {
		body["result"] = r.result
	}
	return json.Marshal(body)
}
