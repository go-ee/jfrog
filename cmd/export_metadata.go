package cmd

import (
	"github.com/go-ee/jfrog/jf"
	"github.com/urfave/cli/v2"
)

type ExportMetaDataCmd struct {
	*ServerCmd
	ServerPathFlag *ServerPathFlag
}

func NewExportMetaDataCmd() (ret *ExportMetaDataCmd) {
	ret = &ExportMetaDataCmd{
		ServerCmd:      NewServerCmd(""),
		ServerPathFlag: NewServerPathFlag(),
	}

	ret.Command = &cli.Command{
		Name:  "export-metadata",
		Usage: "Export metadata of Artifactory to server path",
		Flags: []cli.Flag{
			ret.Server.Url, ret.Server.User, ret.Server.Password, ret.Server.Token, ret.ServerPathFlag,
			ret.DryRunFlag,
		},
	}

	ret.Command.Action = func(context *cli.Context) (err error) {
		var artifactoryManager *jf.ArtifactoryManager
		if artifactoryManager, err = ret.Server.buildArtifactoryManagerAndConnect(ret.DryRunFlag); err == nil {
			err = artifactoryManager.ExportMetadata(ret.ServerPathFlag.CurrentValue)
		}
		return
	}
	return
}
