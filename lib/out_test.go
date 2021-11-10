package lib

import (
	"fmt"
	"os"
	"path"
	"testing"
)

func TestFmt(t *testing.T) {
	var ss = []string{"x", "y", "z"}

	for _, i := range ss {
		switch i {
		case "x":
			fmt.Printf(i)
		case "y":
			fmt.Printf("Y")
		default:
			fmt.Printf("-")
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
