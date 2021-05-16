package jsonrpc

import (
	"encoding/json"
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
		s.writeResponse(WithError(e), false, writer)
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
	if response.error != nil {
		LogError(response.error)
	}
	if !isNotification {
		if err := json.NewEncoder(writer).Encode(response); err != nil {
			LogError(serverError{err})
		}
	}
}
