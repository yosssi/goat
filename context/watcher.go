package context

import "fmt"

// Watcher represents a file watcher.
type Watcher struct {
	Extension string   `json:"extension"`
	Commands  []string `json:"commands"`
	CommandsC chan<- []string
}

// launch launches the watcher's process.
func (w *Watcher) Launch(commandsC chan<- []string) {
	w.CommandsC = commandsC
	fmt.Printf("%+v\n", w)
}
