package cmd

import (
	"github.com/go-ee/jfrog/jf"
	"github.com/urfave/cli/v2"
)

type ClonePermissionsCmd struct {
	*BaseCmd
}

func NewClonePermissionsCmd() (ret *ClonePermissionsCmd) {
	ret = &ClonePermissionsCmd{
		BaseCmd: NewBaseCmd(),
	}

	ret.Command = &cli.Command{
		Name:  "clone-permissions",
		Usage: "Create permissions of source Artifactory server in target server",
		Flags: []cli.Flag{
			ret.Source.Url, ret.Source.User, ret.Source.Password,
			ret.Target.Url, ret.Target.User, ret.Target.Password,
			ret.DryRunFlag,
		},
	}

	ret.Command.Action = func(context *cli.Context) (err error) {
		var syncer *jf.Syncer
		if syncer, err = ret.buildSyncerAndConnect(); err == nil {
			err = syncer.ClonePermissions()
		}
		return
	}
	return
}
