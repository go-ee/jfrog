package cmd

import (
	"github.com/go-ee/jfrog/jf"
	"github.com/go-ee/utils/cliu"
	"github.com/go-ee/utils/exec"
	"github.com/go-ee/utils/lg"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
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
	log := lg.NewZapProdLogger()
	return &ServerCmd{
		BaseCommand: &cliu.BaseCommand{Log: log},
		Server:      NewServerDef(serverLabel, log),
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
		Target:     NewServerDef("target", cmd.Log),
		DryRunFlag: NewDryRunFlag()}
}

func (o *DoubleServerCmd) buildSyncerAndConnect() (ret *jf.Syncer, err error) {
	executor := buildExecutor(o.DryRunFlag, o.Log)

	ret, err = jf.NewSyncerAndConnect(
		buildArtifactoryManager(o.Server, executor),
		buildArtifactoryManager(o.Target, executor))
	return
}

func buildExecutor(dryRunFlag *DryRunFlag, logger *zap.SugaredLogger) (ret exec.Executor) {
	if dryRunFlag.CurrentValue {
		ret = &exec.SkipExecutor{Log: logger}
	} else {
		ret = &exec.LogExecutor{Log: logger}
	}
	return
}

func (o *ServerDef) buildArtifactoryManagerAndConnect(
	dryRunFlag *DryRunFlag) (ret *jf.ArtifactoryManager, err error) {

	executor := buildExecutor(dryRunFlag, o.Log)

	ret = buildArtifactoryManager(o, executor)
	err = ret.Connect()
	return
}

func buildArtifactoryManager(server *ServerDef, executor exec.Executor) *jf.ArtifactoryManager {
	return &jf.ArtifactoryManager{
		Label:    server.BuildLabel(),
		Url:      server.Url.NormalizedUrl(),
		User:     server.User.CurrentValue,
		Password: server.Password.CurrentValue,
		Token:    server.Token.CurrentValue,
		Log:      server.Log,

		Executor: executor,
	}
}

func buildArtifactoryManagerAndConnect(
	server *ServerDef, executor exec.Executor) (ret *jf.ArtifactoryManager, err error) {
	ret = buildArtifactoryManager(server, executor)
	err = ret.Connect()
	return
}
