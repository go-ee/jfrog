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

type Migrator struct {
	Source *ArtifactoryMigrator
	Target *ArtifactoryMigrator
	DryRun bool
}

func (o *Migrator) Migrate() (err error) {

	if err = o.Source.Connect(); err != nil {
		return
	}

	if err = o.Target.Connect(); err != nil {
		return
	}

	var repos *[]services.RepositoryDetails

	repos, err = o.Source.manager.GetAllRepositories()
	for _, repo := range *repos {
		logrus.Infof("%v (%v)", repo.Url, repo.PackageType)
		if err = o.CreateReplication(repo); err != nil {
			return
		}
	}
	return
}

func (o *Migrator) MigrateRepo(repo services.RepositoryDetails) (err error) {
	var targetRepo services.RepositoryDetails
	if targetRepo, err = o.Target.GetOrCreateRepo(repo); err != nil {
		return
	}
	logrus.Infof("%v", targetRepo)
	return
}

func (o *Migrator) CreateReplication(repo services.RepositoryDetails) (err error) {
	if replicationParams, findErr := o.Source.manager.GetReplication(repo.Key); findErr != nil {
		createReplicationParams := o.Target.buildCreateReplicationParams(repo)
		logrus.Infof("create replication: %v", createReplicationParams.Url)
		if !o.DryRun {
			err = o.Source.manager.CreateReplication(*createReplicationParams)
		}
	} else {
		logrus.Infof("%v: replication %v", repo.Url, replicationParams)
	}
	return err
}

type ArtifactoryMigrator struct {
	Url      string
	User     string
	Password string
	DryRun   bool
	manager  artifactory.ArtifactoryServicesManager
}

func (o *ArtifactoryMigrator) Connect() (err error) {
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
		o.manager = servicesManager
	}
	return
}

func (o *ArtifactoryMigrator) GetOrCreateRepo(sourceRepo services.RepositoryDetails) (
	ret services.RepositoryDetails, err error) {

	if repoErr := o.manager.GetRepository(sourceRepo.Key, ret); repoErr != nil {
		switch sourceRepo.Type {
		case "local":
			switch sourceRepo.PackageType {
			case "generic":
				params := services.NewGenericLocalRepositoryParams()
				params.Key = sourceRepo.Key
				params.Description = sourceRepo.Description
				params.Notes = "These are internal notes for generic-repo"
				params.RepoLayoutRef = "simple-default"
				//params.ArchiveBrowsingEnabled = true
				//params.XrayIndex = true
				params.IncludesPattern = "**/*"
				params.ExcludesPattern = "excludedDir/*"
				//params.DownloadRedirect = true

				//err = o.manager.CreateLocalRepository().Generic(params)
			}

		}
	}
	return
}

func (o *ArtifactoryMigrator) buildCreateReplicationParams(
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

func (o *ArtifactoryMigrator) buildReplicationUrl(repo services.RepositoryDetails) (ret string) {
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
