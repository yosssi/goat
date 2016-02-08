package context

import (
	"encoding/json"
	"io/ioutil"

	"github.com/yosssi/goat/consts"
	"gopkg.in/yaml.v2"
)

// Config represents a configuration of a process.
type Config struct {
	InitTasks []*Task    `json:"init_tasks" yaml:"init_tasks"`
	Watchers  []*Watcher `json:"watchers" yaml:"watchers"`
}

// NewConfig parses a JSON file, generates a Config and returns it.
func NewConfig() (*Config, error) {
	bytes, err := ioutil.ReadFile(consts.JSONConfigFile)
	if err != nil {
		bytes, err = ioutil.ReadFile(consts.YAMLConfigFile)
		if err != nil {
			return nil, err
		}

		config := &Config{}
		if err := yaml.Unmarshal(bytes, config); err != nil {
			return nil, err
		}
		return config, nil
	}

	config := &Config{}
	if err := json.Unmarshal(bytes, config); err != nil {
		return nil, err
	}
	return config, nil
}
