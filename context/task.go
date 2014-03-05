package context

// A Task represents a task.
type Task struct {
	Command string `json:"command"`
	Nowait  bool   `json:"nowait"`
}
