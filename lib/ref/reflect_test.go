package ref

import (
	"fmt"
	"reflect"
	"testing"
)

type Data struct {
	ID string `json:"id,omitempty" yaml:"id"`
}

func (d *Data) Do() {
	fmt.Println(d.ID)
}

func TestReflect(t *testing.T) {
	var d = &Data{ID: "x"}
	v := reflect.ValueOf(d)
	t.Log(v.String())
}
