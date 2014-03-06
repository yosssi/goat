package context

import (
	"io/ioutil"

	"github.com/yosssi/goat/consts"
	"launchpad.net/goyaml"
)

// Config represents a configuration of a process.
type Config struct {
	InitTasks []*Task    `json:"init_tasks"`
	Watchers  []*Watcher `json:"watchers"`
}

// NewConfig parses a JSON file, generates a Config and returns it.
func NewConfig() (*Config, error) {
	bytes, err := ioutil.ReadFile(consts.ConfigFile)
	if err != nil {
		return nil, err
	}
	config := &Config{}
	if err := goyaml.Unmarshal(bytes, config); err != nil {
		return nil, err
	}
	return config, nil
}
