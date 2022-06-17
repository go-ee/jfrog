package core

import (
	"fmt"
	"github.com/jfrog/jfrog-client-go/artifactory"
	"github.com/jfrog/jfrog-client-go/artifactory/auth"
	"github.com/jfrog/jfrog-client-go/artifactory/services"
	"github.com/jfrog/jfrog-client-go/artifactory/services/utils"
	"github.com/jfrog/jfrog-client-go/config"
	"github.com/sirupsen/logrus"
)

type Cloner struct {
	Source      *ArtifactoryManager
	Target      *ArtifactoryManager
	DryRun      bool
	repoCloners map[RepoType]map[PackageType]RepositoryCloner
}

func NewCloner(source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *Cloner {
	return &Cloner{Source: source, Target: target, DryRun: dryRun,
		repoCloners: map[RepoType]map[PackageType]RepositoryCloner{}}
}

func (o *Cloner) Clone() (err error) {

	if err = o.Source.Connect(); err != nil {
		return
	}

	if err = o.Target.Connect(); err != nil {
		return
	}

	var repos *[]services.RepositoryDetails

	repos, err = o.Source.GetAllRepositories()
	for _, repo := range *repos {
		logrus.Infof("migrate %v (%v)", repo.Url, repo.PackageType)
		if err = o.CloneRepo(repo); err != nil {
			return
		}
	}
	return
}

func (o *Cloner) CreateReplication(repo services.RepositoryDetails) (err error) {
	if replicationParams, findErr := o.Source.GetReplication(repo.Key); findErr != nil {
		createReplicationParams := o.Target.buildCreateReplicationParams(repo)
		logrus.Infof("create replication: %v", createReplicationParams.Url)
		if !o.DryRun {
			err = o.Source.CreateReplication(*createReplicationParams)
		}
	} else {
		logrus.Debugf("%v: replication already configured %v", repo.Url, replicationParams)
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
		logrus.Infof("repo already cloned %v", sourceRepo.Url)
	}
	return
}

func (o *Cloner) getRepoCloner(repoTypo RepoType, packageType PackageType) (ret RepositoryCloner, err error) {
	typedRepoCloners := o.repoCloners[repoTypo]
	if typedRepoCloners == nil {
		typedRepoCloners = map[PackageType]RepositoryCloner{}
		o.repoCloners[repoTypo] = typedRepoCloners
	}

	if ret = typedRepoCloners[packageType]; ret == nil {
		switch repoTypo {
		case Local:
			ret, err = BuildLocalRepoCloner(packageType, o.Source, o.Target, o.DryRun)
		case Remote:
			ret, err = BuildRemoteRepoCloner(packageType, o.Source, o.Target, o.DryRun)
		case Virtual:
			ret, err = BuildVirtualRepoCloner(packageType, o.Source, o.Target, o.DryRun)
		case Federated:
			ret, err = BuildFederatedRepoCloner(packageType, o.Source, o.Target, o.DryRun)
		default:
			err = fmt.Errorf("repo type '%v' not supported", repoTypo)
			return
		}
		o.repoCloners[repoTypo] = typedRepoCloners
	}
	return
}

type ArtifactoryManager struct {
	artifactory.ArtifactoryServicesManager
	Label    string
	Url      string
	User     string
	Password string
	DryRun   bool
}

func (o *ArtifactoryManager) Connect() (err error) {
	rtDetails := auth.NewArtifactoryDetails()
	rtDetails.SetUrl(o.Url)
	rtDetails.SetUser(o.User)
	rtDetails.SetPassword(o.Password)

	var serviceConfig config.Config
	if serviceConfig, err = config.NewConfigBuilder().
		SetServiceDetails(rtDetails).
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
	DryRun      bool
}

func (o *RepositoryClonerImpl) run(repoKey string, label string, execute func() error) (err error) {
	logrus.Infof("%v(%v): %v=>%v", label, repoKey, o.Source.Label, o.Target.Label)
	if !o.DryRun {
		err = execute()
	}
	return
}

func (o *RepositoryClonerImpl) clone(repoKey string, execute func() error) (err error) {
	return o.run(repoKey, "clone", execute)
}
