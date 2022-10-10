package cmd

import (
	"github.com/go-ee/jfrog/jfrog"
	"github.com/go-ee/utils/lg"
	"github.com/urfave/cli/v2"
)

type CloneServersCmd struct {
	*DoubleServerCmd
	ReplicationFlag *ReplicationFlag
}

func NewCloneServesCmd() (ret *CloneServersCmd) {
	ret = &CloneServersCmd{
		DoubleServerCmd: NewDoubleServerCmd(),
		ReplicationFlag: NewReplicationFlag(),
	}

	ret.Command = &cli.Command{
		Name:  "clone-server",
		Usage: "Create repositories, users, etc. of source Artifactory server in target server",
		Flags: []cli.Flag{
			ret.Server.Url, ret.Server.User, ret.Server.Password, ret.Server.Token,
			ret.Target.Url, ret.Target.User, ret.Target.Password, ret.Target.Token,
			ret.DryRunFlag,
		},
	}

	ret.Command.Action = func(context *cli.Context) (err error) {
		var syncer *jfrog.Syncer
		if syncer, err = ret.buildSyncerAndConnect(); err == nil {
			if err = syncer.CloneUsers(); err != nil {
				lg.LOG.Debugf("an error at cloning of users from %v to %v: %v",
					ret.Server.Url, ret.Target.Url, err)
			}
			if err = syncer.CloneRepos("", ret.ReplicationFlag.CurrentValue); err != nil {
				lg.LOG.Debugf("an error at cloning of repos from %v to %v: %v",
					ret.Server.Url, ret.Target.Url, err)
			}
			if err = syncer.ClonePermissions(); err != nil {
				lg.LOG.Debugf("an error at cloning of permissions from %v to %v: %v",
					ret.Server.Url, ret.Target.Url, err)
			}
			if err = syncer.CloneProjects(); err != nil {
				lg.LOG.Debugf("an error at cloning of projects from %v to %v: %v",
					ret.Server.Url, ret.Target.Url, err)
			}
			if err = syncer.CloneGroups(); err != nil {
				lg.LOG.Debugf("an error at cloning of groups from %v to %v: %v",
					ret.Server.Url, ret.Target.Url, err)
			}
		}
		return
	}
	return
}
