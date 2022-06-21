package cmd

import (
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
			ret.Source.Url, ret.Source.User, ret.Source.Password,
			ret.DryRunFlag,
		},
	}

	ret.Command.Action = func(context *cli.Context) (err error) {
		executor := buildExecutor(ret.DryRunFlag)

		artifactoryManager := buildArtifactoryManager(ret.Source, executor)
		err = artifactoryManager.Connect()

		if err == nil {
			err = artifactoryManager.EnableReplications()
		}
		return err
	}
	return
}
