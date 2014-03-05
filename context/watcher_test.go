package context

import (
	"os"
	"testing"
)

func TestWatcherLaunch(t *testing.T) {
	wd := os.Getenv("GOPATH") + "/src/github.com/yosssi/goat/test/context/TestWatcherLaunch001"
	os.Chdir(wd)
	ctx, err := NewContext(500)
	if err != nil {
		t.Errorf("Error (%s) occurred.", err.Error())
	}
	if ctx.Config == nil || len(ctx.Config.Watchers) != 1 {
		t.Error("Context is invalid.")
	}
	watcher := ctx.Config.Watchers[0]
	if watcher == nil {
		t.Error("Watcher is invalid.")
	}
	watcher.Launch(ctx, make(chan Job))
}
