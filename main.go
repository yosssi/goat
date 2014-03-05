package main

import (
	"flag"
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

	interval := flag.Int("i", consts.DefaultInterval, "An interval(ms) of a watchers' file check loop")
	flag.Parse()

	ctx, err := context.NewContext(*interval)
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
		for _, task := range watcher.Tasks {
			command := task.Command
			tokens := strings.Split(command, " ")
			name := tokens[0]
			var cmdArg []string
			if len(tokens) > 1 {
				cmdArg = tokens[1:]
			}
			cmd := exec.Command(name, cmdArg...)
			if task.Nowait {
				watcher.Printf("execute(nowait): %s", command)
				if err := cmd.Start(); err != nil {
					watcher.Printf("An error occurred: %s", err.Error())
				} else {
					watcher.Printf("end(nowait): %s", command)
				}
			} else {
				watcher.Printf("execute: %s", command)
				bytes, err := cmd.Output()
				if err != nil {
					watcher.Printf("An error occurred: %s", err.Error())
				} else {
					fmt.Print(string(bytes))
					watcher.Printf("end: %s", command)
				}
			}
		}
	}
}
