package cmd

import (
	"github.com/go-ee/jfrog/jf"
	"github.com/urfave/cli/v2"
)

type EnableReplicationsCmd struct {
	*ServerCmd
}

func NewEnableReplicationsCmd() (ret *EnableReplicationsCmd) {
	ret = &EnableReplicationsCmd{
		ServerCmd: NewServerCmd("server"),
	}

	ret.Command = &cli.Command{
		Name:  "enable-replications",
		Usage: "Enable replications in Artifactory",
		Flags: []cli.Flag{
			ret.Server.Url, ret.Server.User, ret.Server.Password, ret.Server.Token,
			ret.DryRunFlag,
		},
	}

	ret.Command.Action = func(context *cli.Context) (err error) {
		var artifactoryManager *jf.ArtifactoryManager
		if artifactoryManager, err = ret.Server.buildArtifactoryManagerAndConnect(ret.DryRunFlag); err == nil {
			err = artifactoryManager.EnableReplications()
		}
		return
	}
	return
}
