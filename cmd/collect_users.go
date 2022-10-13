package cmd

import (
	"github.com/go-ee/jfrog/jfrog"
	"github.com/go-ee/utils/lg"
	"github.com/urfave/cli/v2"
	"strings"
)

type CollectPermissionTargetManagersCmd struct {
	*ServerCmd
}

func NewCollectUsersByAccessLevelsCmd() (ret *CollectPermissionTargetManagersCmd) {
	ret = &CollectPermissionTargetManagersCmd{
		ServerCmd: NewServerCmd(""),
	}

	ret.Command = &cli.Command{
		Name:  "collect-users-by-access-levels",
		Usage: "Collect users by access levels",
		Flags: []cli.Flag{
			ret.Server.Url, ret.Server.User, ret.Server.Password, ret.Server.Token,
			ret.DryRunFlag,
		},
	}

	ret.Command.Action = func(context *cli.Context) (err error) {
		var artifactoryManager *jfrog.ArtifactoryManager
		if artifactoryManager, err = ret.Server.buildArtifactoryManagerAndConnect(ret.DryRunFlag); err == nil {
			var accessLevelToUsers map[string][]string
			if accessLevelToUsers, err = artifactoryManager.CollectUsersByAccessLevels(); err != nil {
				return
			}
			for accessLevel, usersByAccessLevel := range accessLevelToUsers {
				lg.LOG.Infof("accessLevel: %v, users: %v", accessLevel, strings.Join(usersByAccessLevel, ";"))
			}
		}
		return
	}
	return
}
