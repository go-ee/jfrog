package syncer

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
	//workaround: repos assigned to projects without project prefix does not allow create repo.
	if params.Key != "" && !strings.HasPrefix(params.Key, params.ProjectKey) {
		params.ProjectKey = ""
	}
}
