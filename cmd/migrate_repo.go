package cmd

import (
	"github.com/go-ee/jfrog/jf"
	"github.com/urfave/cli/v2"
)

type MigrateRepoCmd struct {
	*DoubleServerCmd
	RepoKeyFlag *RepoKeyFlag
}

func NewMigrateRepoCmd() (ret *MigrateRepoCmd) {
	ret = &MigrateRepoCmd{
		DoubleServerCmd: NewDoubleServerCmd(),
		RepoKeyFlag:     NewRepoKeyFlag(),
	}

	ret.Command = &cli.Command{
		Name:  "migrate-repo",
		Usage: "Create repository of source Artifactory server in target server, and create replication",
		Flags: []cli.Flag{
			ret.Server.Url, ret.Server.User, ret.Server.Password, ret.Server.Token,
			ret.Target.Url, ret.Target.User, ret.Target.Password, ret.Target.Token,
			ret.RepoKeyFlag, ret.DryRunFlag,
		},
	}

	ret.Command.Action = func(context *cli.Context) (err error) {
		var syncer *jf.Syncer
		if syncer, err = ret.buildSyncerAndConnect(); err == nil {
			err = syncer.CloneAndCreateReplication(ret.RepoKeyFlag.CurrentValue)
		}
		return
	}
	return
}
