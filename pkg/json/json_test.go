package json

import (
	"fmt"
	"testing"

	"github.com/gogf/gf/encoding/gjson"
)

func TestJsonList(t *testing.T) {
	var data = `[{"name":"jeyrce", "mark":"txt"},{"name":"x", "mark":"cc"},{"name":"yyy","mark":"zzz"}]`
	items := gjson.New(data).Array()
	for _, item := range items {
		fmt.Println(item)
	}
}
