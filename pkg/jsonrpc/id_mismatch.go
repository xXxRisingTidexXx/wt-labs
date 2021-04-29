package jsonrpc

type idMismatch struct {
	expected ID
	actual   ID
}

func (m idMismatch) code() int {
	return -32103
}

func (m idMismatch) message() string {
	return "ID mismatch"
}

func (m idMismatch) data() interface{} {
	return map[string]ID{"expected": m.expected, "actual": m.actual}
}
