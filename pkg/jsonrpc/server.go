package jsonrpc

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	methods map[string]Method
}

func NewServer(methods map[string]Method) *Server {
	return &Server{methods}
}

func (s *Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	r, e := ParseRequest(request.Body)
	if e != nil {
		s.writeResponse(Response{error: e, id: nullID{}}, false, writer)
	} else if method, ok := s.methods[r.method]; !ok {
		s.writeResponse(
			Response{error: methodNotFound{r.method}, id: r.id},
			r.IsNotification(),
			writer,
		)
	} else {
		response := method.ServeJSONRPC(r)
		response.id = r.id
		s.writeResponse(response, r.IsNotification(), writer)
	}
}

func (s *Server) writeResponse(
	response Response,
	isNotification bool,
	writer http.ResponseWriter,
) {
	if response.HasError() {
		logError(response.error)
	}
	if !isNotification {
		if err := json.NewEncoder(writer).Encode(response); err != nil {
			logError(wrappedError{err})
		}
	}
}

func logError(e Error) {
	log.WithFields(log.Fields{"code": e.code(), "data": e.data()}).Error(e.message())
}
