package math

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// go 1.20 之后推荐设置随机种子的方式发生了变化
func TestNewRandomSed(t *testing.T) {
	t.Log(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(4096))
	t.Log(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(4096))
	t.Log(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(4096))
	t.Log(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(4096))
}

func TestClient_GenerateProcessInstId(t *testing.T) {
	number := rand.New(rand.NewSource(time.Now().UnixNano()))
	id := fmt.Sprintf("%09d", number.Intn(99999999))
	t.Log(id)
}
