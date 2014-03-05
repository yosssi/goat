package context

import "os"

// Context represents a context of a process.
type Context struct {
	Wd       string
	Config   *Config
	Interval int
}

// NewContext generates a Context and returns it.
func NewContext(interval int) (*Context, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	config, err := NewConfig()
	if err != nil {
		return nil, err
	}
	return &Context{Wd: wd, Config: config, Interval: interval}, nil
}
