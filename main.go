package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/yosssi/goat/consts"
	"github.com/yosssi/goat/context"
)

// main executes main processes.
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	version := flag.Bool("v", false, "Show Goat version")
	interval := flag.Int("i", consts.DefaultInterval, "An interval(ms) of a watchers' file check loop")
	flag.Parse()

	if *version {
		fmt.Printf("Goat %s\n", consts.Version)
		os.Exit(0)
	}

	ctx, err := context.NewContext(*interval)
	if err != nil {
		log.Fatal(err)
	}

	initTasks := ctx.Config.InitTasks
	if initTasks != nil && len(initTasks) > 0 {
		executeTasks(initTasks, nil)
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
		executeTasks(watcher.Tasks, watcher)
	}
}

// executeTasks executes tasks.
func executeTasks(tasks []*context.Task, watcher *context.Watcher) {
	for _, task := range tasks {
		command := task.Command
		tokens := strings.Split(command, " ")
		name := tokens[0]
		var cmdArg []string
		if len(tokens) > 1 {
			cmdArg = tokens[1:]
		}
		cmd := exec.Command(name, cmdArg...)
		if task.Nowait {
			printf(watcher, "execute(nowait): %s", command)
			if err := cmd.Start(); err != nil {
				printf(watcher, "An error occurred: %s", err.Error())
			} else {
				printf(watcher, "end(nowait): %s", command)
			}
		} else {
			printf(watcher, "execute: %s", command)
			bytes, err := cmd.Output()
			if err != nil {
				printf(watcher, "An error occurred: %s", err.Error())
			} else {
				fmt.Print(string(bytes))
				printf(watcher, "end: %s", command)
			}
		}
	}
}

func printf(watcher *context.Watcher, format string, v ...interface{}) {
	if watcher != nil {
		watcher.Printf(format, v)
	} else {
		log.Printf(format, v)
	}
}
