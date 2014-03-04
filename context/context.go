package context

// Context represents a context of a process.
type Context struct {
	Config *Config
}

// NewContext generates a Context and returns it.
func NewContext() (*Context, error) {
	config, err := NewConfig()
	if err != nil {
		return nil, err
	}
	return &Context{Config: config}, nil
}
