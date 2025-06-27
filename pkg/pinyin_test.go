package pkg

import (
	"strings"
	"testing"

	"github.com/mozillazg/go-pinyin"
)

func TestPinyin(t *testing.T) {
	convert := pinyin.LazyConvert("我是中国人", nil)
	t.Log(strings.ToUpper(strings.Join(convert, "_")))
}
