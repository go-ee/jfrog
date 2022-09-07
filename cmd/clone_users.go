package cmd

import (
	"github.com/go-ee/jfrog/jfrog"
	"github.com/urfave/cli/v2"
)

type CloneUsersCmd struct {
	*DoubleServerCmd
}

func NewCloneUsersCmd() (ret *CloneUsersCmd) {
	ret = &CloneUsersCmd{
		DoubleServerCmd: NewDoubleServerCmd(),
	}

	ret.Command = &cli.Command{
		Name:  "clone-users",
		Usage: "Create users of source Artifactory server in target server",
		Flags: []cli.Flag{
			ret.Server.Url, ret.Server.User, ret.Server.Password, ret.Server.Token,
			ret.Target.Url, ret.Target.User, ret.Target.Password, ret.Target.Token,
			ret.DryRunFlag,
		},
	}

	ret.Command.Action = func(context *cli.Context) (err error) {
		var syncer *jfrog.Syncer
		if syncer, err = ret.buildSyncerAndConnect(); err == nil {
			err = syncer.CloneUsers()
		}
		return
	}
	return
}
