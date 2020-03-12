// Package runtime is the micro runtime
package runtime

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"sort"
	"strings"
	"syscall"
	"text/tabwriter"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/config/cmd"
	"github.com/micro/go-micro/v2/runtime"
	srvRuntime "github.com/micro/go-micro/v2/runtime/service"
	"github.com/micro/micro/v2/runtime/scheduler"
)

const (
	// RunUsage message for the run command
	RunUsage = "Required usage: micro run [service] [version] [--source github.com/micro/services]"
	// KillUsage message for the kill command
	KillUsage = "Require usage: micro kill [service] [version]"
	// UpdateUsage message for the update command
	UpdateUsage = "Require usage: micro update [service] [version]"
	// GetUsage message for micro get command
	GetUsage = "Require usage: micro ps [service] [version]"
	// ServicesUsage message for micro services command
	ServicesUsage = "Require usage: micro services"
	// CannotWatch message for the run command
	CannotWatch = "Cannot watch filesystem on this runtime"
)

var (
	// DefaultRetries which should be attempted when starting a service
	DefaultRetries = 3
)

func defaultEnv() []string {
	var env []string
	for _, evar := range os.Environ() {
		if strings.HasPrefix(evar, "MICRO_") {
			env = append(env, evar)
		}
	}

	return env
}

func runtimeFromContext(ctx *cli.Context) runtime.Runtime {
	if ctx.Bool("platform") {
		os.Setenv("MICRO_PROXY", "service")
		os.Setenv("MICRO_PROXY_ADDRESS", "proxy.micro.mu:443")
		return srvRuntime.NewRuntime()
	}

	return *cmd.DefaultCmd.Options().Runtime
}

func runService(ctx *cli.Context, srvOpts ...micro.Option) {
	// Init plugins
	for _, p := range Plugins() {
		p.Init(ctx)
	}

	// we need some args to run
	if ctx.Args().Len() == 0 {
		fmt.Println(RunUsage)
		return
	}

	// set and validate the name (arg 1)
	name := ctx.Args().Get(0)
	if strings.HasPrefix(name, ".") || strings.HasPrefix(name, "/") {
		fmt.Println(RunUsage)
		return
	}

	// set the version (arg 2, optional)
	version := "latest"
	if ctx.Args().Len() > 1 {
		version = ctx.Args().Get(1)
	}

	// load the runtime
	r := runtimeFromContext(ctx)

	source := ctx.String("source")
	exec := []string{"go", "run", filepath.Join(source, name)}

	// Determine the filepath
	fp := filepath.Join(os.Getenv("GOPATH"), "src", source, name)

	// Find the filepath or `go run` will pull from git by default
	if r.String() == "local" && os.Chdir(fp) == nil {
		exec = []string{"go", "run", "."}

		// watch the filesystem for changes
		sched := scheduler.New(name, version, fp)
		if err := r.Init(runtime.WithScheduler(sched)); err != nil {
			fmt.Printf("Could not start scheduler: %v", err)
			return
		}
	}

	// start the runtimes
	if err := r.Start(); err != nil {
		fmt.Printf("Could not start: %v", err)
		return
	}

	// add environment variable passed in via cli
	environment := defaultEnv()
	for _, evar := range ctx.StringSlice("env") {
		for _, e := range strings.Split(evar, ",") {
			if len(e) > 0 {
				environment = append(environment, strings.TrimSpace(e))
			}
		}
	}

	var retries = DefaultRetries
	if ctx.IsSet("retries") {
		retries = ctx.Int("retries")
	}

	// specify the options
	opts := []runtime.CreateOption{
		runtime.WithCommand(exec...),
		runtime.WithOutput(os.Stdout),
		runtime.WithEnv(environment),
		runtime.WithRetries(retries),
	}

	// run the service
	service := &runtime.Service{
		Name:     name,
		Source:   source,
		Version:  version,
		Metadata: make(map[string]string),
	}
	if err := r.Create(service, opts...); err != nil {
		fmt.Println(err)
		return
	}

	// if local	 then register signal handlers
	if r.String() == "local" {
		shutdown := make(chan os.Signal, 1)
		signal.Notify(shutdown, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

		// wait for shutdown
		<-shutdown

		// delete service from runtime
		if err := r.Delete(service); err != nil {
			fmt.Println(err)
			return
		}

		if err := r.Stop(); err != nil {
			fmt.Println(err)
			return
		}
	}
}

func killService(ctx *cli.Context, srvOpts ...micro.Option) {
	// we need some args to run
	if ctx.Args().Len() == 0 {
		fmt.Println(RunUsage)
		return
	}

	// set and validate the name (arg 1)
	name := ctx.Args().Get(0)
	if strings.HasPrefix(name, ".") || strings.HasPrefix(name, "/") {
		fmt.Println(RunUsage)
		return
	}

	// set the version (arg 2, optional)
	version := "latest"
	if ctx.Args().Len() > 1 {
		version = ctx.Args().Get(1)
	}

	service := &runtime.Service{
		Name:    name,
		Version: version,
	}

	if err := runtimeFromContext(ctx).Delete(service); err != nil {
		fmt.Println(err)
		return
	}
}

func updateService(ctx *cli.Context, srvOpts ...micro.Option) {
	// we need some args to run
	if ctx.Args().Len() == 0 {
		fmt.Println(RunUsage)
		return
	}

	// set and validate the name (arg 1)
	name := ctx.Args().Get(0)
	if strings.HasPrefix(name, ".") || strings.HasPrefix(name, "/") {
		fmt.Println(RunUsage)
		return
	}

	// set the version (arg 2, optional)
	version := "latest"
	if ctx.Args().Len() > 1 {
		version = ctx.Args().Get(1)
	}

	service := &runtime.Service{
		Name:    name,
		Version: version,
	}

	if err := runtimeFromContext(ctx).Update(service); err != nil {
		fmt.Println(err)
		return
	}
}

func getService(ctx *cli.Context, srvOpts ...micro.Option) {
	runType := ctx.Bool("runtime")

	// get and validate the name (arg 1, optional)
	name := ctx.Args().Get(0)
	if strings.HasPrefix(name, ".") || strings.HasPrefix(name, "/") {
		return
	}

	// get the version (arg 2, optional)
	version := "latest"
	if ctx.Args().Len() > 1 {
		version = ctx.Args().Get(1)
	}

	var list bool

	// zero args so list all
	if ctx.Args().Len() == 0 {
		list = true
	} else {
		// set name as first arg
		name = ctx.Args().Get(0)
		// set version as second arg
		if ctx.Args().Len() > 1 {
			version = ctx.Args().Get(1)
		}
	}

	var services []*runtime.Service
	var err error

	r := runtimeFromContext(ctx)

	// return a list of services
	switch list {
	case true:
		// return the runtiem services
		if runType {
			services, err = r.Read(runtime.ReadType("runtime"))
		} else {
			// list all running services
			services, err = r.List()
		}
	// return one service
	default:
		// check if service name was passed in
		if len(name) == 0 {
			fmt.Println(GetUsage)
			return
		}

		// get service with name and version
		opts := []runtime.ReadOption{
			runtime.ReadService(name),
			runtime.ReadVersion(version),
		}

		// return the runtime services
		if runType {
			opts = append(opts, runtime.ReadType("runtime"))
		}

		// read the service
		services, err = r.Read(opts...)
	}

	// check the error
	if err != nil {
		fmt.Println(err)
		return
	}

	// make sure we return UNKNOWN when empty string is supplied
	parse := func(m string) string {
		if len(m) == 0 {
			return "n/a"
		}
		return m
	}

	// don't do anything if there's no services
	if len(services) == 0 {
		fmt.Println("No services found")
		return
	}

	sort.Slice(services, func(i, j int) bool { return services[i].Name < services[j].Name })

	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
	fmt.Fprintln(writer, "NAME\tVERSION\tSOURCE\tSTATUS\tBUILD\tMETADATA")
	for _, service := range services {
		status := parse(service.Metadata["status"])
		if status == "error" {
			status = service.Metadata["error"]
		}

		fmt.Fprintf(writer, "%s\t%s\t%s\t%s\t%s\t%s\n",
			service.Name,
			parse(service.Version),
			parse(service.Source),
			status,
			parse(service.Metadata["build"]),
			fmt.Sprintf("owner=%s,group=%s", parse(service.Metadata["owner"]), parse(service.Metadata["group"])))
	}
	writer.Flush()
}
