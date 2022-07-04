package cmd

import (
	"github.com/go-ee/jfrog/jf"
	"github.com/urfave/cli/v2"
)

type EnableReplicationsCmd struct {
	*BaseCmd
}

func NewEnableReplicationsCmd() (ret *EnableReplicationsCmd) {
	ret = &EnableReplicationsCmd{
		BaseCmd: NewBaseCmd(),
	}

	ret.Command = &cli.Command{
		Name:  "enable-replications",
		Usage: "Enable replications in Artifactory",
		Flags: []cli.Flag{
			ret.Source.Url, ret.Source.User, ret.Source.Password, ret.Source.Token,
			ret.DryRunFlag,
		},
	}

	ret.Command.Action = func(context *cli.Context) (err error) {
		var artifactoryManager *jf.ArtifactoryManager
		if artifactoryManager, err = ret.buildArtifactoryManagerAndConnect(); err == nil {
			err = artifactoryManager.EnableReplications()
		}
		return
	}
	return
}
