package cmd

import (
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
		executor := buildExecutor(ret.DryRunFlag)

		artifactoryManager := buildArtifactoryManager(ret.Source, executor)
		err = artifactoryManager.Connect()

		if err == nil {
			err = artifactoryManager.DisableReplications()
		}
		return err
	}
	return
}
