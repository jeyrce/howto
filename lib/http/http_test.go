package http

import (
	"crypto/tls"
	"io"
	"net/http"
	"testing"
	"time"
)

func TestHTTPPing(t *testing.T) {
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: false,
			},
		},
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       time.Second * 2,
	}
	request, err := http.NewRequest(http.MethodGet, "lujianxin.com:80", nil)
	if err != nil {
		t.Fatal(err)
	}
	response, err := client.Do(request)
	if err != nil {
		t.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)
	t.Log("OK")
}
