package jsonrpc

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	methods map[string]Method
	logger  log.FieldLogger
}

func NewServer(methods map[string]Method) *Server {
	return &Server{methods, log.New()}
}

func (s *Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	r, e := ParseRequest(request.Body)
	if e != nil {
		s.writeResponse(Response{error: e, id: nullID{}}, true, writer)
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

func (s *Server) writeResponse(response Response, shouldReply bool, writer http.ResponseWriter) {
	if response.HasError() {
		s.logError(response.error)
	}
	if shouldReply {
		if err := json.NewEncoder(writer).Encode(response); err != nil {
			s.logError(wrappedError{err})
		}
	}
}

func (s *Server) logError(e Error) {
	s.logger.WithFields(log.Fields{"code": e.code(), "data": e.data()}).Error(e.message())
}
