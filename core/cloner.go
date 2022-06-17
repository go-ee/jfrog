package core

import (
	"fmt"
	"github.com/go-ee/utils/exec"
	"github.com/jfrog/jfrog-client-go/artifactory"
	"github.com/jfrog/jfrog-client-go/artifactory/auth"
	"github.com/jfrog/jfrog-client-go/artifactory/services"
	"github.com/jfrog/jfrog-client-go/artifactory/services/utils"
	"github.com/jfrog/jfrog-client-go/config"
	"github.com/sirupsen/logrus"
)

type Cloner struct {
	Source *ArtifactoryManager
	Target *ArtifactoryManager

	cloners map[RepoType]map[PackageType]RepositoryCloner
}

func NewCloner(source *ArtifactoryManager, target *ArtifactoryManager) *Cloner {
	return &Cloner{Source: source, Target: target,
		cloners: map[RepoType]map[PackageType]RepositoryCloner{}}
}

func (o *Cloner) Clone() (err error) {
	logrus.Infof("clone artifactory server from '%v' to '%v'", o.Source.Url, o.Target.Url)

	if err = o.Source.Connect(); err != nil {
		return
	}

	if err = o.Target.Connect(); err != nil {
		return
	}

	var repos *[]services.RepositoryDetails

	repos, err = o.Source.GetAllRepositories()
	for _, repo := range *repos {
		if err = o.CloneRepo(repo); err != nil {
			return
		}
	}
	return
}

func (o *Cloner) CreateReplication(repo services.RepositoryDetails) (err error) {
	repoType := RepoType(repo.Type)
	switch repoType {
	case Local:
		if replicationParams, findErr := o.Source.GetReplication(repo.Key); findErr != nil {
			createReplicationParams := o.Target.buildCreateReplicationParams(repo)
			err = o.Source.Execute(fmt.Sprintf("create PUSH replication: %v", createReplicationParams.Url), func() error {
				return o.Source.CreateReplication(*createReplicationParams)
			})
		} else {
			logrus.Debugf("%v: replication already configured %v", repo.Url, replicationParams)
		}
	case Remote:
		if replicationParams, findErr := o.Target.GetReplication(repo.Key); findErr != nil {
			createReplicationParams := o.Source.buildCreateReplicationParams(repo)
			logrus.Infof("create replication: %v", createReplicationParams.Url)
			err = o.Target.Execute(fmt.Sprintf("create PULL replication: %v", createReplicationParams.Url), func() error {
				return o.Target.CreateReplication(*createReplicationParams)
			})
		} else {
			logrus.Debugf("%v: replication already configured %v", repo.Url, replicationParams)
		}
	default:
		logrus.Infof("no need for replication of '%v' repository '%v'", repoType, repo.Url)
	}
	return err
}

func (o *Cloner) CloneRepo(sourceRepo services.RepositoryDetails) (err error) {

	var repoExists bool
	if repoExists, err = o.Target.IsRepoExists(sourceRepo.Key); err != nil {
		return
	}

	if !repoExists {
		repoType := RepoType(sourceRepo.Type)
		packageType := PackageType(sourceRepo.PackageType)

		var repoCloner RepositoryCloner
		if repoCloner, err = o.getRepoCloner(repoType, packageType); err == nil {
			err = repoCloner.Clone(sourceRepo.Key)
		}
	} else {
		logrus.Infof(o.Target.buildLog("repo already exists " + sourceRepo.Key))
	}
	return
}

func (o *Cloner) getRepoCloner(repoTypo RepoType, packageType PackageType) (ret RepositoryCloner, err error) {
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

type ArtifactoryManager struct {
	artifactory.ArtifactoryServicesManager
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
	return
}

func (o *ArtifactoryManager) buildCreateReplicationParams(
	repo services.RepositoryDetails) (ret *services.CreateReplicationParams) {

	ret = &services.CreateReplicationParams{
		ReplicationParams: utils.ReplicationParams{
			Username:               o.User,
			Password:               o.Password,
			Url:                    o.buildReplicationUrl(repo),
			CronExp:                "",
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

func buildRepoPackageTypeUrlPrefix(repo services.RepositoryDetails) (ret string) {
	switch PackageType(repo.PackageType) {
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

type PackageType string

const (
	Bower       PackageType = "Bower"
	Chef                    = "Chef"
	CocoaPods               = "CocoaPods"
	Conan                   = "Conan"
	Docker                  = "Docker"
	Go                      = "Go"
	NuGet                   = "NuGet"
	Npm                     = "Npm"
	PhpComposer             = "PHP Composer"
	Puppet                  = "Puppet"
	PyPi                    = "Pypi"
	RubyGems                = "RubyGems"
	Generic                 = "Generic"
	Maven                   = "Maven"
	Helm                    = "Helm"
	GitLfs                  = "GitLfs"
	Debian                  = "Debian"
	YUM                     = "YUM"
	Vagrant                 = "Vagrant"
	Cargo                   = "Cargo"
	Gradle                  = "Gradle"
)

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
