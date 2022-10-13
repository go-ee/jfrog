package cmd

import (
	"encoding/json"
	"github.com/go-ee/jfrog/jfrog"
	"github.com/jfrog/jfrog-client-go/artifactory/services/utils"
	"github.com/urfave/cli/v2"
	"os"
)

type CollectTrashCanCmd struct {
	*ServerCmd
	TargetFileFlag *OutputFileFlag
}

func NewCollectTrashCan() (ret *CollectTrashCanCmd) {
	ret = &CollectTrashCanCmd{
		ServerCmd:      NewServerCmd(""),
		TargetFileFlag: NewOutputFileFlag(""),
	}

	ret.TargetFileFlag.Value = "trash-files.json"

	ret.Command = &cli.Command{
		Name:  "collect-trash",
		Usage: "Collect files infos of trash can",
		Flags: []cli.Flag{
			ret.Server.Url, ret.Server.User, ret.Server.Password, ret.Server.Token,
			ret.TargetFileFlag,
			ret.DryRunFlag,
		},
	}

	ret.Command.Action = func(context *cli.Context) (err error) {
		var artifactoryManager *jfrog.ArtifactoryManager
		if artifactoryManager, err = ret.Server.buildArtifactoryManagerAndConnect(ret.DryRunFlag); err == nil {
			var fileListResponse *utils.FileListResponse
			if fileListResponse, err = artifactoryManager.CollectTrashCan(); err != nil {
				return
			}

			file, _ := json.MarshalIndent(fileListResponse, "", " ")
			_ = os.WriteFile(ret.TargetFileFlag.CurrentValue, file, 0644)
		}
		return
	}
	return
}
