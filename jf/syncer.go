package jf

import (
	"fmt"
	"github.com/go-ee/utils/stringu"
	accessServices "github.com/jfrog/jfrog-client-go/access/services"
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
		err = o.CreateOrUpdateReplication(repo)
	}
	return
}

func (o *Syncer) Connect() (err error) {
	if err = o.Source.Connect(); err == nil {
		err = o.Target.Connect()
	}
	return
}

func (o *Syncer) CreateOrUpdateReplication(repo services.RepositoryDetails) (err error) {

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
	if reps, findErr := o.Source.GetReplication(repo.Key); findErr != nil {
		createReplicationParams := o.Target.buildCreateReplicationParams(repo)
		err = o.Source.Execute(fmt.Sprintf("create PUSH replication '%v'", createReplicationParams.Url),
			func() error {
				return o.Source.CreateReplication(*createReplicationParams)
			})
	} else {
		rep := reps[0]
		updateRepParams := o.Target.buildUpdateReplicationParams(repo)

		if rep.Username != updateRepParams.Username ||
			//rep.Password != updateRepParams.Password ||
			rep.Url != updateRepParams.Url {
			err = o.Source.Execute(fmt.Sprintf("update PUSH replication '%v'", updateRepParams.Url),
				func() error {
					return o.Source.UpdateReplication(*updateRepParams)
				})
		} else {
			logrus.Debugf(o.Target.buildLog(fmt.Sprintf("replication already configured '%v'", repo.Key)))
		}
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
		logrus.Debugf(o.Target.buildLog("repo already exists " + sourceRepo.Key))
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
	logrus.Infof("create artifactory sourceItems from '%v' to '%v'", o.Source.Url, o.Target.Url)

	var sourceItems []*services.User
	if sourceItems, err = o.Source.GetAllUsers(); err != nil {
		return
	}
	var targetItems []*services.User
	if targetItems, err = o.Target.GetAllUsers(); err != nil {
		return
	}

	notExistentItems := findNonExistentUsers(sourceItems, targetItems)
	for _, item := range notExistentItems {
		if err = o.cloneUser(item.Name); err != nil {
			logrus.Warnf("clone error, %v, %v", item, err)
		}
	}
	return
}

func (o *Syncer) cloneUser(userName string) (err error) {

	var user *services.User
	if user, err = o.Source.GetUser(*wrapNameToUserParams(userName)); err != nil {
		return
	}
	if user.Email == "" {
		user.Email = user.Name
	}
	if user.Password == "" {
		user.Password = stringu.GeneratePassword()
	}
	logrus.Infof("create user %v", user.Name)
	err = o.Target.CreateUser(services.UserParams{UserDetails: *user})
	return
}

func (o *Syncer) ClonePermissions() (err error) {
	logrus.Infof("create artifactory permissions from '%v' to '%v'", o.Source.Url, o.Target.Url)

	var sourceItems []*services.PermissionTargetParams
	if sourceItems, err = o.Source.GetPermissionTargets(); err != nil {
		return
	}
	var targetItems []*services.PermissionTargetParams
	if targetItems, err = o.Target.GetPermissionTargets(); err != nil {
		return
	}

	notExistentItems := findNonExistentPermissions(sourceItems, targetItems)
	for _, item := range notExistentItems {
		if err = o.clonePermission(item.Name); err != nil {
			logrus.Warnf("clone error, %v, %v", item, err)
		}
	}
	return
}

func (o *Syncer) clonePermission(permissionTargetName string) (err error) {
	var sourcePermissionTargetParams *services.PermissionTargetParams
	if sourcePermissionTargetParams, err = o.Source.GetPermissionTarget(permissionTargetName); err != nil {
		return
	}

	logrus.Infof("create permission target %v", sourcePermissionTargetParams.Name)
	err = o.Target.CreatePermissionTarget(*sourcePermissionTargetParams)
	return
}

func (o *Syncer) CloneGroups() (err error) {
	logrus.Infof("create artifactory groups from '%v' to '%v'", o.Source.Url, o.Target.Url)

	var sourceItems []*services.Group
	if sourceItems, err = o.Source.GetGroups(); err != nil {
		return
	}
	var targetItems []*services.Group
	if targetItems, err = o.Target.GetGroups(); err != nil {
		return
	}

	notExistentItems := findNonExistentGroups(sourceItems, targetItems)
	for _, item := range notExistentItems {
		if err = o.cloneGroup(item.Name); err != nil {
			logrus.Warnf("clone error, %v, %v", item, err)
		}
	}
	return
}

func (o *Syncer) cloneGroup(groupName string) (err error) {
	var group *services.Group
	if group, err = o.Source.GetGroup(*wrapNameToGroupParams(groupName)); err != nil {
		return
	}

	logrus.Infof("create group %v", group.Name)
	err = o.Target.CreateGroup(*wrapGroupToGroupParams(group))
	return
}

func (o *Syncer) CloneProjects() (err error) {
	logrus.Infof("create artifactory projects from '%v' to '%v'", o.Source.Url, o.Target.Url)

	var sourceItems []*accessServices.Project
	if sourceItems, err = o.Source.ProjectService.GetAllProjects(); err != nil {
		return
	}
	var targetItems []*accessServices.Project
	if targetItems, err = o.Target.ProjectService.GetAllProjects(); err != nil {
		return
	}

	notExistentItems := findNonExistentProjects(sourceItems, targetItems)
	for _, item := range notExistentItems {
		if err = o.cloneProject(item.ProjectKey); err != nil {
			logrus.Warnf("clone project error, %v, %v", item, err)
		}
	}

	for _, item := range sourceItems {
		if err = o.cloneProjectRoles(item.ProjectKey); err != nil {
			logrus.Warnf("clone project role error, %v, %v", item, err)
		}

		if err = o.cloneProjectUsers(item.ProjectKey); err != nil {
			logrus.Warnf("clone project user error, %v, %v", item, err)
		}
	}
	return
}

func (o *Syncer) cloneProject(projectKey string) (err error) {
	var sourceItem *accessServices.Project
	if sourceItem, err = o.Source.ProjectService.Get(projectKey); err != nil {
		return
	}

	var projectExists bool
	if projectExists, err = o.Target.IsProjectExists(projectKey); !projectExists {
		logrus.Infof("create project %v", sourceItem.ProjectKey)
		err = o.Target.ProjectService.Create(*wrapProjectToProjectParams(sourceItem))
	} else {
		logrus.Debugf("project already exists %v", projectKey)
	}
	return
}

func (o *Syncer) cloneProjectRoles(projectKey string) (err error) {
	var sourceItems []*accessServices.Role
	if sourceItems, err = o.Source.ProjectService.GetRoles(projectKey); err != nil {
		return
	}

	var targetItems []*accessServices.Role
	if targetItems, err = o.Target.ProjectService.GetRoles(projectKey); err != nil {
		return
	}

	notExistentItems := findNonExistentRoles(sourceItems, targetItems)
	for _, item := range notExistentItems {
		if err = o.cloneProjectRole(projectKey, item); err != nil {
			logrus.Warnf("clone error, %v, %v", item, err)
		}
	}
	return
}

func (o *Syncer) cloneProjectRole(projectKey string, role *accessServices.Role) (err error) {
	logrus.Infof("create project '%v' role '%v'", projectKey, role.Name)
	err = o.Target.ProjectService.CreateRole(projectKey, role)
	return
}

func (o *Syncer) cloneProjectUsers(projectKey string) (err error) {
	var sourceItems *accessServices.ProjectUsers
	if sourceItems, err = o.Source.ProjectService.GetUsers(projectKey); err != nil {
		return
	}

	var targetItems *accessServices.ProjectUsers
	if targetItems, err = o.Target.ProjectService.GetUsers(projectKey); err != nil {
		return
	}

	notExistentItems := findNonExistentProjectUsers(sourceItems, targetItems)
	for _, item := range notExistentItems {
		if err = o.cloneProjectUser(projectKey, item); err != nil {
			logrus.Warnf("clone error, %v, %v", item, err)
		}
	}
	return
}

func (o *Syncer) cloneProjectUser(projectKey string, user *accessServices.ProjectUser) (err error) {
	logrus.Infof("create project '%v' user '%v'", projectKey, user.Name)
	err = o.Target.ProjectService.UpdateUser(projectKey, user)
	return
}
