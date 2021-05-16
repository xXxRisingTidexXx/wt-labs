package jsonrpc

import (
	log "github.com/sirupsen/logrus"
)

func LogError(e Error) {
	log.WithFields(log.Fields{"code": e.code(), "data": e.data()}).Error(e.message())
}
