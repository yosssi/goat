package main

import (
	"log"

	"github.com/yosssi/goat/consts"
	"github.com/yosssi/goat/context"
)

// main executes main processes.
func main() {
	ctx, err := context.NewContext()
	if err != nil {
		log.Fatal(err)
	}

	commandsC := make(chan []string, consts.CommandsChannelBuffer)

	launchWatchers(ctx, commandsC)
}

func launchWatchers(ctx *context.Context, commandsC chan<- []string) {
	for _, watcher := range ctx.Config.Watchers {
		watcher.Launch(commandsC)
	}
}
