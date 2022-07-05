package cmd

import (
	"github.com/go-ee/jfrog/jf"
	"github.com/jfrog/jfrog-client-go/artifactory/services"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type ExportUsersCmd struct {
	*ServerCmd
	UsersFileFlag *UsersFileFlag
}

func NewExportUsersCmd() (ret *ExportUsersCmd) {
	ret = &ExportUsersCmd{
		ServerCmd:     NewServerCmd(""),
		UsersFileFlag: NewUsersFileFlag(),
	}

	ret.Command = &cli.Command{
		Name:  "export-users",
		Usage: "Export users to an yaml file",
		Flags: []cli.Flag{
			ret.Server.Url, ret.Server.User, ret.Server.Password, ret.Server.Token,
			ret.UsersFileFlag,
			ret.DryRunFlag,
		},
	}

	ret.Command.Action = func(context *cli.Context) (err error) {
		var artifactoryManager *jf.ArtifactoryManager
		if artifactoryManager, err = ret.Server.buildArtifactoryManagerAndConnect(ret.DryRunFlag); err == nil {
			var users []*services.User
			if users, err = artifactoryManager.GetAllUsers(); err != nil {
				return
			}
			var data []byte
			if data, err = yaml.Marshal(&users); err == nil {
				err = ioutil.WriteFile(ret.UsersFileFlag.CurrentValue, data, 0)
			}
		}
		return
	}
	return
}
