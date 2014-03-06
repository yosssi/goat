package context

import (
	"os"
	"testing"
)

func TestNewConfig(t *testing.T) {
	// When ioutil.ReadFile returns an error.
	expectedErrMsg := "open goat.json: no such file or directory"
	_, err := NewConfig()
	if err == nil || err.Error() != expectedErrMsg {
		t.Errorf("Error (%s) should be returned.", expectedErrMsg)
	}

	// When json.Unmarshal returns an error.
	expectedErrMsg = "unexpected end of JSON input"
	os.Chdir(os.Getenv("GOPATH") + "/src/github.com/yosssi/goat/test/context/TestNewConfig001")
	_, err = NewConfig()
	if err == nil || err.Error() != expectedErrMsg {
		t.Errorf("Error (%s) should be returned.", expectedErrMsg)
	}

	// When config is returned.
	os.Chdir(os.Getenv("GOPATH") + "/src/github.com/yosssi/goat/test/context/TestNewConfig002")
	config, err := NewConfig()
	if err != nil {
		t.Errorf("Error (%s) occurred.", err.Error())
	}
	if len(config.Watchers) != 1 || config.Watchers[0].Extension != "go" || len(config.Watchers[0].Tasks) != 1 || config.Watchers[0].Tasks[0].Command != "make rerun" || config.Watchers[0].Tasks[0].Nowait != true {
		t.Errorf("Config is invalid.")
	}
}
