package cmd

import (
	"github.com/go-ee/jfrog/jf"
	"github.com/urfave/cli/v2"
)

type MigrateReposCmd struct {
	*BaseCmd
}

func NewMigrateReposCmd() (ret *MigrateReposCmd) {
	ret = &MigrateReposCmd{
		BaseCmd: NewBaseCmd(),
	}

	ret.Command = &cli.Command{
		Name:  "migrate-repos",
		Usage: "Create repositories of source Artifactory server in target server and create replications",
		Flags: []cli.Flag{
			ret.Source.Url, ret.Source.User, ret.Source.Password,
			ret.Target.Url, ret.Target.User, ret.Target.Password,
			ret.DryRunFlag,
		},
	}

	ret.Command.Action = func(context *cli.Context) error {
		executor := buildExecutor(ret.DryRunFlag)

		syncer, err := jf.NewSyncerAndConnect(
			buildArtifactoryManager(ret.Source, executor),
			buildArtifactoryManager(ret.Target, executor))

		if err == nil {
			err = syncer.CloneReposAndCreateReplications()
		}
		return err
	}
	return
}
