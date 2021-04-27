package jsonrpc

type numberID int64

func NewInt(id int) ID {
	return numberID(id)
}

func NewInt64(id int64) ID {
	return numberID(id)
}

func (id numberID) toValue() (interface{}, bool) {
	return id, true
}

func (id numberID) toNumber() (int64, bool) {
	return int64(id), true
}

func (id numberID) toString() (string, bool) {
	return "", false
}

func (id numberID) isNull() bool {
	return false
}
