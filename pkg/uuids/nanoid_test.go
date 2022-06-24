package uuids

import (
	"testing"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

// 自定义字符以及长度url友好
func TestNanoid(t *testing.T) {
	alphabet := "1234567890"
	letters := "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < 999; i++ {
		t.Log(gonanoid.MustGenerate(alphabet, 9))
		t.Log(gonanoid.MustGenerate(letters, 9))
	}
}

// https://ioseek.cn/p/852248652
// https://ioseek.cn/p/hpsxchtei
