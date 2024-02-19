package main

import (
	"testing"
	"time"

	"github.com/gogf/gf/v2/os/gtime"
)

func TestGTimeLayout(t *testing.T) {
	now := gtime.Now().Local().Layout(time.DateOnly)
	t.Log(now)
	var threeDaysAgo = time.Now().AddDate(0, 0, -3).Format(time.RFC3339)
	t.Log(threeDaysAgo)
	today := time.Now().Local().Format(time.DateOnly)
	t.Log(now == today)
}

func TestTimeLocal(t *testing.T) {
	t.Log(gtime.Now().Local().String())
	s := gtime.New("2024-01-09").String()
	t.Log(s)
}
