package cmd

import (
	"github.com/go-ee/jfrog/jf"
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
		NewMigrateRepoCmd().Command,
		NewCloneRepoCmd().Command,
		NewCloneReposCmd().Command,
		NewMigrateRepoCmd().Command,
		NewMigrateReposCmd().Command,
		NewEnableReplicationsCmd().Command,
		NewDisableReplicationsCmd().Command,
		NewCloneUsersCmd().Command,
		NewClonePermissionsCmd().Command,
		NewCipherCmd().Command,
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

func (o *BaseCmd) buildSyncerAndConnect() (ret *jf.Syncer, err error) {
	executor := buildExecutor(o.DryRunFlag)

	ret, err = jf.NewSyncerAndConnect(
		buildArtifactoryManager(o.Source, executor),
		buildArtifactoryManager(o.Target, executor))
	return
}

func buildExecutor(dryRunFlag *DryRunFlag) (ret exec.Executor) {
	if dryRunFlag.CurrentValue {
		ret = &exec.SkipExecutor{}
	} else {
		ret = &exec.LogExecutor{}
	}
	return
}

func (o *BaseCmd) buildArtifactoryManagerAndConnect() (ret *jf.ArtifactoryManager, err error) {
	executor := buildExecutor(o.DryRunFlag)

	ret = buildArtifactoryManager(o.Source, executor)
	err = ret.Connect()
	return
}

func buildArtifactoryManager(server *ServerFlagLabels, executor exec.Executor) *jf.ArtifactoryManager {
	return &jf.ArtifactoryManager{
		Label:    server.BuildLabel(),
		Url:      server.Url.CurrentValue,
		User:     server.User.CurrentValue,
		Password: server.Password.CurrentValue,

		Executor: executor,
	}
}

func buildArtifactoryManagerAndConnect(
	server *ServerFlagLabels, executor exec.Executor) (ret *jf.ArtifactoryManager, err error) {
	ret = buildArtifactoryManager(server, executor)
	err = ret.Connect()
	return
}
