package pkg

import (
	"os"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestYml(t *testing.T) {
	file, err := os.ReadFile("test.yml")
	if err != nil {
		t.Error(err)
	}
	var out interface{}
	err = yaml.Unmarshal(file, &out)
	if err != nil {
		t.Error(err)
	}
	t.Log(out)
}
