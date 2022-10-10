package jfrog

import (
	"fmt"
	"github.com/go-ee/utils/lg"
	"github.com/go-ee/utils/stringu"
	accessServices "github.com/jfrog/jfrog-client-go/access/services"
	"github.com/jfrog/jfrog-client-go/artifactory/services"
	"reflect"
	"strings"
)

type Syncer struct {
	Source *ArtifactoryManager
	Target *ArtifactoryManager

	cloners map[RepoType]map[PackageType]RepositoryCloner
}

func NewSyncerAndConnect(source *ArtifactoryManager, target *ArtifactoryManager) (ret *Syncer, err error) {
	ret = &Syncer{Source: source, Target: target,
		cloners: map[RepoType]map[PackageType]RepositoryCloner{}}
	err = ret.Connect()
	return
}

func (o *Syncer) CloneRepos(packageType string, createReplication bool) (err error) {
	lg.LOG.Infof(
		"clone artifactory repos and create replications from '%v' to '%v'", o.Source.Url, o.Target.Url)

	var repos *[]services.RepositoryDetails
	repos, err = o.Source.GetAllRepositories()
	if packageType == "" {
		for _, repo := range *repos {
			if err = o.cloneRepoAndCreateReplication(repo, createReplication); err != nil {
				lg.LOG.Warnf("clone and create replication error, %v, %v", repo.Key, err)
			}
		}
	} else {
		packageTypeLowCase := strings.ToLower(packageType)
		for _, repo := range *repos {
			if packageTypeLowCase == strings.ToLower(repo.PackageType) {
				if err = o.cloneRepoAndCreateReplication(repo, createReplication); err != nil {
					lg.LOG.Warnf("clone and create replication error, %v, %v", repo.Key, err)
				}
			}
		}
	}
	return
}

func (o *Syncer) Clone(repoKey string, createReplication bool) (err error) {
	var repo services.RepositoryDetails
	if err = o.Source.GetRepository(repoKey, &repo); err == nil {
		err = o.cloneRepoAndCreateReplication(repo, createReplication)
	}
	return
}

func (o *Syncer) cloneRepoAndCreateReplication(repo services.RepositoryDetails, createReplication bool) (err error) {
	if err = o.cloneRepo(repo); err == nil {
		if createReplication {
			err = o.CreateOrUpdateReplication(repo)
		}
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
		lg.LOG.Infof(o.Target.buildLog(
			fmt.Sprintf("repository does not exists '%v', creation of replication not possible", repo.Key)))
		return
	}

	repoType := AsRepoType(repo.GetRepoType())
	switch repoType {
	case Local:
		err = o.createReplicationLocal(repo)
	default:
		lg.LOG.Debugf(o.Target.buildLog(
			fmt.Sprintf("no need for replication of '%v' repository '%v'", repoType, repo.Key)))
	}
	return
}

func (o *Syncer) createReplicationRemote(repo services.RepositoryDetails) (err error) {
	if _, findErr := o.Target.GetReplication(repo.Key); findErr != nil {
		createReplicationParams := o.Source.buildCreateReplicationParams(repo)
		lg.LOG.Infof("create replication: %v", createReplicationParams.Url)
		err = o.Target.Execute(fmt.Sprintf("create PULL replication '%v'", createReplicationParams.Url),
			func() error {
				return o.Target.CreateReplication(*createReplicationParams)
			})
	} else {
		lg.LOG.Infof(o.Target.buildLog(fmt.Sprintf("replication already configured '%v'", repo.Key)))
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
			lg.LOG.Debugf(o.Target.buildLog(fmt.Sprintf("replication already configured '%v'", repo.Key)))
		}
	}
	return
}

func (o *Syncer) CloneRepo(repoKey string, createReplication bool) (err error) {
	var repo services.RepositoryDetails
	if err = o.Source.GetRepository(repoKey, &repo); err == nil {
		err = o.cloneRepoAndCreateReplication(repo, createReplication)
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
		lg.LOG.Debugf(o.Target.buildLog("repo already exists " + sourceRepo.Key))
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
	lg.LOG.Infof("clone artifactory users from '%v' to '%v'", o.Source.Url, o.Target.Url)

	var sourceItems []*services.User
	if sourceItems, err = o.Source.GetAllUsers(); err != nil {
		return
	}
	var targetItems []*services.User
	if targetItems, err = o.Target.GetAllUsers(); err != nil {
		return
	}

	notExistentItems, _ := splitNonExistentAndExistentUsers(sourceItems, targetItems)
	for _, item := range notExistentItems {
		if err = o.cloneUser(item.Name); err != nil {
			lg.LOG.Warnf("clone error, %v, %v", item, err)
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
		if strings.Contains(user.Name, "@") {
			user.Email = user.Name
		} else {
			user.Email = fmt.Sprintf("%v@internal.local", user.Name)
		}
	}
	if user.Password == "" {
		user.Password = stringu.GeneratePassword()
	}
	lg.LOG.Infof("create user %v", user.Name)
	err = o.Target.CreateUser(services.UserParams{UserDetails: *user})
	return
}

func (o *Syncer) ClonePermissions() (err error) {
	lg.LOG.Infof("create artifactory permissions from '%v' to '%v'", o.Source.Url, o.Target.Url)

	var sourceItems []*services.PermissionTargetParams
	if sourceItems, err = o.Source.GetPermissionTargets(); err != nil {
		return
	}
	var targetItems []*services.PermissionTargetParams
	if targetItems, err = o.Target.GetPermissionTargets(); err != nil {
		return
	}

	notExistentItems, existentItems := splitNonExistentAndExistentPermissions(sourceItems, targetItems)
	for _, item := range notExistentItems {
		if err = o.clonePermission(item.Name); err != nil {
			lg.LOG.Warnf("clone error, %v, %v", item, err)
		}
	}

	for _, item := range existentItems {
		if err = o.updatePermission(item.Name); err != nil {
			lg.LOG.Warnf("update permission error, %v, %v", item, err)
		}
	}

	return
}

func (o *Syncer) clonePermission(permissionTargetName string) (err error) {
	if strings.HasPrefix(permissionTargetName, "INTERNAL_") {
		lg.LOG.Debugf("skip creation permission target %v", permissionTargetName)
		return
	}

	var sourcePermissionTargetParams *services.PermissionTargetParams
	if sourcePermissionTargetParams, err = o.Source.GetPermissionTarget(permissionTargetName); err != nil {
		return
	}

	lg.LOG.Infof("create permission target %v", sourcePermissionTargetParams.Name)
	err = o.Target.CreatePermissionTarget(*sourcePermissionTargetParams)
	return
}

func (o *Syncer) updatePermission(permissionTargetName string) (err error) {
	if strings.HasPrefix(permissionTargetName, "INTERNAL_") {
		lg.LOG.Debugf("skip updating permission target %v", permissionTargetName)
		return
	}

	var source, target *services.PermissionTargetParams
	if source, err = o.Source.GetPermissionTarget(permissionTargetName); err != nil {
		return
	}

	if target, err = o.Target.GetPermissionTarget(permissionTargetName); err != nil {
		return
	}

	if !reflect.DeepEqual(source, target) {
		lg.LOG.Infof("update permission target %v", source.Name)
		err = o.Target.UpdatePermissionTarget(*source)
	}
	return
}

func (o *Syncer) CloneGroups() (err error) {
	lg.LOG.Infof("create artifactory groups from '%v' to '%v'", o.Source.Url, o.Target.Url)

	var sourceItems []*services.Group
	if sourceItems, err = o.Source.GetGroups(); err != nil {
		return
	}
	var targetItems []*services.Group
	if targetItems, err = o.Target.GetGroups(); err != nil {
		return
	}

	notExistentItems, _ := splitNonExistentAndExistentGroups(sourceItems, targetItems)
	for _, item := range notExistentItems {
		if err = o.cloneGroup(item.Name); err != nil {
			lg.LOG.Warnf("clone error, %v, %v", item, err)
		}
	}
	return
}

func (o *Syncer) cloneGroup(groupName string) (err error) {
	var group *services.Group
	if group, err = o.Source.GetGroup(*wrapNameToGroupParams(groupName)); err != nil {
		return
	}

	lg.LOG.Infof("create group %v", group.Name)
	err = o.Target.CreateGroup(*wrapGroupToGroupParams(group))
	return
}

func (o *Syncer) CloneProjects() (err error) {
	lg.LOG.Infof("create artifactory projects from '%v' to '%v'", o.Source.Url, o.Target.Url)

	var sourceItems []*accessServices.Project
	if sourceItems, err = o.Source.ProjectService.GetAllProjects(); err != nil {
		return
	}
	var targetItems []*accessServices.Project
	if targetItems, err = o.Target.ProjectService.GetAllProjects(); err != nil {
		return
	}

	notExistentItems, _ := splitNonExistentAndExistentProjects(sourceItems, targetItems)
	for _, item := range notExistentItems {
		if err = o.cloneProject(item.ProjectKey); err != nil {
			lg.LOG.Warnf("clone project error, %v, %v", item, err)
		}
	}

	for _, item := range sourceItems {
		if err = o.cloneProjectRoles(item.ProjectKey); err != nil {
			lg.LOG.Warnf("clone project role error, %v, %v", item, err)
		}

		if err = o.cloneProjectUsers(item.ProjectKey); err != nil {
			lg.LOG.Warnf("clone project user error, %v, %v", item, err)
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
		lg.LOG.Infof("create project %v", sourceItem.ProjectKey)
		err = o.Target.ProjectService.Create(*wrapProjectToProjectParams(sourceItem))
	} else {
		lg.LOG.Debugf("project already exists %v", projectKey)
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

	notExistentItems, existentItems := splitNonExistentAndExistentRoles(sourceItems, targetItems)
	for _, item := range notExistentItems {
		if err = o.cloneProjectRole(projectKey, item); err != nil {
			lg.LOG.Warnf("clone error, %v, %v", item, err)
		}
	}

	for _, item := range existentItems {
		if err = o.updateProjectRole(projectKey, item); err != nil {
			lg.LOG.Warnf("clone error, %v, %v", item, err)
		}
	}
	return
}

func (o *Syncer) cloneProjectRole(projectKey string, role *accessServices.Role) (err error) {
	if role.Type == "PREDEFINED" || role.Type == "ADMIN" {
		lg.LOG.Infof("skip creation of %v project '%v' role '%v'", role.Type, projectKey, role.Name)
		return
	}
	lg.LOG.Infof("create project '%v' role '%v'", projectKey, role.Name)
	err = o.Target.ProjectService.CreateRole(projectKey, role)
	return
}

func (o *Syncer) updateProjectRole(projectKey string, role *accessServices.Role) (err error) {
	if role.Type == "PREDEFINED" || role.Type == "ADMIN" {
		lg.LOG.Infof("skip updating of %v project '%v' role '%v'", role.Type, projectKey, role.Name)
		return
	}
	lg.LOG.Infof("update project '%v' role '%v'", projectKey, role.Name)
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

	notExistentItems, existentItems := findNonExistentProjectUsers(sourceItems.Members, targetItems.Members)
	for _, item := range notExistentItems {
		if err = o.cloneProjectUser(projectKey, item); err != nil {
			lg.LOG.Warnf("clone error, %v, %v", item, err)
		}
	}

	for _, item := range existentItems {
		if err = o.updateProjectUser(projectKey, item); err != nil {
			lg.LOG.Warnf("update error, %v, %v", item, err)
		}
	}
	return
}

func (o *Syncer) cloneProjectUser(projectKey string, user *accessServices.ProjectUser) (err error) {
	lg.LOG.Infof("create project '%v' user '%v'", projectKey, user.Name)
	err = o.Target.ProjectService.UpdateUser(projectKey, user)
	return
}

func (o *Syncer) updateProjectUser(projectKey string, user *accessServices.ProjectUser) (err error) {
	lg.LOG.Infof("update project '%v' user '%v'", projectKey, user.Name)
	err = o.Target.ProjectService.UpdateUser(projectKey, user)
	return
}
