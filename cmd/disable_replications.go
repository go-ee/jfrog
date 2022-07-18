package cmd

import (
	"github.com/go-ee/jfrog/jf"
	"github.com/urfave/cli/v2"
)

type DisableReplicationsCmd struct {
	*ServerCmd
}

func NewDisableReplicationsCmd() (ret *DisableReplicationsCmd) {
	ret = &DisableReplicationsCmd{
		ServerCmd: NewServerCmd(""),
	}

	ret.Command = &cli.Command{
		Name:  "disable-replications",
		Usage: "Disable replications in Artifactory",
		Flags: []cli.Flag{
			ret.Server.Url, ret.Server.User, ret.Server.Password, ret.Server.Token,
			ret.DryRunFlag,
		},
	}

	ret.Command.Action = func(context *cli.Context) (err error) {
		var artifactoryManager *jf.ArtifactoryManager
		if artifactoryManager, err = ret.Server.buildArtifactoryManagerAndConnect(ret.DryRunFlag); err == nil {
			err = artifactoryManager.DisableReplications()
		}
		return
	}
	return
}
