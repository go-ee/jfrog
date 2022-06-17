package core

import (
	"fmt"
	"github.com/jfrog/jfrog-client-go/artifactory/services"
)

func BuildRemoteRepoCloner(packageType PackageType,
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) (ret RepositoryCloner, err error) {

	switch packageType {
	case Bower:
		ret = NewBowerRemoteRepositoryCloner(source, target, dryRun)
	case Chef:
		ret = NewChefRemoteRepositoryCloner(source, target, dryRun)
	case CocoaPods:
		ret = NewCocoaPodsRemotePodsRepositoryCloner(source, target, dryRun)
	case Conan:
		ret = NewConanRemoteRepositoryCloner(source, target, dryRun)
	case Docker:
		ret = NewDockerRemoteRepositoryCloner(source, target, dryRun)
	case Go:
		ret = NewGoRemoteRepositoryCloner(source, target, dryRun)
	case NuGet:
		ret = NewNuGetRemoteRepositoryCloner(source, target, dryRun)
	case Npm:
		ret = NewNpmRemoteRepositoryCloner(source, target, dryRun)
	case PhpComposer:
		ret = NewPhpComposerRemoteRepositoryCloner(source, target, dryRun)
	case Puppet:
		ret = NewPuppetRemoteRepositoryCloner(source, target, dryRun)
	case PyPi:
		ret = NewPyPiRemoteRepositoryCloner(source, target, dryRun)
	case RubyGems:
		ret = NewRubyRemoteGemsRepositoryCloner(source, target, dryRun)
	case Generic:
		ret = NewGenericRemoteRepositoryCloner(source, target, dryRun)
	case Maven:
		ret = NewMavenRemoteRepositoryCloner(source, target, dryRun)
	case Helm:
		ret = NewHelmRemoteRepositoryCloner(source, target, dryRun)
	case GitLfs:
		ret = NewGitLfsRemoteRepositoryCloner(source, target, dryRun)
	case Debian:
		ret = NewDebianRemoteRepositoryCloner(source, target, dryRun)
	case YUM:
		ret = NewYumRemoteRepositoryCloner(source, target, dryRun)
	case Cargo:
		ret = NewCargoRemoteRepositoryCloner(source, target, dryRun)
	case Gradle:
		ret = NewGradleRemoteRepositoryCloner(source, target, dryRun)
	default:
		err = fmt.Errorf("repo type '%v' with package type '%v' not supported", Remote, packageType)

	}
	return
}

func NewRemoteRepositoryClonerImpl(packageType PackageType,
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *RepositoryClonerImpl {
	return &RepositoryClonerImpl{RepoType: Remote, PackageType: packageType,
		Source: source, Target: target, DryRun: dryRun}
}

type BowerRemoteRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewBowerRemoteRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *BowerRemoteRepositoryCloner {
	return &BowerRemoteRepositoryCloner{
		NewRemoteRepositoryClonerImpl(Bower, source, target, dryRun)}
}

func (o *BowerRemoteRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.BowerRemoteRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateRemoteRepository().Bower(sourceRepoDetails)
		})
	}
	return
}

type ChefRemoteRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewChefRemoteRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *ChefRemoteRepositoryCloner {
	return &ChefRemoteRepositoryCloner{
		NewRemoteRepositoryClonerImpl(Chef, source, target, dryRun)}
}

func (o *ChefRemoteRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.ChefRemoteRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateRemoteRepository().Chef(sourceRepoDetails)
		})
	}
	return
}

type CocoaPodsRemotePodsRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewCocoaPodsRemotePodsRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *CocoaPodsRemotePodsRepositoryCloner {
	return &CocoaPodsRemotePodsRepositoryCloner{
		NewRemoteRepositoryClonerImpl(CocoaPods, source, target, dryRun)}
}

func (o *CocoaPodsRemotePodsRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.CocoapodsRemoteRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateRemoteRepository().Cocoapods(sourceRepoDetails)
		})
	}
	return
}

type ConanRemoteRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewConanRemoteRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *ConanRemoteRepositoryCloner {
	return &ConanRemoteRepositoryCloner{
		NewRemoteRepositoryClonerImpl(Conan, source, target, dryRun)}
}

func (o *ConanRemoteRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.ConanRemoteRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateRemoteRepository().Conan(sourceRepoDetails)
		})
	}
	return
}

type DockerRemoteRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewDockerRemoteRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *DockerRemoteRepositoryCloner {
	return &DockerRemoteRepositoryCloner{
		NewRemoteRepositoryClonerImpl(Docker, source, target, dryRun)}
}

func (o *DockerRemoteRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.DockerRemoteRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateRemoteRepository().Docker(sourceRepoDetails)
		})
	}
	return
}

type GoRemoteRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewGoRemoteRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *GoRemoteRepositoryCloner {
	return &GoRemoteRepositoryCloner{
		NewRemoteRepositoryClonerImpl(Go, source, target, dryRun)}
}

func (o *GoRemoteRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.GoRemoteRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateRemoteRepository().Go(sourceRepoDetails)
		})
	}
	return
}

type NuGetRemoteRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewNuGetRemoteRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *NuGetRemoteRepositoryCloner {
	return &NuGetRemoteRepositoryCloner{
		NewRemoteRepositoryClonerImpl(NuGet, source, target, dryRun)}
}

func (o *NuGetRemoteRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.NugetRemoteRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateRemoteRepository().Nuget(sourceRepoDetails)
		})
	}
	return
}

type NpmRemoteRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewNpmRemoteRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *NpmRemoteRepositoryCloner {
	return &NpmRemoteRepositoryCloner{
		NewRemoteRepositoryClonerImpl(Npm, source, target, dryRun)}
}

func (o *NpmRemoteRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.NpmRemoteRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateRemoteRepository().Npm(sourceRepoDetails)
		})
	}
	return
}

type PhpComposerRemoteRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewPhpComposerRemoteRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *PhpComposerRemoteRepositoryCloner {
	return &PhpComposerRemoteRepositoryCloner{
		NewRemoteRepositoryClonerImpl(PhpComposer, source, target, dryRun)}
}

func (o *PhpComposerRemoteRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.ComposerRemoteRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateRemoteRepository().Composer(sourceRepoDetails)
		})
	}
	return
}

type PuppetRemoteRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewPuppetRemoteRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *PuppetRemoteRepositoryCloner {
	return &PuppetRemoteRepositoryCloner{
		NewRemoteRepositoryClonerImpl(Puppet, source, target, dryRun)}
}

func (o *PuppetRemoteRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.PuppetRemoteRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateRemoteRepository().Puppet(sourceRepoDetails)
		})
	}
	return
}

type PyPIRemoteRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewPyPiRemoteRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *PyPIRemoteRepositoryCloner {
	return &PyPIRemoteRepositoryCloner{
		NewRemoteRepositoryClonerImpl(PyPi, source, target, dryRun)}
}

func (o *PyPIRemoteRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.PypiRemoteRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateRemoteRepository().Pypi(sourceRepoDetails)
		})
	}
	return
}

type RubyRemoteGemsRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewRubyRemoteGemsRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *RubyRemoteGemsRepositoryCloner {
	return &RubyRemoteGemsRepositoryCloner{
		NewRemoteRepositoryClonerImpl(RubyGems, source, target, dryRun)}
}

func (o *RubyRemoteGemsRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.GemsRemoteRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateRemoteRepository().Gems(sourceRepoDetails)
		})
	}
	return
}

type GenericRemoteRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewGenericRemoteRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *GenericRemoteRepositoryCloner {
	return &GenericRemoteRepositoryCloner{
		NewRemoteRepositoryClonerImpl(Generic, source, target, dryRun)}
}

func (o *GenericRemoteRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.GenericRemoteRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateRemoteRepository().Generic(sourceRepoDetails)
		})
	}
	return
}

type MavenRemoteRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewMavenRemoteRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *MavenRemoteRepositoryCloner {
	return &MavenRemoteRepositoryCloner{
		NewRemoteRepositoryClonerImpl(Maven, source, target, dryRun)}
}

func (o *MavenRemoteRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.MavenRemoteRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateRemoteRepository().Maven(sourceRepoDetails)
		})
	}
	return
}

type HelmRemoteRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewHelmRemoteRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *HelmRemoteRepositoryCloner {
	return &HelmRemoteRepositoryCloner{
		NewRemoteRepositoryClonerImpl(Helm, source, target, dryRun)}
}

func (o *HelmRemoteRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.HelmRemoteRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateRemoteRepository().Helm(sourceRepoDetails)
		})
	}
	return
}

type GitLfsRemoteRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewGitLfsRemoteRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *GitLfsRemoteRepositoryCloner {
	return &GitLfsRemoteRepositoryCloner{
		NewRemoteRepositoryClonerImpl(GitLfs, source, target, dryRun)}
}

func (o *GitLfsRemoteRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.GitlfsRemoteRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateRemoteRepository().Gitlfs(sourceRepoDetails)
		})
	}
	return
}

type DebianRemoteRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewDebianRemoteRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *DebianRemoteRepositoryCloner {
	return &DebianRemoteRepositoryCloner{
		NewRemoteRepositoryClonerImpl(GitLfs, source, target, dryRun)}
}

func (o *DebianRemoteRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.DebianRemoteRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateRemoteRepository().Debian(sourceRepoDetails)
		})
	}
	return
}

type YumRemoteRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewYumRemoteRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *YumRemoteRepositoryCloner {
	return &YumRemoteRepositoryCloner{
		NewRemoteRepositoryClonerImpl(YUM, source, target, dryRun)}
}

func (o *YumRemoteRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.YumRemoteRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateRemoteRepository().Yum(sourceRepoDetails)
		})
	}
	return
}

type CargoRemoteRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewCargoRemoteRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *CargoRemoteRepositoryCloner {
	return &CargoRemoteRepositoryCloner{
		NewRemoteRepositoryClonerImpl(Cargo, source, target, dryRun)}
}

func (o *CargoRemoteRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.CargoRemoteRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateRemoteRepository().Cargo(sourceRepoDetails)
		})
	}
	return
}

type GradleRemoteRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewGradleRemoteRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *GradleRemoteRepositoryCloner {
	return &GradleRemoteRepositoryCloner{
		NewRemoteRepositoryClonerImpl(Gradle, source, target, dryRun)}
}

func (o *GradleRemoteRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.GradleRemoteRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateRemoteRepository().Gradle(sourceRepoDetails)
		})
	}
	return
}
