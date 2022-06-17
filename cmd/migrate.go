package cmd

import (
	"github.com/go-ee/jfrog/core"
	"github.com/go-ee/utils/cliu"
	"github.com/urfave/cli/v2"
)

type MigrateCmd struct {
	*cliu.BaseCommand
	Source     *ServerDef
	Target     *ServerDef
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
		cloner := core.NewCloner(
			buildArtifactoryMigrator(ret.Source), buildArtifactoryMigrator(ret.Target),
			ret.DryRunFlag.CurrentValue)

		err = cloner.Clone()

		return
	}
	return
}

func buildArtifactoryMigrator(server *ServerDef) *core.ArtifactoryManager {
	return &core.ArtifactoryManager{
		Label:    server.Label,
		Url:      server.Url.CurrentValue,
		User:     server.User.CurrentValue,
		Password: server.Password.CurrentValue,
	}
}
