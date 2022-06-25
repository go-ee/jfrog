package jf

import (
	"fmt"
	"github.com/go-ee/utils/exec"
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
	//*access.AccessServicesManager
	Label    string
	Url      string
	User     string
	Password string

	Executor exec.Executor
}

func (o *ArtifactoryManager) Connect() (err error) {
	details := auth.NewArtifactoryDetails()
	details.SetUrl(o.Url)
	details.SetUser(o.User)
	details.SetPassword(o.Password)

	var serviceConfig config.Config
	if serviceConfig, err = config.NewConfigBuilder().
		SetServiceDetails(details).
		SetDryRun(false).
		//SetHttpClient(myCustomClient).
		Build(); err != nil {
		return
	}

	var servicesManager artifactory.ArtifactoryServicesManager
	if servicesManager, err = artifactory.New(serviceConfig); err == nil {
		o.ArtifactoryServicesManager = servicesManager
	}

	/*
		var accessManager *access.AccessServicesManager
		if accessManager, err = access.New(serviceConfig); err == nil {
			o.AccessServicesManager = accessManager
		}
	*/

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
		ReplicationParams: utils.ReplicationParams{
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
		}}
	return
}

func (o *ArtifactoryManager) buildReplicationUrl(repo services.RepositoryDetails) (ret string) {
	return fmt.Sprintf(
		"%v%v%v", o.Url, buildRepoPackageTypeUrlPrefix(repo), repo.Key)
}

func (o *ArtifactoryManager) buildLog(info string) (ret string) {
	return fmt.Sprintf("%v: %v", o.Label, info)
}

func (o *ArtifactoryManager) Execute(info string, execute func() error) (err error) {
	return o.Executor.Execute(o.buildLog(info), execute)
}

func (o *ArtifactoryManager) IsUserExists(userName string) (ret bool, err error) {
	var user *services.User
	user, err = o.GetUser(services.UserParams{UserDetails: services.User{Name: userName}})
	ret = user != nil
	return
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
		ret = "api/pypi/pypi-local/"
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