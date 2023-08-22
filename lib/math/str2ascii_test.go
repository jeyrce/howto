package math

import (
	"strconv"
	"testing"
)

func TestStringToAscii(t *testing.T) {
	var data = []rune{
		'1',
		'丨',
		'|',
		'x',
	}
	for _, s := range data {
		t.Log(strconv.FormatInt(int64(s), 16))
	}
}
