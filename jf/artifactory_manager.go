package jf

import (
	"fmt"
	"github.com/go-ee/utils/exec"
	"github.com/jfrog/jfrog-client-go/access"
	accessServices "github.com/jfrog/jfrog-client-go/access/services"
	"github.com/jfrog/jfrog-client-go/artifactory"
	"github.com/jfrog/jfrog-client-go/artifactory/auth"
	"github.com/jfrog/jfrog-client-go/artifactory/services"
	"github.com/jfrog/jfrog-client-go/artifactory/services/utils"
	"github.com/jfrog/jfrog-client-go/config"
	"github.com/sirupsen/logrus"
	"strings"
)

type ArtifactoryManager struct {
	artifactory.ArtifactoryServicesManager
	*access.AccessServicesManager
	ProjectService *accessServices.ProjectService
	Label          string
	Url            string
	User           string
	Password       string
	Token          string

	Executor exec.Executor

	urlArtifactory string
	urlAccess      string
}

func (o *ArtifactoryManager) Connect() (err error) {
	var artifactoryServiceConfig config.Config
	if artifactoryServiceConfig, err = o.buildArtifactoryConfig(); err != nil {
		return
	}

	var servicesManager artifactory.ArtifactoryServicesManager
	if servicesManager, err = artifactory.New(artifactoryServiceConfig); err == nil {
		o.ArtifactoryServicesManager = servicesManager
	}

	var accessToken string
	if accessToken, err = o.getOrCreateAccessToken(); err != nil {
		return
	}

	var accessServiceConfig config.Config
	if accessServiceConfig, err = o.buildAccessConfig(accessToken); err != nil {
		return
	}

	var accessManager *access.AccessServicesManager
	if accessManager, err = access.New(accessServiceConfig); err == nil {
		o.AccessServicesManager = accessManager
	}

	o.ProjectService = accessServices.NewProjectService(accessManager.Client())
	o.ProjectService.ServiceDetails = accessServiceConfig.GetServiceDetails()
	return
}

func (o *ArtifactoryManager) getOrCreateAccessToken() (ret string, err error) {
	if o.Token != "" {
		ret = o.Token
		return
	}

	var tokens []string
	if tokens, err = o.GetUserTokens(o.User); err != nil {
		return
	}

	if len(tokens) > 0 {
		ret = tokens[0]
	} else {
		logrus.Infof("create access token, not implemented yet")
		//err = fmt.Errorf("create access token, not implemented yet")
	}
	return
}

func (o *ArtifactoryManager) buildArtifactoryConfig() (ret config.Config, err error) {
	o.urlArtifactory = fmt.Sprintf("%vartifactory/", o.Url)
	return o.buildConfig(o.urlArtifactory, "")
}

func (o *ArtifactoryManager) buildAccessConfig(accessToken string) (ret config.Config, err error) {
	o.urlAccess = fmt.Sprintf("%vaccess/", o.Url)
	return o.buildConfig(o.urlAccess, accessToken)
}

func (o *ArtifactoryManager) buildConfig(url string, accessToken string) (ret config.Config, err error) {
	details := auth.NewArtifactoryDetails()
	details.SetUrl(url)

	if accessToken != "" {
		//encodedAccessToken := base64.StdEncoding.EncodeToString([]byte(accessToken))
		details.SetAccessToken(accessToken)
	} else {
		details.SetUser(o.User)
		details.SetPassword(o.Password)
	}

	ret, err = config.NewConfigBuilder().
		SetServiceDetails(details).
		SetDryRun(false).
		//SetHttpClient(myCustomClient).
		Build()
	return
}

func (o *ArtifactoryManager) EnableReplications() (err error) {
	err = o.ChangeReplicationsStatus(true)
	return
}

func (o *ArtifactoryManager) DisableReplications() (err error) {
	err = o.ChangeReplicationsStatus(false)
	return
}

func (o *ArtifactoryManager) DisableReplication(repo services.RepositoryDetails) (err error) {
	err = o.ChangeReplicationStatus(repo, false)
	return
}

func (o *ArtifactoryManager) EnableReplication(repo services.RepositoryDetails) (err error) {
	err = o.ChangeReplicationStatus(repo, true)
	return
}

func (o *ArtifactoryManager) ChangeReplicationsStatus(enable bool) (err error) {
	var repos *[]services.RepositoryDetails
	repos, err = o.GetAllRepositories()
	for _, repo := range *repos {
		if err = o.ChangeReplicationStatus(repo, enable); err != nil {
			logrus.Warnf("change replication error, %v, %v", repo.Key, err)
		}
	}
	return
}

func (o *ArtifactoryManager) ChangeReplicationStatus(repo services.RepositoryDetails, enable bool) (err error) {
	if replications, findErr := o.GetReplication(repo.Key); findErr == nil {

		logrus.Debugf(o.buildLog(fmt.Sprintf("disable replication '%v'", repo.Key)))
		for _, replication := range replications {

			if replication.Enabled != enable {
				replication.Enabled = enable
				updateReplicationParams := services.NewUpdateReplicationParams()
				updateReplicationParams.ReplicationParams = replication

				err = o.Execute(fmt.Sprintf("disable replication '%v'", replication.Url),
					func() error {
						return o.UpdateReplication(updateReplicationParams)
					})
			}
		}
	} else {
		logrus.Debugf(o.buildLog(fmt.Sprintf("no replication configured '%v'", repo.Key)))
	}

	return
}

func (o *ArtifactoryManager) buildCreateReplicationParams(
	repo services.RepositoryDetails) (ret *services.CreateReplicationParams) {

	ret = &services.CreateReplicationParams{
		ReplicationParams: o.buildReplicationParams(repo)}
	return
}

func (o *ArtifactoryManager) buildUpdateReplicationParams(
	repo services.RepositoryDetails) (ret *services.UpdateReplicationParams) {

	ret = &services.UpdateReplicationParams{
		ReplicationParams: o.buildReplicationParams(repo)}
	return
}

func (o *ArtifactoryManager) buildReplicationParams(repo services.RepositoryDetails) utils.ReplicationParams {
	return utils.ReplicationParams{
		Username:               o.User,
		Password:               o.Password,
		Url:                    o.buildReplicationUrl(repo),
		CronExp:                "0 0 1 * * ?",
		RepoKey:                repo.Key,
		EnableEventReplication: true,
		SocketTimeoutMillis:    0,
		Enabled:                true,
		SyncDeletes:            false,
		SyncProperties:         true,
		SyncStatistics:         true,
	}
}

func (o *ArtifactoryManager) buildReplicationUrl(repo services.RepositoryDetails) (ret string) {
	return fmt.Sprintf(
		"%v%v%v", o.urlArtifactory, buildRepoPackageTypeUrlPrefix(repo), repo.Key)
}

func (o *ArtifactoryManager) buildLog(info string) (ret string) {
	return fmt.Sprintf("%v: %v", o.Label, info)
}

func (o *ArtifactoryManager) Execute(info string, execute func() error) (err error) {
	return o.Executor.Execute(o.buildLog(info), execute)
}

func (o *ArtifactoryManager) IsUserExists(userName string) (ret bool, err error) {
	var user *services.User
	user, err = o.GetUser(*wrapNameToUserParams(userName))
	ret = user != nil
	return
}

func (o *ArtifactoryManager) IsProjectExists(projectKey string) (ret bool, err error) {
	var project *accessServices.Project
	project, err = o.ProjectService.Get(projectKey)
	ret = project != nil
	return
}

func wrapNameToUserParams(userName string) *services.UserParams {
	return &services.UserParams{UserDetails: services.User{Name: userName}}
}

func wrapNameToGroupParams(groupName string) *services.GroupParams {
	return &services.GroupParams{GroupDetails: services.Group{Name: groupName}}
}

func wrapGroupToGroupParams(group *services.Group) *services.GroupParams {
	return &services.GroupParams{GroupDetails: *group}
}

func wrapProjectToProjectParams(project *accessServices.Project) *accessServices.ProjectParams {
	return &accessServices.ProjectParams{ProjectDetails: *project}
}

func buildRepoPackageTypeUrlPrefix(repo services.RepositoryDetails) (ret string) {
	switch AsPackageType(repo.PackageType) {
	case Bower:
		ret = "api/bower/"
	case Chef:
		ret = "api/chef/"
	case CocoaPods:
		ret = "api/pods/"
	case Conan:
		ret = "api/conan/"
	case Docker:
		ret = "api/docker/"
	case Go:
		ret = "api/go/"
	case NuGet:
		ret = "api/nuget/"
	case Npm:
		ret = "api/npm/"
	case PhpComposer:
		ret = "api/composer/"
	case Puppet:
		ret = "api/puppet/"
	case PyPi:
		ret = "api/pypi/"
	case RubyGems:
		ret = "api/gems/"
	case GitLfs:
		ret = "api/lfs"
	}
	return
}

type RepoType string

const (
	Local     RepoType = "LOCAL"
	Remote             = "REMOTE"
	Virtual            = "VIRTUAL"
	Federated          = "FEDERATED"
)

func AsRepoType(repoType string) RepoType {
	return RepoType(strings.ToUpper(repoType))
}

type PackageType string

const (
	Bower       PackageType = "BOWER"
	Chef                    = "CHEF"
	CocoaPods               = "COCOAPODS"
	Conan                   = "CONAN"
	Docker                  = "DOCKER"
	Go                      = "GO"
	NuGet                   = "NUGET"
	Npm                     = "NPM"
	PhpComposer             = "PHP COMPOSER"
	Puppet                  = "PUPPET"
	PyPi                    = "PYPI"
	RubyGems                = "RUBYGEMS"
	Generic                 = "GENERIC"
	Maven                   = "MAVEN"
	Helm                    = "HELM"
	GitLfs                  = "GITLFS"
	Debian                  = "DEBIAN"
	YUM                     = "YUM"
	Vagrant                 = "VAGRANT"
	Cargo                   = "CARGO"
	Gradle                  = "GRADLE"
)

func AsPackageType(packageType string) PackageType {
	return PackageType(strings.ToUpper(packageType))
}

type RepositoryCloner interface {
	Clone(repoKey string) error
}

type RepositoryClonerImpl struct {
	RepoType    RepoType
	PackageType PackageType
	Source      *ArtifactoryManager
	Target      *ArtifactoryManager

	labelCreateRepository string
}

func (o *RepositoryClonerImpl) buildLabelCreateRepository(repoKey string) string {
	return o.labelCreateRepository + repoKey
}

func NewRepositoryCloner(repoType RepoType, packageType PackageType,
	source *ArtifactoryManager, target *ArtifactoryManager) *RepositoryClonerImpl {
	return &RepositoryClonerImpl{RepoType: repoType, PackageType: packageType,
		Source: source, Target: target,
		labelCreateRepository: fmt.Sprintf("create repository[%v,%v] ", repoType, packageType),
	}
}

func prepareRepositoryBaseParams(params *services.RepositoryBaseParams) {
	//workaround: repos assigned to project without project prefix does not allow to create repo.
	if params.Key != "" && !strings.HasPrefix(params.Key, params.ProjectKey) {
		params.ProjectKey = ""
	}
}

func findNonExistentUsers(
	sources []*services.User, targets []*services.User) (ret []*services.User) {

	targetNames := map[string]*services.User{}
	for _, target := range targets {
		targetNames[target.Name] = target
	}

	for _, source := range sources {
		found := targetNames[source.Name]
		if found == nil {
			ret = append(ret, source)
		}
	}
	return
}

func findNonExistentPermissions(
	sources []*services.PermissionTargetParams, targets []*services.PermissionTargetParams) (
	ret []*services.PermissionTargetParams) {

	targetNames := map[string]*services.PermissionTargetParams{}
	for _, target := range targets {
		targetNames[target.Name] = target
	}

	for _, source := range sources {
		found := targetNames[source.Name]
		if found == nil {
			ret = append(ret, source)
		}
	}
	return
}

func findNonExistentGroups(
	sources []*services.Group, targets []*services.Group) (
	ret []*services.Group) {

	targetNames := map[string]*services.Group{}
	for _, target := range targets {
		targetNames[target.Name] = target
	}

	for _, source := range sources {
		found := targetNames[source.Name]
		if found == nil {
			ret = append(ret, source)
		}
	}
	return
}

func findNonExistentProjects(
	sources []*accessServices.Project, targets []*accessServices.Project) (
	ret []*accessServices.Project) {

	targetNames := map[string]*accessServices.Project{}
	for _, target := range targets {
		targetNames[target.ProjectKey] = target
	}

	for _, source := range sources {
		found := targetNames[source.ProjectKey]
		if found == nil {
			ret = append(ret, source)
		}
	}
	return
}

func findNonExistentRoles(
	sources []*accessServices.Role, targets []*accessServices.Role) (
	ret []*accessServices.Role) {

	targetNames := map[string]*accessServices.Role{}
	for _, target := range targets {
		targetNames[target.Name] = target
	}

	for _, source := range sources {
		found := targetNames[source.Name]
		if found == nil {
			ret = append(ret, source)
		}
	}
	return
}

func findNonExistentProjectUsers(
	sources *accessServices.ProjectUsers, targets *accessServices.ProjectUsers) (
	ret []*accessServices.ProjectUser) {

	targetNames := map[string]*accessServices.ProjectUser{}
	for _, target := range targets.Members {
		targetNames[target.Name] = target
	}

	for _, source := range sources.Members {
		found := targetNames[source.Name]
		if found == nil {
			ret = append(ret, source)
		}
	}
	return
}
