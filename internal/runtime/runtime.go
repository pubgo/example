package runtime

import (
	"os"
	"os/signal"
	"sort"
	"syscall"

	"github.com/pubgo/x/q"
	"github.com/pubgo/x/stack"
	"github.com/pubgo/xerror"
	"github.com/urfave/cli/v2"

	"github.com/pubgo/lava/config"
	"github.com/pubgo/lava/entry"
	"github.com/pubgo/lava/logger"
	"github.com/pubgo/lava/logz"
	"github.com/pubgo/lava/pkg/watcher"
	"github.com/pubgo/lava/plugin"
	"github.com/pubgo/lava/plugins/healthy"
	"github.com/pubgo/lava/plugins/syncx"
	"github.com/pubgo/lava/runenv"
	"github.com/pubgo/lava/vars"
	"github.com/pubgo/lava/version"
)

const name = "runtime"

var logs = logz.Component(name)
var app = &cli.App{
	Name:    runenv.Domain,
	Version: version.Version,
}

func handleSignal() {
	if runenv.CatchSigpipe {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGPIPE)
		syncx.GoSafe(func() {
			<-sigChan
			logs.Warn("Caught SIGPIPE (ignoring all future SIGPIPE)")
			signal.Ignore(syscall.SIGPIPE)
		})
	}

	if !runenv.Block {
		return
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGHUP)
	runenv.Signal = <-ch
	logs.Infof("signal [%s] trigger", runenv.Signal.String())
}

func start(ent entry.Runtime) {
	logs.StepAndThrow("before-start running", func() error {
		var beforeList []func()
		for _, p := range plugin.All() {
			beforeList = append(beforeList, p.BeforeStarts()...)
		}
		beforeList = append(beforeList, ent.Options().BeforeStarts...)
		for i := range beforeList {
			logs.Infof("running %s", stack.Func(beforeList[i]))
			xerror.PanicF(xerror.Try(beforeList[i]), stack.Func(beforeList[i]))
		}
		return nil
	})

	logs.StepAndThrow("server start", ent.Start)

	logs.StepAndThrow("after-start running", func() error {
		var afterList []func()
		for _, p := range plugin.All() {
			afterList = append(afterList, p.AfterStarts()...)
		}
		afterList = append(afterList, ent.Options().AfterStarts...)
		for i := range afterList {
			logs.Infof("running %s", stack.Func(afterList[i]))
			xerror.PanicF(xerror.Try(afterList[i]), stack.Func(afterList[i]))
		}
		return nil
	})
}

func stop(ent entry.Runtime) {
	logs.Step("before-stop running", func() error {
		var beforeList []func()
		for _, p := range plugin.All() {
			beforeList = append(beforeList, p.BeforeStops()...)
		}
		beforeList = append(beforeList, ent.Options().BeforeStops...)
		for i := range beforeList {
			logs.Infof("running %s", stack.Func(beforeList[i]))
			xerror.PanicF(xerror.Try(beforeList[i]), stack.Func(beforeList[i]))
		}
		return nil
	})

	logs.Step("server stop", ent.Stop)

	logs.Step("after-stop running", func() error {
		var afterList []func()
		for _, p := range plugin.All() {
			afterList = append(afterList, p.AfterStops()...)
		}
		afterList = append(afterList, ent.Options().AfterStops...)
		for i := range afterList {
			logs.Infof("running %s", stack.Func(afterList[i]))
			xerror.PanicF(xerror.Try(afterList[i]), stack.Func(afterList[i]))
		}
		return nil
	})
}

func Run(description string, entries ...entry.Entry) {
	defer xerror.RespExit()

	xerror.Assert(len(entries) == 0, "[entries] should not be zero")

	for _, ent := range entries {
		xerror.Assert(ent == nil, "[ent] should not be nil")

		_, ok := ent.(entry.Runtime)
		xerror.Assert(!ok, "[ent] not implement runtime, \n%s", q.Sq(ent))
	}

	app.Usage = description
	app.Description = description

	// ????????????flags
	app.Flags = append(app.Flags, config.DefaultFlags()...)

	// ????????????plugin
	for _, plg := range plugin.All() {
		app.Flags = append(app.Flags, plg.Flags()...)

		var cmd = plg.Commands()
		if cmd != nil {
			// ??????Command????????????
			xerror.Assert(app.Command(cmd.Name) != nil, "command(%s) already exists", cmd.Name)
			app.Commands = append(app.Commands, cmd)
		}

		// ??????????????????
		if plg.Health() != nil {
			healthy.Register(plg.UniqueName(), plg.Health())
		}

		// ??????vars
		xerror.Panic(plg.Vars(vars.Register))

		// ??????watcher
		if plg.Watch() != nil {
			watcher.Watch(plg.UniqueName(), plg.Watch())
		}
	}

	for i := range entries {
		ent := entries[i]
		entRT := ent.(entry.Runtime)
		cmd := entRT.Options().Command

		// ????????????Command????????????
		xerror.Assert(app.Command(cmd.Name) != nil, "command(%s) already exists", cmd.Name)

		cmd.Before = func(ctx *cli.Context) error {
			defer xerror.RespExit()
			// ??????????????????
			runenv.Project = entRT.Options().Name

			// ?????????????????????
			xerror.Panic(config.Init())

			// ???????????????
			logger.Init()

			// plugin?????????
			for _, plg := range plugin.All() {
				entRT.MiddlewareInter(plg.Middleware())
				logs.LogAndThrow("plugin init", plg.Init, logger.Name(plg.UniqueName()))
			}

			// entry?????????
			entRT.InitRT()

			// watcher?????????, ???????????????, ??????????????????????????????
			xerror.Panic(watcher.Init(entRT.Options().WatchProjects...))
			return nil
		}

		cmd.Action = func(ctx *cli.Context) error {
			defer xerror.RespExit()
			start(entRT)
			handleSignal()
			stop(entRT)
			return nil
		}

		app.Commands = append(app.Commands, cmd)
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	xerror.Panic(app.Run(os.Args))
}
