package cmd

import (
	"github.com/go-ee/jfrog/jf"
	"github.com/urfave/cli/v2"
)

type CloneRepoCmd struct {
	*BaseCmd
	RepoKeyFlag *RepoKeyFlag
}

func NewCloneRepoCmd() (ret *CloneRepoCmd) {
	ret = &CloneRepoCmd{
		BaseCmd:     NewBaseCmd(),
		RepoKeyFlag: NewRepoKeyFlag(),
	}

	ret.Command = &cli.Command{
		Name:  "clone-repo",
		Usage: "Create repository of source Artifactory server in target server",
		Flags: []cli.Flag{
			ret.Source.Url, ret.Source.User, ret.Source.Password,
			ret.Target.Url, ret.Target.User, ret.Target.Password,
			ret.RepoKeyFlag, ret.DryRunFlag,
		},
	}

	ret.Command.Action = func(context *cli.Context) (err error) {
		var syncer *jf.Syncer
		if syncer, err = ret.buildSyncerAndConnect(); err == nil {
			err = syncer.CloneRepo(ret.RepoKeyFlag.CurrentValue)
		}
		return
	}
	return
}
