package lib

import (
	"net/http"
	"testing"
)

type server struct {
	http.Handler
	maxConn uint
}

func (s *server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if s.maxConn < 0 {
		s.maxConn = 1
	}
	s.Handler.ServeHTTP(writer, request)
}

func TestHttpVersion(t *testing.T) {
	_ = http.NewServeMux()
}
