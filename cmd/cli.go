package cmd

import (
	"github.com/go-ee/jfrog/core"
	"github.com/go-ee/utils/cliu"
	"github.com/go-ee/utils/exec"
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
		NewCloneRepoCmd().Command,
		NewCloneReposCmd().Command,
		NewMigrateRepoCmd().Command,
		NewMigrateReposCmd().Command,
		cliu.NewMarkdownCmd(ret.App).Command,
	}
	return
}

type BaseCmd struct {
	*cliu.BaseCommand
	Source     *ServerFlagLabels
	Target     *ServerFlagLabels
	DryRunFlag *DryRunFlag
}

func NewBaseCmd() *BaseCmd {
	return &BaseCmd{
		BaseCommand: &cliu.BaseCommand{},
		Source:      NewServerDef("source"),
		Target:      NewServerDef("target"),
		DryRunFlag:  NewDryRunFlag()}
}

func buildExecutor(dryRunFlag *DryRunFlag) (ret exec.Executor) {
	if dryRunFlag.CurrentValue {
		ret = &exec.SkipExecutor{}
	} else {
		ret = &exec.LogExecutor{}
	}
	return
}

func buildArtifactoryManager(server *ServerFlagLabels, executor exec.Executor) *core.ArtifactoryManager {
	return &core.ArtifactoryManager{
		Label:    server.BuildLabel(),
		Url:      server.Url.CurrentValue,
		User:     server.User.CurrentValue,
		Password: server.Password.CurrentValue,

		Executor: executor,
	}
}
