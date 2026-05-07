package lib

import (
	"os"
	"path"
	"testing"
)

func TestFmt(t *testing.T) {
	var ss = []string{"x", "y", "z"}

	for _, i := range ss {
		switch i {
		case "x":
			t.Log(i)
		case "y":
			t.Log("Y")
		default:
			t.Log("-")
		}
	}
}

func TestPWD(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(path.Join(dir, "xxx.txt"))
}
