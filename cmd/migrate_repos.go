package cmd

import (
	"github.com/go-ee/jfrog/jf"
	"github.com/urfave/cli/v2"
)

type MigrateReposCmd struct {
	*DoubleServerCmd
}

func NewMigrateReposCmd() (ret *MigrateReposCmd) {
	ret = &MigrateReposCmd{
		DoubleServerCmd: NewDoubleServerCmd(),
	}

	ret.Command = &cli.Command{
		Name:  "migrate-repos",
		Usage: "Create repositories of source Artifactory server in target server and create replications",
		Flags: []cli.Flag{
			ret.Server.Url, ret.Server.User, ret.Server.Password, ret.Server.Token,
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
