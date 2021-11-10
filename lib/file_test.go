package lib

import (
	"os"
	"testing"
)

func TestWriteFile(t *testing.T) {
	if err := os.WriteFile("test.txt", []byte("Let's go"), 0666); err != nil {
		t.Fatal(err)
	}
}
