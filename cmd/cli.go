package cmd

import (
	"github.com/go-ee/jfrog/jfrog"
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
		NewCloneServesCmd().Command,
		NewCloneRepoCmd().Command,
		NewCloneReposCmd().Command,
		NewEnableReplicationsCmd().Command,
		NewDisableReplicationsCmd().Command,
		NewCloneUsersCmd().Command,
		NewClonePermissionsCmd().Command,
		NewCloneGroupsCmd().Command,
		NewCloneProjectsCmd().Command,
		NewCipherCmd().Command,
		NewExportUsersCmd().Command,
		NewExportMetaDataCmd().Command,
		NewCollectUsersByAccessLevelsCmd().Command,
		NewCollectTrashCan().Command,
		cliu.NewMarkdownCmd(ret.App).Command,
	}
	return
}

type ServerCmd struct {
	*cliu.BaseCommand
	Server     *ServerDef
	DryRunFlag *DryRunFlag
}

func NewServerCmd(serverLabel string) *ServerCmd {
	return &ServerCmd{
		BaseCommand: &cliu.BaseCommand{},
		Server:      NewServerDef(serverLabel),
		DryRunFlag:  NewDryRunFlag()}
}

type DoubleServerCmd struct {
	*ServerCmd
	Target     *ServerDef
	DryRunFlag *DryRunFlag
}

func NewDoubleServerCmd() *DoubleServerCmd {
	cmd := NewServerCmd("source")
	return &DoubleServerCmd{
		ServerCmd:  cmd,
		Target:     NewServerDef("target"),
		DryRunFlag: NewDryRunFlag()}
}

func (o *DoubleServerCmd) buildSyncerAndConnect() (ret *jfrog.Syncer, err error) {
	executor := buildExecutor(o.DryRunFlag)

	ret, err = jfrog.NewSyncerAndConnect(
		buildArtifactoryManager(o.Server, executor),
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

func (o *ServerDef) buildArtifactoryManagerAndConnect(
	dryRunFlag *DryRunFlag) (ret *jfrog.ArtifactoryManager, err error) {

	executor := buildExecutor(dryRunFlag)

	ret = buildArtifactoryManager(o, executor)
	err = ret.Connect()
	return
}

func buildArtifactoryManager(server *ServerDef, executor exec.Executor) *jfrog.ArtifactoryManager {
	return &jfrog.ArtifactoryManager{
		Label:    server.BuildLabel(),
		Url:      server.Url.NormalizedUrl(),
		User:     server.User.CurrentValue,
		Password: server.Password.CurrentValue,
		Token:    server.Token.CurrentValue,

		Executor: executor,
	}
}

func buildArtifactoryManagerAndConnect(
	server *ServerDef, executor exec.Executor) (ret *jfrog.ArtifactoryManager, err error) {
	ret = buildArtifactoryManager(server, executor)
	err = ret.Connect()
	return
}
