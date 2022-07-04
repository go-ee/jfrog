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
			ret.Source.Url, ret.Source.User, ret.Source.Password, ret.Source.Token,
			ret.Target.Url, ret.Target.User, ret.Target.Password, ret.Target.Token,
			ret.DryRunFlag,
		},
	}

	ret.Command.Action = func(context *cli.Context) (err error) {
		var syncer *jf.Syncer
		if syncer, err = ret.buildSyncerAndConnect(); err == nil {
			err = syncer.CloneReposAndCreateReplications()
		}
		return err
	}
	return
}
