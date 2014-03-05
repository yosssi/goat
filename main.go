package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strings"

	"github.com/yosssi/goat/consts"
	"github.com/yosssi/goat/context"
)

// main executes main processes.
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	ctx, err := context.NewContext()
	if err != nil {
		log.Fatal(err)
	}
	jobsC := make(chan context.Job, consts.JobsChannelBuffer)

	launchWatchers(ctx, jobsC)

	handleJobs(jobsC)
}

// launchWatchers launches watchers.
func launchWatchers(ctx *context.Context, jobsC chan<- context.Job) {
	for _, watcher := range ctx.Config.Watchers {
		go watcher.Launch(ctx, jobsC)
	}
}

// handleJobs handle jobs.
func handleJobs(jobsC <-chan context.Job) {
	for job := range jobsC {
		watcher := job.Watcher
		watcher.Printf("%s", job.Message)
		for _, command := range watcher.Commands {
			tokens := strings.Split(command, " ")
			name := tokens[0]
			var cmdArg []string
			if len(tokens) > 1 {
				cmdArg = tokens[1:]
			}
			watcher.Printf("execute: %s", command)
			bytes, err := exec.Command(name, cmdArg...).Output()
			if len(bytes) > 0 {
				fmt.Print(string(bytes))
			}
			if err != nil {
				watcher.Printf("An error occurred: %s", err.Error())
			}
		}
	}
}
