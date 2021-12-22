package main

import (
	"github.com/pubgo/example/services/entry/cli_entry"
	"github.com/pubgo/example/services/entry/gid"
	"github.com/pubgo/example/services/entry/gin_entry"
	"github.com/pubgo/example/services/entry/grpc_entry"
	"github.com/pubgo/example/services/entry/task_entry"
	"github.com/pubgo/example/services/entry/version_entry"
	"github.com/pubgo/lava"
)

func main() {
	lava.Run(
		"gid service",
		gid.GetEntry(),
		gin_entry.GetEntry(),
		cli_entry.GetEntry(),
		grpc_entry.GetEntry(),
		version_entry.GetEntry(),
		task_entry.GetEntry(),
	)
}
