package cmd

import (
	"github.com/go-ee/jfrog/jfrog"
	"github.com/urfave/cli/v2"
)

type CloneReposCmd struct {
	*DoubleServerCmd
	PackageTypeFlag *PackageTypeFlag
	ReplicationFlag *ReplicationFlag
}

func NewCloneReposCmd() (ret *CloneReposCmd) {
	ret = &CloneReposCmd{
		DoubleServerCmd: NewDoubleServerCmd(),
		PackageTypeFlag: NewPackageTypeFlag(),
		ReplicationFlag: NewReplicationFlag(),
	}

	ret.Command = &cli.Command{
		Name:  "clone-repos",
		Usage: "Create repositories of source Artifactory server in target server",
		Flags: []cli.Flag{
			ret.Server.Url, ret.Server.User, ret.Server.Password, ret.Server.Token,
			ret.Target.Url, ret.Target.User, ret.Target.Password, ret.Target.Token,
			ret.PackageTypeFlag,
			ret.DryRunFlag,
		},
	}

	ret.Command.Action = func(context *cli.Context) (err error) {
		var syncer *jfrog.Syncer
		if syncer, err = ret.buildSyncerAndConnect(); err == nil {
			err = syncer.CloneRepos(
				ret.PackageTypeFlag.CurrentValue, ret.ReplicationFlag.CurrentValue)
		}
		return
	}
	return
}
