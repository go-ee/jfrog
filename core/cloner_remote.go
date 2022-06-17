package core

import (
	"fmt"
	"github.com/jfrog/jfrog-client-go/artifactory/services"
)

func BuildRemoteRepoCloner(packageType PackageType,
	source *ArtifactoryManager, target *ArtifactoryManager) (ret RepositoryCloner, err error) {

	switch packageType {
	case Bower:
		ret = NewBowerRemoteRepositoryCloner(source, target)
	case Chef:
		ret = NewChefRemoteRepositoryCloner(source, target)
	case CocoaPods:
		ret = NewCocoaPodsRemotePodsRepositoryCloner(source, target)
	case Conan:
		ret = NewConanRemoteRepositoryCloner(source, target)
	case Docker:
		ret = NewDockerRemoteRepositoryCloner(source, target)
	case Go:
		ret = NewGoRemoteRepositoryCloner(source, target)
	case NuGet:
		ret = NewNuGetRemoteRepositoryCloner(source, target)
	case Npm:
		ret = NewNpmRemoteRepositoryCloner(source, target)
	case PhpComposer:
		ret = NewPhpComposerRemoteRepositoryCloner(source, target)
	case Puppet:
		ret = NewPuppetRemoteRepositoryCloner(source, target)
	case PyPi:
		ret = NewPyPiRemoteRepositoryCloner(source, target)
	case RubyGems:
		ret = NewRubyRemoteGemsRepositoryCloner(source, target)
	case Generic:
		ret = NewGenericRemoteRepositoryCloner(source, target)
	case Maven:
		ret = NewMavenRemoteRepositoryCloner(source, target)
	case Helm:
		ret = NewHelmRemoteRepositoryCloner(source, target)
	case GitLfs:
		ret = NewGitLfsRemoteRepositoryCloner(source, target)
	case Debian:
		ret = NewDebianRemoteRepositoryCloner(source, target)
	case YUM:
		ret = NewYumRemoteRepositoryCloner(source, target)
	case Cargo:
		ret = NewCargoRemoteRepositoryCloner(source, target)
	case Gradle:
		ret = NewGradleRemoteRepositoryCloner(source, target)
	default:
		err = fmt.Errorf("repo type '%v' with package type '%v' not supported", Remote, packageType)

	}
	return
}

func NewRemoteRepositoryClonerImpl(packageType PackageType,
	source *ArtifactoryManager, target *ArtifactoryManager) *RepositoryClonerImpl {
	return NewRepositoryCloner(Remote, packageType, source, target)
}

type BowerRemoteRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewBowerRemoteRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *BowerRemoteRepositoryCloner {
	return &BowerRemoteRepositoryCloner{
		NewRemoteRepositoryClonerImpl(Bower, source, target)}
}

func (o *BowerRemoteRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.BowerRemoteRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			return o.Target.CreateRemoteRepository().Bower(sourceRepoDetails)
		})
	}
	return
}

type ChefRemoteRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewChefRemoteRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *ChefRemoteRepositoryCloner {
	return &ChefRemoteRepositoryCloner{
		NewRemoteRepositoryClonerImpl(Chef, source, target)}
}

func (o *ChefRemoteRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.ChefRemoteRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			return o.Target.CreateRemoteRepository().Chef(sourceRepoDetails)
		})
	}
	return
}

type CocoaPodsRemotePodsRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewCocoaPodsRemotePodsRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *CocoaPodsRemotePodsRepositoryCloner {
	return &CocoaPodsRemotePodsRepositoryCloner{
		NewRemoteRepositoryClonerImpl(CocoaPods, source, target)}
}

func (o *CocoaPodsRemotePodsRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.CocoapodsRemoteRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			return o.Target.CreateRemoteRepository().Cocoapods(sourceRepoDetails)
		})
	}
	return
}

type ConanRemoteRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewConanRemoteRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *ConanRemoteRepositoryCloner {
	return &ConanRemoteRepositoryCloner{
		NewRemoteRepositoryClonerImpl(Conan, source, target)}
}

func (o *ConanRemoteRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.ConanRemoteRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			return o.Target.CreateRemoteRepository().Conan(sourceRepoDetails)
		})
	}
	return
}

type DockerRemoteRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewDockerRemoteRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *DockerRemoteRepositoryCloner {
	return &DockerRemoteRepositoryCloner{
		NewRemoteRepositoryClonerImpl(Docker, source, target)}
}

func (o *DockerRemoteRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.DockerRemoteRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			return o.Target.CreateRemoteRepository().Docker(sourceRepoDetails)
		})
	}
	return
}

type GoRemoteRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewGoRemoteRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *GoRemoteRepositoryCloner {
	return &GoRemoteRepositoryCloner{
		NewRemoteRepositoryClonerImpl(Go, source, target)}
}

func (o *GoRemoteRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.GoRemoteRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			return o.Target.CreateRemoteRepository().Go(sourceRepoDetails)
		})
	}
	return
}

type NuGetRemoteRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewNuGetRemoteRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *NuGetRemoteRepositoryCloner {
	return &NuGetRemoteRepositoryCloner{
		NewRemoteRepositoryClonerImpl(NuGet, source, target)}
}

func (o *NuGetRemoteRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.NugetRemoteRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			return o.Target.CreateRemoteRepository().Nuget(sourceRepoDetails)
		})
	}
	return
}

type NpmRemoteRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewNpmRemoteRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *NpmRemoteRepositoryCloner {
	return &NpmRemoteRepositoryCloner{
		NewRemoteRepositoryClonerImpl(Npm, source, target)}
}

func (o *NpmRemoteRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.NpmRemoteRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			return o.Target.CreateRemoteRepository().Npm(sourceRepoDetails)
		})
	}
	return
}

type PhpComposerRemoteRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewPhpComposerRemoteRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *PhpComposerRemoteRepositoryCloner {
	return &PhpComposerRemoteRepositoryCloner{
		NewRemoteRepositoryClonerImpl(PhpComposer, source, target)}
}

func (o *PhpComposerRemoteRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.ComposerRemoteRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			return o.Target.CreateRemoteRepository().Composer(sourceRepoDetails)
		})
	}
	return
}

type PuppetRemoteRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewPuppetRemoteRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *PuppetRemoteRepositoryCloner {
	return &PuppetRemoteRepositoryCloner{
		NewRemoteRepositoryClonerImpl(Puppet, source, target)}
}

func (o *PuppetRemoteRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.PuppetRemoteRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			return o.Target.CreateRemoteRepository().Puppet(sourceRepoDetails)
		})
	}
	return
}

type PyPIRemoteRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewPyPiRemoteRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *PyPIRemoteRepositoryCloner {
	return &PyPIRemoteRepositoryCloner{
		NewRemoteRepositoryClonerImpl(PyPi, source, target)}
}

func (o *PyPIRemoteRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.PypiRemoteRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			return o.Target.CreateRemoteRepository().Pypi(sourceRepoDetails)
		})
	}
	return
}

type RubyRemoteGemsRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewRubyRemoteGemsRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *RubyRemoteGemsRepositoryCloner {
	return &RubyRemoteGemsRepositoryCloner{
		NewRemoteRepositoryClonerImpl(RubyGems, source, target)}
}

func (o *RubyRemoteGemsRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.GemsRemoteRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			return o.Target.CreateRemoteRepository().Gems(sourceRepoDetails)
		})
	}
	return
}

type GenericRemoteRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewGenericRemoteRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *GenericRemoteRepositoryCloner {
	return &GenericRemoteRepositoryCloner{
		NewRemoteRepositoryClonerImpl(Generic, source, target)}
}

func (o *GenericRemoteRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.GenericRemoteRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			return o.Target.CreateRemoteRepository().Generic(sourceRepoDetails)
		})
	}
	return
}

type MavenRemoteRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewMavenRemoteRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *MavenRemoteRepositoryCloner {
	return &MavenRemoteRepositoryCloner{
		NewRemoteRepositoryClonerImpl(Maven, source, target)}
}

func (o *MavenRemoteRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.MavenRemoteRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			return o.Target.CreateRemoteRepository().Maven(sourceRepoDetails)
		})
	}
	return
}

type HelmRemoteRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewHelmRemoteRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *HelmRemoteRepositoryCloner {
	return &HelmRemoteRepositoryCloner{
		NewRemoteRepositoryClonerImpl(Helm, source, target)}
}

func (o *HelmRemoteRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.HelmRemoteRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			return o.Target.CreateRemoteRepository().Helm(sourceRepoDetails)
		})
	}
	return
}

type GitLfsRemoteRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewGitLfsRemoteRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *GitLfsRemoteRepositoryCloner {
	return &GitLfsRemoteRepositoryCloner{
		NewRemoteRepositoryClonerImpl(GitLfs, source, target)}
}

func (o *GitLfsRemoteRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.GitlfsRemoteRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			return o.Target.CreateRemoteRepository().Gitlfs(sourceRepoDetails)
		})
	}
	return
}

type DebianRemoteRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewDebianRemoteRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *DebianRemoteRepositoryCloner {
	return &DebianRemoteRepositoryCloner{
		NewRemoteRepositoryClonerImpl(GitLfs, source, target)}
}

func (o *DebianRemoteRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.DebianRemoteRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			return o.Target.CreateRemoteRepository().Debian(sourceRepoDetails)
		})
	}
	return
}

type YumRemoteRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewYumRemoteRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *YumRemoteRepositoryCloner {
	return &YumRemoteRepositoryCloner{
		NewRemoteRepositoryClonerImpl(YUM, source, target)}
}

func (o *YumRemoteRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.YumRemoteRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			return o.Target.CreateRemoteRepository().Yum(sourceRepoDetails)
		})
	}
	return
}

type CargoRemoteRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewCargoRemoteRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *CargoRemoteRepositoryCloner {
	return &CargoRemoteRepositoryCloner{
		NewRemoteRepositoryClonerImpl(Cargo, source, target)}
}

func (o *CargoRemoteRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.CargoRemoteRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			return o.Target.CreateRemoteRepository().Cargo(sourceRepoDetails)
		})
	}
	return
}

type GradleRemoteRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewGradleRemoteRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *GradleRemoteRepositoryCloner {
	return &GradleRemoteRepositoryCloner{
		NewRemoteRepositoryClonerImpl(Gradle, source, target)}
}

func (o *GradleRemoteRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.GradleRemoteRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			return o.Target.CreateRemoteRepository().Gradle(sourceRepoDetails)
		})
	}
	return
}
