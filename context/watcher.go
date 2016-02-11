package context

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

// Watcher represents a file watcher.
type Watcher struct {
	Extension string   `json:"extension" yaml:"extension"`
	Directory string   `json:"directory" yaml:"directory"`
	Excludes  []string `json:"excludes" yaml:"excludes"`
	Tasks     []*Task  `json:"tasks" yaml:"tasks"`
	JobsC     chan<- Job
	Targets   map[string]map[string]os.FileInfo
}

// launch launches the watcher's process.
func (w *Watcher) Launch(ctx *Context, jobsC chan<- Job) {
	w.JobsC = jobsC
	w.Targets = make(map[string]map[string]os.FileInfo)
	watchDir := ctx.Wd
	if w.Directory != "" {
		watchDir = watchDir+"/"+w.Directory
	}
	w.readDir(watchDir, true)
	w.Printf("%s", "Watching...")
	for {
		time.Sleep(time.Duration(ctx.Interval) * time.Millisecond)
		w.readDir(watchDir, false)
	}
}

// readDir reads the directory named by dirname.
func (w *Watcher) readDir(dirname string, init bool) error {
	fileInfos, err := ioutil.ReadDir(dirname)
	if err != nil {
		return err
	}
	for _, fileInfo := range fileInfos {
		name := fileInfo.Name()
		switch {
		case strings.HasPrefix(name, "."):
		case fileInfo.IsDir():
			if err := w.readDir(dirname+"/"+name, init); err != nil {
				return err
			}
		case w.exclude(name):
		case strings.HasSuffix(name, "."+w.Extension):
			_, prs := w.Targets[dirname]
			if !prs {
				w.Targets[dirname] = make(map[string]os.FileInfo)
			}
			if init {
				w.Targets[dirname][name] = fileInfo
			} else {
				preservedFileInfo, prs := w.Targets[dirname][name]
				if !prs || preservedFileInfo.ModTime() != fileInfo.ModTime() {
					w.Targets[dirname][name] = fileInfo
					var action string
					if !prs {
						action = "created"
					} else {
						action = "updated"
					}
					w.sendJob(dirname, name, action)
				}
			}
		}
	}
	if !init {
		preservedFileInfos, prs := w.Targets[dirname]
		if prs {
			for name, _ := range preservedFileInfos {
				exist := false
				for _, fileInfo := range fileInfos {
					if name == fileInfo.Name() {
						exist = true
						break
					}
				}
				if !exist {
					delete(w.Targets[dirname], name)
					w.sendJob(dirname, name, "deleted")
				}
			}
		}
	}
	return nil
}

// sendJob sends a job to the channel.
func (w *Watcher) sendJob(dirname, name, action string) {
	message := fmt.Sprintf("%s was %s.", dirname+"/"+name, action)
	w.JobsC <- Job{Watcher: w, Message: message}
}

// Printf calls log.Printf.
func (w *Watcher) Printf(format string, v ...interface{}) {
	watchDir := "project root"
	if w.Directory != "" {
		watchDir = w.Directory
	}
	log.Printf("[Watcher for "+w.Extension+" files under "+watchDir+"] "+format, v...)
}

// exclude returns true if the file should be not checked.
func (w *Watcher) exclude(filename string) bool {
	for _, excludeFilename := range w.Excludes {
		if filename == excludeFilename {
			return true
		}
	}
	return false
}
