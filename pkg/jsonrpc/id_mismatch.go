package jsonrpc

type idMismatch struct {
	expected interface{}
	actual   interface{}
}

func (m idMismatch) code() int {
	return -32103
}

func (m idMismatch) message() string {
	return "ID mismatch"
}

func (m idMismatch) data() interface{} {
	return map[string]interface{}{"expected": m.expected, "actual": m.actual}
}
