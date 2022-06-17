package cmd

import (
	"github.com/go-ee/jfrog/core"
	"github.com/go-ee/utils/cliu"
	"github.com/go-ee/utils/exec"
	"github.com/urfave/cli/v2"
	"strings"
)

type MigrateCmd struct {
	*cliu.BaseCommand
	Source     *ServerFlagLabels
	Target     *ServerFlagLabels
	DryRunFlag *DryRunFlag
}

func NewMigrateCmd() (ret *MigrateCmd) {
	ret = &MigrateCmd{
		BaseCommand: &cliu.BaseCommand{},
		Source:      NewServerDef("source"),
		Target:      NewServerDef("target"),
		DryRunFlag:  NewDryRunFlag(),
	}

	ret.Command = &cli.Command{
		Name:  "migrate-by-replication",
		Usage: "Create target repositories and create replications for all local repositories",
		Flags: []cli.Flag{
			ret.Source.Url, ret.Source.User, ret.Source.Password,
			ret.Target.Url, ret.Target.User, ret.Target.Password,
			ret.DryRunFlag,
		},
	}

	ret.Command.Action = func(context *cli.Context) (err error) {
		executor := buildExecutor(ret.DryRunFlag)

		cloner := core.NewCloner(
			buildArtifactoryManager(ret.Source, executor),
			buildArtifactoryManager(ret.Target, executor))

		err = cloner.Clone()

		return
	}
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

func buildArtifactoryManager(server *ServerFlagLabels, executor exec.Executor) *core.ArtifactoryManager {
	return &core.ArtifactoryManager{
		Label:    buildLabel(server),
		Url:      server.Url.CurrentValue,
		User:     server.User.CurrentValue,
		Password: server.Password.CurrentValue,

		Executor: executor,
	}
}

func buildLabel(server *ServerFlagLabels) (ret string) {
	ret = strings.TrimPrefix(server.Url.CurrentValue, "https://")
	ret = strings.Split(ret, ".")[0]
	return ret
}
