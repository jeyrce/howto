package lib

import (
	"path"
	"testing"
)

func TestU(t *testing.T) {
	var ps = []string{
		"",
		"/",
		"/ioseek",
		"/ioseek/",
	}
	var urls = []string{
		"/api/v1",
		"/version",
		"metrics",
		"article/1",
	}
	for _, url := range urls {
		for _, p := range ps {
			t.Log(path.Join(p, url))
		}
	}
}

func TestPathJoin(t *testing.T) {
	t.Log(path.Join("/", "", "api/v2"))
	t.Log(path.Join("/", "app", "api/v3"))
}
