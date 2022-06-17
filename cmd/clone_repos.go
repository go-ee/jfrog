package cmd

import (
	"github.com/go-ee/jfrog/core"
	"github.com/go-ee/utils/cliu"
	"github.com/urfave/cli/v2"
)

type CloneReposCmd struct {
	*cliu.BaseCommand
	Source     *ServerFlagLabels
	Target     *ServerFlagLabels
	DryRunFlag *DryRunFlag
}

func NewCloneReposCmd() (ret *CloneReposCmd) {
	ret = &CloneReposCmd{
		BaseCommand: &cliu.BaseCommand{},
		Source:      NewServerDef("source"),
		Target:      NewServerDef("target"),
		DryRunFlag:  NewDryRunFlag(),
	}

	ret.Command = &cli.Command{
		Name:  "clone-repos",
		Usage: "Create repositories of source Artifactory server in target server",
		Flags: []cli.Flag{
			ret.Source.Url, ret.Source.User, ret.Source.Password,
			ret.Target.Url, ret.Target.User, ret.Target.Password,
			ret.DryRunFlag,
		},
	}

	ret.Command.Action = func(context *cli.Context) (err error) {
		executor := buildExecutor(ret.DryRunFlag)

		syncer, err := core.NewSyncerAndConnect(
			buildArtifactoryManager(ret.Source, executor),
			buildArtifactoryManager(ret.Target, executor))

		if err == nil {
			err = syncer.CloneReposAndCreateReplications()
		}
		return
	}
	return
}
