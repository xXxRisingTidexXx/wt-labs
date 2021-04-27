package jsonrpc

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io"
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
	response := s.serveJSONRPC(request.Body)
	if response.HasError() {
		s.logError(response.error)
	}
	if response.shouldReturn {
		if err := json.NewEncoder(writer).Encode(response); err != nil {
			s.logError(wrappedError{err})
		}
	}
}

func (s *Server) logError(e Error) {
	s.logger.WithFields(log.Fields{"code": e.code(), "data": e.data()}).Error(e.message())
}

func (s *Server) serveJSONRPC(reader io.Reader) Response {
	var (
		body map[string]interface{}
		r    Request
	)
	if err := json.NewDecoder(reader).Decode(&body); err != nil {
		return Response{shouldReturn: true, error: parseError{err}, id: nullID{}}
	}
	if version, ok := body["jsonrpc"]; !ok || version != Version {
		return Response{
			shouldReturn: true,
			error:        invalidRequest{"Field \"jsonrpc\" is either absent or invalid"},
			id:           nullID{},
		}
	}
	delete(body, "jsonrpc")
	method, ok := body["method"]
	if !ok {
		return Response{
			shouldReturn: true,
			error:        invalidRequest{"Field \"method\" is absent"},
			id:           nullID{},
		}
	}
	switch method := method.(type) {
	case string:
		r.method = method
	default:
		return Response{
			shouldReturn: true,
			error:        invalidRequest{"Field \"method\" is not string"},
			id:           nullID{},
		}
	}
	delete(body, "method")
	params, ok := body["params"]
	if !ok {
		return Response{
			shouldReturn: true,
			error:        invalidRequest{"Field \"params\" is absent"},
			id:           nullID{},
		}
	}
	switch params := params.(type) {
	case []interface{}:
		r.params = positionalParams(params)
	case map[string]interface{}:
		r.params = namedParams(params)
	default:
		return Response{
			shouldReturn: true,
			error:        invalidRequest{"Field \"params\" is neither array nor object"},
			id:           nullID{},
		}
	}
	delete(body, "params")
	id, shouldReturn := body["id"]
	if !shouldReturn {
		switch id := id.(type) {
		case int64:
			r.id = numberID(id)
		case string:
			r.id = stringID(id)
		case nil:
			r.id = nullID{}
		default:
			return Response{
				shouldReturn: true,
				error:        invalidRequest{"Field \"id\" is neither number nor string nor null"},
				id:           nullID{},
			}
		}
	} else {
		r.id = notificationID{}
	}
	delete(body, "id")
	if len(body) > 0 {
		return Response{
			shouldReturn: true,
			error:        invalidRequest{"Request contains extra fields"},
			id:           nullID{},
		}
	}
	m, ok := s.methods[r.method]
	if !ok {
		return Response{shouldReturn: shouldReturn, error: methodNotFound{r.method}, id: r.id}
	}
	response := m.ServeJSONRPC(r)
	response.shouldReturn = shouldReturn
	return response
}
