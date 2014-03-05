package context

import (
	"os"
	"testing"
)

func TestNewContext(t *testing.T) {
	// When NewConfig returns an error.
	wd := os.Getenv("GOPATH") + "/src/github.com/yosssi/goat/test/context/TestNewContext001"
	os.Chdir(wd)
	expectedErrMsg := "open goat.json: no such file or directory"
	_, err := NewContext(500)
	if err == nil || err.Error() != expectedErrMsg {
		t.Errorf("Error (%s) should be returned. %s %s", expectedErrMsg, err.Error(), wd)
	}

	// When NewConfig returns a context.
	wd = os.Getenv("GOPATH") + "/src/github.com/yosssi/goat/test/context/TestNewContext002"
	os.Chdir(wd)
	ctx, err := NewContext(500)
	if err != nil {
		t.Errorf("Error (%s) occurred.", err.Error())
	}
	if ctx == nil || ctx.Wd != wd || ctx.Config == nil || len(ctx.Config.Watchers) != 1 || ctx.Config.Watchers[0].Extension != "go" || len(ctx.Config.Watchers[0].Tasks) != 1 || ctx.Config.Watchers[0].Tasks[0].Command != "make rerun" || ctx.Config.Watchers[0].Tasks[0].Nowait != true || ctx.Interval != 500 {
		t.Errorf("Context is invalid.")
	}
}
