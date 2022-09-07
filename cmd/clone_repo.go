package cmd

import (
	"github.com/go-ee/jfrog/jfrog"
	"github.com/urfave/cli/v2"
)

type CloneRepoCmd struct {
	*DoubleServerCmd
	RepoKeyFlag *RepoKeyFlag
}

func NewCloneRepoCmd() (ret *CloneRepoCmd) {
	ret = &CloneRepoCmd{
		DoubleServerCmd: NewDoubleServerCmd(),
		RepoKeyFlag:     NewRepoKeyFlag(),
	}

	ret.Command = &cli.Command{
		Name:  "clone-repo",
		Usage: "Create repository of source Artifactory server in target server",
		Flags: []cli.Flag{
			ret.Server.Url, ret.Server.User, ret.Server.Password, ret.Server.Token,
			ret.Target.Url, ret.Target.User, ret.Target.Password, ret.Target.Token,
			ret.RepoKeyFlag, ret.DryRunFlag,
		},
	}

	ret.Command.Action = func(context *cli.Context) (err error) {
		var syncer *jfrog.Syncer
		if syncer, err = ret.buildSyncerAndConnect(); err == nil {
			err = syncer.CloneRepo(ret.RepoKeyFlag.CurrentValue)
		}
		return
	}
	return
}
