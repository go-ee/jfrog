package jfrog

import (
	"fmt"
	"github.com/jfrog/jfrog-client-go/artifactory/services"
)

func BuildVirtualRepoCloner(packageType PackageType,
	source *ArtifactoryManager, target *ArtifactoryManager) (ret RepositoryCloner, err error) {

	switch packageType {
	case Bower:
		ret = NewBowerVirtualRepositoryCloner(source, target)
	case Chef:
		ret = NewChefVirtualRepositoryCloner(source, target)
	case Conan:
		ret = NewConanVirtualRepositoryCloner(source, target)
	case Docker:
		ret = NewDockerVirtualRepositoryCloner(source, target)
	case Go:
		ret = NewGoVirtualRepositoryCloner(source, target)
	case NuGet:
		ret = NewNuGetVirtualRepositoryCloner(source, target)
	case Npm:
		ret = NewNpmVirtualRepositoryCloner(source, target)
	case Puppet:
		ret = NewPuppetVirtualRepositoryCloner(source, target)
	case PyPi:
		ret = NewPyPiVirtualRepositoryCloner(source, target)
	case RubyGems:
		ret = NewRubyVirtualGemsRepositoryCloner(source, target)
	case Generic:
		ret = NewGenericVirtualRepositoryCloner(source, target)
	case Maven:
		ret = NewMavenVirtualRepositoryCloner(source, target)
	case Helm:
		ret = NewHelmVirtualRepositoryCloner(source, target)
	case GitLfs:
		ret = NewGitLfsVirtualRepositoryCloner(source, target)
	case Debian:
		ret = NewDebianVirtualRepositoryCloner(source, target)
	case YUM:
		ret = NewYumVirtualRepositoryCloner(source, target)
	case Gradle:
		ret = NewGradleVirtualRepositoryCloner(source, target)
	default:
		err = fmt.Errorf("repo type '%v' with package type '%v' not supported", Virtual, packageType)

	}
	return
}

func NewVirtualRepositoryClonerImpl(packageType PackageType,
	source *ArtifactoryManager, target *ArtifactoryManager) *RepositoryClonerImpl {
	return NewRepositoryCloner(Virtual, packageType, source, target)
}

type BowerVirtualRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewBowerVirtualRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *BowerVirtualRepositoryCloner {
	return &BowerVirtualRepositoryCloner{
		NewVirtualRepositoryClonerImpl(Bower, source, target)}
}

func (o *BowerVirtualRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.BowerVirtualRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			prepareRepositoryBaseParams(&sourceRepoDetails.RepositoryBaseParams)
			return o.Target.CreateVirtualRepository().Bower(sourceRepoDetails)
		})
	}
	return
}

type ChefVirtualRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewChefVirtualRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *ChefVirtualRepositoryCloner {
	return &ChefVirtualRepositoryCloner{
		NewVirtualRepositoryClonerImpl(Chef, source, target)}
}

func (o *ChefVirtualRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.ChefVirtualRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			prepareRepositoryBaseParams(&sourceRepoDetails.RepositoryBaseParams)
			return o.Target.CreateVirtualRepository().Chef(sourceRepoDetails)
		})
	}
	return
}

type ConanVirtualRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewConanVirtualRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *ConanVirtualRepositoryCloner {
	return &ConanVirtualRepositoryCloner{
		NewVirtualRepositoryClonerImpl(Conan, source, target)}
}

func (o *ConanVirtualRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.ConanVirtualRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			prepareRepositoryBaseParams(&sourceRepoDetails.RepositoryBaseParams)
			return o.Target.CreateVirtualRepository().Conan(sourceRepoDetails)
		})
	}
	return
}

type DockerVirtualRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewDockerVirtualRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *DockerVirtualRepositoryCloner {
	return &DockerVirtualRepositoryCloner{
		NewVirtualRepositoryClonerImpl(Docker, source, target)}
}

func (o *DockerVirtualRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.DockerVirtualRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			prepareRepositoryBaseParams(&sourceRepoDetails.RepositoryBaseParams)
			return o.Target.CreateVirtualRepository().Docker(sourceRepoDetails)
		})
	}
	return
}

type GoVirtualRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewGoVirtualRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *GoVirtualRepositoryCloner {
	return &GoVirtualRepositoryCloner{
		NewVirtualRepositoryClonerImpl(Go, source, target)}
}

func (o *GoVirtualRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.GoVirtualRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			prepareRepositoryBaseParams(&sourceRepoDetails.RepositoryBaseParams)
			return o.Target.CreateVirtualRepository().Go(sourceRepoDetails)
		})
	}
	return
}

type NuGetVirtualRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewNuGetVirtualRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *NuGetVirtualRepositoryCloner {
	return &NuGetVirtualRepositoryCloner{
		NewVirtualRepositoryClonerImpl(NuGet, source, target)}
}

func (o *NuGetVirtualRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.NugetVirtualRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			prepareRepositoryBaseParams(&sourceRepoDetails.RepositoryBaseParams)
			return o.Target.CreateVirtualRepository().Nuget(sourceRepoDetails)
		})
	}
	return
}

type NpmVirtualRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewNpmVirtualRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *NpmVirtualRepositoryCloner {
	return &NpmVirtualRepositoryCloner{
		NewVirtualRepositoryClonerImpl(Npm, source, target)}
}

func (o *NpmVirtualRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.NpmVirtualRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			prepareRepositoryBaseParams(&sourceRepoDetails.RepositoryBaseParams)
			return o.Target.CreateVirtualRepository().Npm(sourceRepoDetails)
		})
	}
	return
}

type PuppetVirtualRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewPuppetVirtualRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *PuppetVirtualRepositoryCloner {
	return &PuppetVirtualRepositoryCloner{
		NewVirtualRepositoryClonerImpl(Puppet, source, target)}
}

func (o *PuppetVirtualRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.PuppetVirtualRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			prepareRepositoryBaseParams(&sourceRepoDetails.RepositoryBaseParams)
			return o.Target.CreateVirtualRepository().Puppet(sourceRepoDetails)
		})
	}
	return
}

type PyPiVirtualRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewPyPiVirtualRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *PyPiVirtualRepositoryCloner {
	return &PyPiVirtualRepositoryCloner{
		NewVirtualRepositoryClonerImpl(PyPi, source, target)}
}

func (o *PyPiVirtualRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.PypiVirtualRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			prepareRepositoryBaseParams(&sourceRepoDetails.RepositoryBaseParams)
			return o.Target.CreateVirtualRepository().Pypi(sourceRepoDetails)
		})
	}
	return
}

type RubyVirtualGemsRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewRubyVirtualGemsRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *RubyVirtualGemsRepositoryCloner {
	return &RubyVirtualGemsRepositoryCloner{
		NewVirtualRepositoryClonerImpl(RubyGems, source, target)}
}

func (o *RubyVirtualGemsRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.GemsVirtualRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			prepareRepositoryBaseParams(&sourceRepoDetails.RepositoryBaseParams)
			return o.Target.CreateVirtualRepository().Gems(sourceRepoDetails)
		})
	}
	return
}

type GenericVirtualRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewGenericVirtualRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *GenericVirtualRepositoryCloner {
	return &GenericVirtualRepositoryCloner{
		NewVirtualRepositoryClonerImpl(Generic, source, target)}
}

func (o *GenericVirtualRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.GenericVirtualRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			prepareRepositoryBaseParams(&sourceRepoDetails.RepositoryBaseParams)
			return o.Target.CreateVirtualRepository().Generic(sourceRepoDetails)
		})
	}
	return
}

type MavenVirtualRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewMavenVirtualRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *MavenVirtualRepositoryCloner {
	return &MavenVirtualRepositoryCloner{
		NewVirtualRepositoryClonerImpl(Maven, source, target)}
}

func (o *MavenVirtualRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.MavenVirtualRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			prepareRepositoryBaseParams(&sourceRepoDetails.RepositoryBaseParams)
			return o.Target.CreateVirtualRepository().Maven(sourceRepoDetails)
		})
	}
	return
}

type HelmVirtualRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewHelmVirtualRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *HelmVirtualRepositoryCloner {
	return &HelmVirtualRepositoryCloner{
		NewVirtualRepositoryClonerImpl(Helm, source, target)}
}

func (o *HelmVirtualRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.HelmVirtualRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			prepareRepositoryBaseParams(&sourceRepoDetails.RepositoryBaseParams)
			return o.Target.CreateVirtualRepository().Helm(sourceRepoDetails)
		})
	}
	return
}

type GitLfsVirtualRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewGitLfsVirtualRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *GitLfsVirtualRepositoryCloner {
	return &GitLfsVirtualRepositoryCloner{
		NewVirtualRepositoryClonerImpl(GitLfs, source, target)}
}

func (o *GitLfsVirtualRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.GitlfsVirtualRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			prepareRepositoryBaseParams(&sourceRepoDetails.RepositoryBaseParams)
			return o.Target.CreateVirtualRepository().Gitlfs(sourceRepoDetails)
		})
	}
	return
}

type DebianVirtualRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewDebianVirtualRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *DebianVirtualRepositoryCloner {
	return &DebianVirtualRepositoryCloner{
		NewVirtualRepositoryClonerImpl(GitLfs, source, target)}
}

func (o *DebianVirtualRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.DebianVirtualRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			prepareRepositoryBaseParams(&sourceRepoDetails.RepositoryBaseParams)
			return o.Target.CreateVirtualRepository().Debian(sourceRepoDetails)
		})
	}
	return
}

type YumVirtualRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewYumVirtualRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *YumVirtualRepositoryCloner {
	return &YumVirtualRepositoryCloner{
		NewVirtualRepositoryClonerImpl(YUM, source, target)}
}

func (o *YumVirtualRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.YumVirtualRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			prepareRepositoryBaseParams(&sourceRepoDetails.RepositoryBaseParams)
			return o.Target.CreateVirtualRepository().Yum(sourceRepoDetails)
		})
	}
	return
}

type GradleVirtualRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewGradleVirtualRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *GradleVirtualRepositoryCloner {
	return &GradleVirtualRepositoryCloner{
		NewVirtualRepositoryClonerImpl(Gradle, source, target)}
}

func (o *GradleVirtualRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.GradleVirtualRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			prepareRepositoryBaseParams(&sourceRepoDetails.RepositoryBaseParams)
			return o.Target.CreateVirtualRepository().Gradle(sourceRepoDetails)
		})
	}
	return
}
