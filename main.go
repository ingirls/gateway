package main

import (
	"github.com/micro/go-micro/v2"

	rpc "github.com/micro/go-plugins/micro/disable_rpc/v2"
	"github.com/micro/go-plugins/micro/metrics/v2"
	"github.com/micro/micro/v2/cmd"
	"github.com/micro/micro/v2/plugin"
)

func init() {
	plugin.Register(metrics.NewPlugin())
	plugin.Register(rpc.NewPlugin())
}

func main() {
	cmd.Init(
		micro.Name("go.micro.api"),
	)
}
