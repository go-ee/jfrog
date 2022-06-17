package core

import (
	"fmt"
	"github.com/jfrog/jfrog-client-go/artifactory/services"
)

func BuildVirtualRepoCloner(packageType PackageType,
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) (ret RepositoryCloner, err error) {

	switch packageType {
	case Bower:
		ret = NewBowerVirtualRepositoryCloner(source, target, dryRun)
	case Chef:
		ret = NewChefVirtualRepositoryCloner(source, target, dryRun)
	case Conan:
		ret = NewConanVirtualRepositoryCloner(source, target, dryRun)
	case Docker:
		ret = NewDockerVirtualRepositoryCloner(source, target, dryRun)
	case Go:
		ret = NewGoVirtualRepositoryCloner(source, target, dryRun)
	case NuGet:
		ret = NewNuGetVirtualRepositoryCloner(source, target, dryRun)
	case Npm:
		ret = NewNpmVirtualRepositoryCloner(source, target, dryRun)
	case Puppet:
		ret = NewPuppetVirtualRepositoryCloner(source, target, dryRun)
	case PyPi:
		ret = NewPyPiVirtualRepositoryCloner(source, target, dryRun)
	case RubyGems:
		ret = NewRubyVirtualGemsRepositoryCloner(source, target, dryRun)
	case Generic:
		ret = NewGenericVirtualRepositoryCloner(source, target, dryRun)
	case Maven:
		ret = NewMavenVirtualRepositoryCloner(source, target, dryRun)
	case Helm:
		ret = NewHelmVirtualRepositoryCloner(source, target, dryRun)
	case GitLfs:
		ret = NewGitLfsVirtualRepositoryCloner(source, target, dryRun)
	case Debian:
		ret = NewDebianVirtualRepositoryCloner(source, target, dryRun)
	case YUM:
		ret = NewYumVirtualRepositoryCloner(source, target, dryRun)
	case Gradle:
		ret = NewGradleVirtualRepositoryCloner(source, target, dryRun)
	default:
		err = fmt.Errorf("repo type '%v' with package type '%v' not supported", Virtual, packageType)

	}
	return
}

func NewVirtualRepositoryClonerImpl(packageType PackageType,
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *RepositoryClonerImpl {
	return &RepositoryClonerImpl{RepoType: Virtual, PackageType: packageType,
		Source: source, Target: target, DryRun: dryRun}
}

type BowerVirtualRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewBowerVirtualRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *BowerVirtualRepositoryCloner {
	return &BowerVirtualRepositoryCloner{
		NewVirtualRepositoryClonerImpl(Bower, source, target, dryRun)}
}

func (o *BowerVirtualRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.BowerVirtualRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateVirtualRepository().Bower(sourceRepoDetails)
		})
	}
	return
}

type ChefVirtualRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewChefVirtualRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *ChefVirtualRepositoryCloner {
	return &ChefVirtualRepositoryCloner{
		NewVirtualRepositoryClonerImpl(Chef, source, target, dryRun)}
}

func (o *ChefVirtualRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.ChefVirtualRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateVirtualRepository().Chef(sourceRepoDetails)
		})
	}
	return
}

type ConanVirtualRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewConanVirtualRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *ConanVirtualRepositoryCloner {
	return &ConanVirtualRepositoryCloner{
		NewVirtualRepositoryClonerImpl(Conan, source, target, dryRun)}
}

func (o *ConanVirtualRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.ConanVirtualRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateVirtualRepository().Conan(sourceRepoDetails)
		})
	}
	return
}

type DockerVirtualRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewDockerVirtualRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *DockerVirtualRepositoryCloner {
	return &DockerVirtualRepositoryCloner{
		NewVirtualRepositoryClonerImpl(Docker, source, target, dryRun)}
}

func (o *DockerVirtualRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.DockerVirtualRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateVirtualRepository().Docker(sourceRepoDetails)
		})
	}
	return
}

type GoVirtualRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewGoVirtualRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *GoVirtualRepositoryCloner {
	return &GoVirtualRepositoryCloner{
		NewVirtualRepositoryClonerImpl(Go, source, target, dryRun)}
}

func (o *GoVirtualRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.GoVirtualRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateVirtualRepository().Go(sourceRepoDetails)
		})
	}
	return
}

type NuGetVirtualRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewNuGetVirtualRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *NuGetVirtualRepositoryCloner {
	return &NuGetVirtualRepositoryCloner{
		NewVirtualRepositoryClonerImpl(NuGet, source, target, dryRun)}
}

func (o *NuGetVirtualRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.NugetVirtualRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateVirtualRepository().Nuget(sourceRepoDetails)
		})
	}
	return
}

type NpmVirtualRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewNpmVirtualRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *NpmVirtualRepositoryCloner {
	return &NpmVirtualRepositoryCloner{
		NewVirtualRepositoryClonerImpl(Npm, source, target, dryRun)}
}

func (o *NpmVirtualRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.NpmVirtualRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateVirtualRepository().Npm(sourceRepoDetails)
		})
	}
	return
}

type PuppetVirtualRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewPuppetVirtualRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *PuppetVirtualRepositoryCloner {
	return &PuppetVirtualRepositoryCloner{
		NewVirtualRepositoryClonerImpl(Puppet, source, target, dryRun)}
}

func (o *PuppetVirtualRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.PuppetVirtualRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateVirtualRepository().Puppet(sourceRepoDetails)
		})
	}
	return
}

type PyPiVirtualRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewPyPiVirtualRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *PyPiVirtualRepositoryCloner {
	return &PyPiVirtualRepositoryCloner{
		NewVirtualRepositoryClonerImpl(PyPi, source, target, dryRun)}
}

func (o *PyPiVirtualRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.PypiVirtualRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateVirtualRepository().Pypi(sourceRepoDetails)
		})
	}
	return
}

type RubyVirtualGemsRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewRubyVirtualGemsRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *RubyVirtualGemsRepositoryCloner {
	return &RubyVirtualGemsRepositoryCloner{
		NewVirtualRepositoryClonerImpl(RubyGems, source, target, dryRun)}
}

func (o *RubyVirtualGemsRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.GemsVirtualRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateVirtualRepository().Gems(sourceRepoDetails)
		})
	}
	return
}

type GenericVirtualRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewGenericVirtualRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *GenericVirtualRepositoryCloner {
	return &GenericVirtualRepositoryCloner{
		NewVirtualRepositoryClonerImpl(Generic, source, target, dryRun)}
}

func (o *GenericVirtualRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.GenericVirtualRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateVirtualRepository().Generic(sourceRepoDetails)
		})
	}
	return
}

type MavenVirtualRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewMavenVirtualRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *MavenVirtualRepositoryCloner {
	return &MavenVirtualRepositoryCloner{
		NewVirtualRepositoryClonerImpl(Maven, source, target, dryRun)}
}

func (o *MavenVirtualRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.MavenVirtualRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateVirtualRepository().Maven(sourceRepoDetails)
		})
	}
	return
}

type HelmVirtualRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewHelmVirtualRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *HelmVirtualRepositoryCloner {
	return &HelmVirtualRepositoryCloner{
		NewVirtualRepositoryClonerImpl(Helm, source, target, dryRun)}
}

func (o *HelmVirtualRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.HelmVirtualRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateVirtualRepository().Helm(sourceRepoDetails)
		})
	}
	return
}

type GitLfsVirtualRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewGitLfsVirtualRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *GitLfsVirtualRepositoryCloner {
	return &GitLfsVirtualRepositoryCloner{
		NewVirtualRepositoryClonerImpl(GitLfs, source, target, dryRun)}
}

func (o *GitLfsVirtualRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.GitlfsVirtualRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateVirtualRepository().Gitlfs(sourceRepoDetails)
		})
	}
	return
}

type DebianVirtualRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewDebianVirtualRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *DebianVirtualRepositoryCloner {
	return &DebianVirtualRepositoryCloner{
		NewVirtualRepositoryClonerImpl(GitLfs, source, target, dryRun)}
}

func (o *DebianVirtualRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.DebianVirtualRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateVirtualRepository().Debian(sourceRepoDetails)
		})
	}
	return
}

type YumVirtualRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewYumVirtualRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *YumVirtualRepositoryCloner {
	return &YumVirtualRepositoryCloner{
		NewVirtualRepositoryClonerImpl(YUM, source, target, dryRun)}
}

func (o *YumVirtualRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.YumVirtualRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateVirtualRepository().Yum(sourceRepoDetails)
		})
	}
	return
}

type GradleVirtualRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewGradleVirtualRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *GradleVirtualRepositoryCloner {
	return &GradleVirtualRepositoryCloner{
		NewVirtualRepositoryClonerImpl(Gradle, source, target, dryRun)}
}

func (o *GradleVirtualRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.GradleVirtualRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateVirtualRepository().Gradle(sourceRepoDetails)
		})
	}
	return
}
