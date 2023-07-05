package valute

import (
	"github.com/Knetic/govaluate"
	"testing"
)

func TestValute(t *testing.T) {
	expression, err := govaluate.NewEvaluableExpression("1/c")
	if err != nil {
		t.Fatal(err)
	}
	evaluate, err := expression.Evaluate(map[string]interface{}{
		"c": "1",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(evaluate)
}

func TestLen(t *testing.T) {
	var x = []interface{}{
		"x",
		1,
		struct {
			Name string
		}{"j"},
	}
	t.Log(x)
}
