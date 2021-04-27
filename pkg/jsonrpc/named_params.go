package jsonrpc

type namedParams map[string]interface{}

func NewNamed(params map[string]interface{}) Params {
	return namedParams(params)
}

func (p namedParams) len() int {
	return len(p)
}
