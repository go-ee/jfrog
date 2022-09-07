package cmd

import (
	"github.com/go-ee/jfrog/jfrog"
	"github.com/urfave/cli/v2"
)

type ClonePermissionsCmd struct {
	*DoubleServerCmd
}

func NewClonePermissionsCmd() (ret *ClonePermissionsCmd) {
	ret = &ClonePermissionsCmd{
		DoubleServerCmd: NewDoubleServerCmd(),
	}

	ret.Command = &cli.Command{
		Name:  "clone-permissions",
		Usage: "Create permissions of source Artifactory server in target server",
		Flags: []cli.Flag{
			ret.Server.Url, ret.Server.User, ret.Server.Password, ret.Server.Token,
			ret.Target.Url, ret.Target.User, ret.Target.Password, ret.Target.Token,
			ret.DryRunFlag,
		},
	}

	ret.Command.Action = func(context *cli.Context) (err error) {
		var syncer *jfrog.Syncer
		if syncer, err = ret.buildSyncerAndConnect(); err == nil {
			err = syncer.ClonePermissions()
		}
		return
	}
	return
}
