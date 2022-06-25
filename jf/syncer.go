package jf

import (
	"fmt"
	"github.com/jfrog/jfrog-client-go/artifactory/services"
	"github.com/sirupsen/logrus"
)

type Syncer struct {
	Source *ArtifactoryManager
	Target *ArtifactoryManager

	cloners map[RepoType]map[PackageType]RepositoryCloner
}

func NewSyncerAndConnect(source *ArtifactoryManager, target *ArtifactoryManager) (ret *Syncer, err error) {
	ret = &Syncer{Source: source, Target: target, cloners: map[RepoType]map[PackageType]RepositoryCloner{}}
	err = ret.Connect()
	return
}

func (o *Syncer) CloneRepos() (err error) {
	logrus.Infof("create artifactory repos from '%v' to '%v'", o.Source.Url, o.Target.Url)

	var repos *[]services.RepositoryDetails
	repos, err = o.Source.GetAllRepositories()
	for _, repo := range *repos {
		if err = o.cloneRepo(repo); err != nil {
			logrus.Warnf("clone error, %v, %v", repo.Key, err)
		}
	}
	return
}

func (o *Syncer) CloneReposAndCreateReplications() (err error) {
	logrus.Infof("clone artifactory repos and create replications from '%v' to '%v'", o.Source.Url, o.Target.Url)

	var repos *[]services.RepositoryDetails
	repos, err = o.Source.GetAllRepositories()
	for _, repo := range *repos {
		if err = o.cloneRepoAndCreateReplication(repo); err != nil {
			logrus.Warnf("clone and create replication error, %v, %v", repo.Key, err)
		}
	}
	return
}

func (o *Syncer) CloneAndCreateReplication(repoKey string) (err error) {
	var repo services.RepositoryDetails
	if err = o.Source.GetRepository(repoKey, &repo); err == nil {
		err = o.cloneRepoAndCreateReplication(repo)
	}
	return
}

func (o *Syncer) cloneRepoAndCreateReplication(repo services.RepositoryDetails) (err error) {
	if err = o.cloneRepo(repo); err == nil {
		err = o.CreateReplication(repo)
	}
	return
}

func (o *Syncer) Connect() (err error) {
	if err = o.Source.Connect(); err == nil {
		err = o.Target.Connect()
	}
	return
}

func (o *Syncer) CreateReplication(repo services.RepositoryDetails) (err error) {

	var repoExists bool
	if repoExists, err = o.Target.IsRepoExists(repo.Key); err != nil {
		return
	}

	if !repoExists {
		logrus.Infof(o.Target.buildLog(
			fmt.Sprintf("repository does not exists '%v', creation of replication not possible", repo.Key)))
		return
	}

	repoType := AsRepoType(repo.GetRepoType())
	switch repoType {
	case Local:
		err = o.createReplicationLocal(repo)
	default:
		logrus.Infof(o.Target.buildLog(
			fmt.Sprintf("no need for replication of '%v' repository '%v'", repoType, repo.Key)))
	}
	return
}

func (o *Syncer) createReplicationRemote(repo services.RepositoryDetails) (err error) {
	if _, findErr := o.Target.GetReplication(repo.Key); findErr != nil {
		createReplicationParams := o.Source.buildCreateReplicationParams(repo)
		logrus.Infof("create replication: %v", createReplicationParams.Url)
		err = o.Target.Execute(fmt.Sprintf("create PULL replication '%v'", createReplicationParams.Url),
			func() error {
				return o.Target.CreateReplication(*createReplicationParams)
			})
	} else {
		logrus.Infof(o.Target.buildLog(fmt.Sprintf("replication already configured '%v'", repo.Key)))
	}
	return
}

func (o *Syncer) createReplicationLocal(repo services.RepositoryDetails) (err error) {
	if _, findErr := o.Source.GetReplication(repo.Key); findErr != nil {
		createReplicationParams := o.Target.buildCreateReplicationParams(repo)
		err = o.Source.Execute(fmt.Sprintf("create PUSH replication '%v'", createReplicationParams.Url),
			func() error {
				return o.Source.CreateReplication(*createReplicationParams)
			})
	} else {
		logrus.Infof(o.Target.buildLog(fmt.Sprintf("replication already configured '%v'", repo.Key)))
	}
	return
}

func (o *Syncer) CloneRepo(repoKey string) (err error) {
	var repo services.RepositoryDetails
	if err = o.Source.GetRepository(repoKey, &repo); err == nil {
		err = o.cloneRepo(repo)
	}
	return
}

func (o *Syncer) cloneRepo(sourceRepo services.RepositoryDetails) (err error) {

	var repoExists bool
	if repoExists, err = o.Target.IsRepoExists(sourceRepo.Key); err != nil {
		return
	}

	if !repoExists {
		repoType := AsRepoType(sourceRepo.GetRepoType())
		packageType := AsPackageType(sourceRepo.PackageType)

		var repoCloner RepositoryCloner
		if repoCloner, err = o.getRepoCloner(repoType, packageType); err == nil {
			err = repoCloner.Clone(sourceRepo.Key)
		}
	} else {
		logrus.Infof(o.Target.buildLog("repo already exists " + sourceRepo.Key))
	}
	return
}

func (o *Syncer) getRepoCloner(repoTypo RepoType, packageType PackageType) (ret RepositoryCloner, err error) {
	typedRepoCloners := o.cloners[repoTypo]
	if typedRepoCloners == nil {
		typedRepoCloners = map[PackageType]RepositoryCloner{}
		o.cloners[repoTypo] = typedRepoCloners
	}

	if ret = typedRepoCloners[packageType]; ret == nil {
		switch repoTypo {
		case Local:
			ret, err = BuildLocalRepoCloner(packageType, o.Source, o.Target)
		case Remote:
			ret, err = BuildRemoteRepoCloner(packageType, o.Source, o.Target)
		case Virtual:
			ret, err = BuildVirtualRepoCloner(packageType, o.Source, o.Target)
		case Federated:
			ret, err = BuildFederatedRepoCloner(packageType, o.Source, o.Target)
		default:
			err = fmt.Errorf("repo type '%v' not supported", repoTypo)
			return
		}
		o.cloners[repoTypo] = typedRepoCloners
	}
	return
}
func (o *Syncer) CloneUsers() (err error) {
	logrus.Infof("create artifactory users from '%v' to '%v'", o.Source.Url, o.Target.Url)

	var users []*services.User
	users, err = o.Source.GetAllUsers()
	for _, user := range users {
		if err = o.cloneUser(user); err != nil {
			logrus.Warnf("clone error, %v, %v", user, err)
		}
	}
	return
}

func (o *Syncer) cloneUser(user *services.User) (err error) {

	var userExists bool
	if userExists, err = o.Target.IsUserExists(user.Name); err != nil {
		return
	}

	if !userExists {
		err = o.Target.CreateUser(services.UserParams{UserDetails: *user})
	} else {
		logrus.Infof(o.Target.buildLog("user already exists: " + user.Name))
	}
	return
}
