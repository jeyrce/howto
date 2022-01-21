package lib

import (
	"strconv"
	"testing"
	"time"
)

func TestTimeToTimeStamp(t *testing.T) {
	var n = time.Now()
	t.Log(n.UnixNano())
	// 1636702193748667173
	// 1636690156498253038
	t.Log(n.UnixNano() / 1e6)
	// 1636702380177
	// 1636698230100
	t.Log(strconv.FormatInt(n.UnixNano()/1e6, 10))
}

func TestTimeFromString(t *testing.T) {
	var s = "[65s]"
	duration, err := time.ParseDuration(s[1 : len(s)-1])
	if err != nil {
		t.Fatal(err)
	}
	t.Log(duration)
}

func TestTimeStamp(t *testing.T) {
	ts := time.Now().Unix()
	t.Log(ts)
	rt := time.Unix(ts, 0).AddDate(0, 3, 0)
	t.Log(rt.Format("2006-01-02 15:04:05"))
}
