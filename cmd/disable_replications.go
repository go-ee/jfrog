package cmd

import (
	"github.com/go-ee/jfrog/jf"
	"github.com/urfave/cli/v2"
)

type DisableReplicationsCmd struct {
	*BaseCmd
}

func NewDisableReplicationsCmd() (ret *DisableReplicationsCmd) {
	ret = &DisableReplicationsCmd{
		BaseCmd: NewBaseCmd(),
	}

	ret.Command = &cli.Command{
		Name:  "disable-replications",
		Usage: "Disable replications in Artifactory",
		Flags: []cli.Flag{
			ret.Source.Url, ret.Source.User, ret.Source.Password,
			ret.DryRunFlag,
		},
	}

	ret.Command.Action = func(context *cli.Context) (err error) {
		var artifactoryManager *jf.ArtifactoryManager
		if artifactoryManager, err = ret.buildArtifactoryManagerAndConnect(); err == nil {
			err = artifactoryManager.DisableReplications()
		}
		return
	}
	return
}
