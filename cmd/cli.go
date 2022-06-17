package cmd

import (
	"github.com/go-ee/utils/cliu"
	"github.com/urfave/cli/v2"
)

type Cli struct {
	*cli.App
	*cliu.CommonFlags
}

func NewCli(common *cliu.CommonFlags, appName string, usage string) (ret *Cli) {
	app := cli.NewApp()
	ret = &Cli{
		App:         app,
		CommonFlags: common,
	}

	app.Name = appName
	app.Usage = usage
	app.Version = AppVersion()

	app.Flags = []cli.Flag{
		ret.Debug,
	}

	app.Commands = []*cli.Command{
		NewMigrateCmd().Command,
		cliu.NewMarkdownCmd(ret.App).Command,
	}
	return
}
