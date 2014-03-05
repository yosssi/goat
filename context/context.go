package context

import (
	"flag"
	"os"

	"github.com/yosssi/goat/consts"
)

// Context represents a context of a process.
type Context struct {
	Wd       string
	Config   *Config
	Interval int
}

// NewContext generates a Context and returns it.
func NewContext() (*Context, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	config, err := NewConfig()
	if err != nil {
		return nil, err
	}
	interval := flag.Int("i", consts.DefaultInterval, "An interval(ms) of a watchers' file check loop")
	flag.Parse()
	return &Context{Wd: wd, Config: config, Interval: *interval}, nil
}
