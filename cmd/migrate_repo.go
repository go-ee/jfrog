package cmd

import (
	"github.com/go-ee/jfrog/core"
	"github.com/urfave/cli/v2"
)

type MigrateRepoCmd struct {
	*BaseCmd
	RepoKeyFlag *RepoKeyFlag
}

func NewMigrateRepoCmd() (ret *MigrateRepoCmd) {
	ret = &MigrateRepoCmd{
		BaseCmd:     NewBaseCmd(),
		RepoKeyFlag: NewRepoKeyFlag(),
	}

	ret.Command = &cli.Command{
		Name:  "migrate-repo",
		Usage: "Create repository of source Artifactory server in target server, and create replication",
		Flags: []cli.Flag{
			ret.Source.Url, ret.Source.User, ret.Source.Password,
			ret.Target.Url, ret.Target.User, ret.Target.Password,
			ret.RepoKeyFlag, ret.DryRunFlag,
		},
	}

	ret.Command.Action = func(context *cli.Context) (err error) {
		executor := buildExecutor(ret.DryRunFlag)

		syncer, err := core.NewSyncerAndConnect(
			buildArtifactoryManager(ret.Source, executor),
			buildArtifactoryManager(ret.Target, executor))

		if err == nil {
			err = syncer.CloneAndCreateReplication(ret.RepoKeyFlag.CurrentValue)
		}
		return
	}
	return
}
