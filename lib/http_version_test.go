package lib

import (
	"net/http"
	"net/url"
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

func TestUulJoin(t *testing.T) {
	path, err := url.JoinPath("https://polaris.woa.com/#/services/info/instance/", "Development", "txstore_global_api-server")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(path)
}
