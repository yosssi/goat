package context

// A Task represents a task.
type Task struct {
	Command string `json:"command" yaml:"command"`
	Nowait  bool   `json:"nowait" yaml:"nowait"`
}
