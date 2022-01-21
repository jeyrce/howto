package pkg

import (
	"github.com/fsnotify/fsnotify"
	"testing"
	"time"
)

func TestFsWatch(t *testing.T) {
	file := "/tmp/fs.txt"
	t.Log("test file watch")
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		t.Fatal(err)
	}
	watcher.Add(file)
	defer watcher.Close()
	if err != nil {
		t.Fatal(err)
	}
	reset := false
	for {
		if reset {
			if err := watcher.Add(file); err != nil {
				t.Errorf("failed to add file")
				time.Sleep(5 * time.Second)
				continue
			}
		}
		select {
		case e := <-watcher.Events:
			reset = false
			switch {
			case e.Op&fsnotify.Create == fsnotify.Create:
				t.Log("new file")
			case e.Op&fsnotify.Remove == fsnotify.Remove || e.Op&fsnotify.Rename == fsnotify.Rename:
				t.Log("restart")
				reset = true
			case e.Op&fsnotify.Write == fsnotify.Write:
				t.Log("file changed")
			default:
				t.Log(e.Op.String())
			}
		}
	}
}
