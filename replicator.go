package jfrog

import (
	"fmt"
	"github.com/jfrog/jfrog-client-go/artifactory"
	"github.com/jfrog/jfrog-client-go/artifactory/auth"
	"github.com/jfrog/jfrog-client-go/artifactory/services"
	"github.com/jfrog/jfrog-client-go/artifactory/services/utils"
	"github.com/jfrog/jfrog-client-go/config"
	"github.com/sirupsen/logrus"
)

type Replicator struct {
	Source *ArtifactoryReplicator
	Target *ArtifactoryReplicator
	DryRun bool
}

type ArtifactoryReplicator struct {
	Url             string
	User            string
	Password        string
	DryRun          bool
	servicesManager artifactory.ArtifactoryServicesManager
}

func (o *ArtifactoryReplicator) Connect() (err error) {
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
		o.servicesManager = servicesManager
	}
	return
}

func (o *Replicator) CreateRepoReplications() (err error) {

	if err = o.Source.Connect(); err != nil {
		return
	}

	if err = o.Target.Connect(); err != nil {
		return
	}

	var repos *[]services.RepositoryDetails

	params := services.NewRepositoriesFilterParams()
	params.RepoType = "local"

	repos, err = o.Source.servicesManager.GetAllRepositoriesFiltered(params)
	for _, repo := range *repos {
		logrus.Infof("%v (%v)", repo.Url, repo.PackageType)
		if err = o.CreateReplication(repo); err != nil {
			return
		}
	}
	return
}

func (o *Replicator) CreateReplication(repo services.RepositoryDetails) (err error) {
	if replicationParams, findErr := o.Source.servicesManager.GetReplication(repo.Key); findErr != nil {
		createReplicationParams := o.Target.BuildCreateReplicationParams(repo)
		logrus.Infof("create replication: %v", createReplicationParams.Url)
		if !o.DryRun {
			err = o.Source.servicesManager.CreateReplication(*createReplicationParams)
		}
	} else {
		logrus.Infof("%v: replication %v", repo.Url, replicationParams)
	}
	return err
}

func (o *ArtifactoryReplicator) BuildCreateReplicationParams(
	repo services.RepositoryDetails) (ret *services.CreateReplicationParams) {

	ret = &services.CreateReplicationParams{
		ReplicationParams: utils.ReplicationParams{
			Username:               o.User,
			Password:               o.Password,
			Url:                    o.BuildReplicationUrl(repo),
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

func (o *ArtifactoryReplicator) BuildReplicationUrl(repo services.RepositoryDetails) (ret string) {
	return fmt.Sprintf(
		"%v%v%v", o.Url, buildRepoPackageTypeUrlPrefix(repo), repo.Key)
}

func buildRepoPackageTypeUrlPrefix(repo services.RepositoryDetails) (ret string) {
	switch repo.PackageType {
	case "Bower":
		ret = "api/bower/"
	case "Chef":
		ret = "api/chef/"
	case "CocoaPods":
		ret = "api/pods/"
	case "Conan":
		ret = "api/conan/"
	case "Docker":
		ret = "api/docker/"
	case "Go":
		ret = "api/go/"
	case "NuGet":
		ret = "api/nuget/"
	case "Npm":
		ret = "api/npm/"
	case "PHP Composer":
		ret = "api/composer/"
	case "Puppet":
		ret = "api/puppet/"
	case "PyPI":
		ret = "api/pypi/pypi-local/"
	case "RubyGems":
		ret = "api/gems/"
	}
	return
}
