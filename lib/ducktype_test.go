package lib

import "testing"

type slice[T string | int64] []T

type Map[K string, V string | float64] map[K]V

func TestDuckType(t *testing.T) {
	var (
		ss slice[string] = []string{
			"1", "2",
		}
		mm Map[string, float64] = map[string]float64{
			"s": 3.15,
		}
	)
	for _, s := range ss {
		t.Log(s)
	}
	for k, v := range mm {
		t.Log(k, v)
	}
}
