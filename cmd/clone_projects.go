package cmd

import (
	"github.com/go-ee/jfrog/jf"
	"github.com/urfave/cli/v2"
)

type CloneProjectsCmd struct {
	*BaseCmd
	Projects *ProjectsFlag
}

func NewCloneProjectsCmd() (ret *CloneProjectsCmd) {
	ret = &CloneProjectsCmd{
		BaseCmd:  NewBaseCmd(),
		Projects: NewProjectsFlag(),
	}

	ret.Command = &cli.Command{
		Name:  "clone-projects",
		Usage: "Create projects of source Artifactory server in target server",
		Flags: []cli.Flag{
			ret.Source.Url, ret.Source.User, ret.Source.Password,
			ret.Target.Url, ret.Target.User, ret.Target.Password,
			ret.Projects, ret.DryRunFlag,
		},
	}

	ret.Command.Action = func(context *cli.Context) (err error) {
		var syncer *jf.Syncer
		if syncer, err = ret.buildSyncerAndConnect(); err == nil {
			err = syncer.CloneProjects(ret.Projects.ProjectKeys())
		}
		return
	}
	return
}
