package core

import (
	"fmt"
	"github.com/jfrog/jfrog-client-go/artifactory/services"
)

func BuildLocalRepoCloner(packageType PackageType,
	source *ArtifactoryManager, target *ArtifactoryManager) (ret RepositoryCloner, err error) {

	switch packageType {
	case Bower:
		ret = NewBowerLocalRepositoryCloner(source, target)
	case Chef:
		ret = NewChefLocalRepositoryCloner(source, target)
	case CocoaPods:
		ret = NewCocoaPodsLocalPodsRepositoryCloner(source, target)
	case Conan:
		ret = NewConanLocalRepositoryCloner(source, target)
	case Docker:
		ret = NewDockerLocalRepositoryCloner(source, target)
	case Go:
		ret = NewGoLocalRepositoryCloner(source, target)
	case NuGet:
		ret = NewNuGetLocalRepositoryCloner(source, target)
	case Npm:
		ret = NewNpmLocalRepositoryCloner(source, target)
	case PhpComposer:
		ret = NewPhpComposerLocalRepositoryCloner(source, target)
	case Puppet:
		ret = NewPuppetLocalRepositoryCloner(source, target)
	case PyPi:
		ret = NewPyPiLocalRepositoryCloner(source, target)
	case RubyGems:
		ret = NewRubyLocalGemsRepositoryCloner(source, target)
	case Generic:
		ret = NewGenericLocalRepositoryCloner(source, target)
	case Maven:
		ret = NewMavenLocalRepositoryCloner(source, target)
	case Helm:
		ret = NewHelmLocalRepositoryCloner(source, target)
	case GitLfs:
		ret = NewGitLfsLocalRepositoryCloner(source, target)
	case Debian:
		ret = NewDebianLocalRepositoryCloner(source, target)
	case YUM:
		ret = NewYumLocalRepositoryCloner(source, target)
	case Vagrant:
		ret = NewVagrantLocalRepositoryCloner(source, target)
	case Cargo:
		ret = NewCargoLocalRepositoryCloner(source, target)
	case Gradle:
		ret = NewGradleLocalRepositoryCloner(source, target)
	default:
		err = fmt.Errorf("repo type '%v' with package type '%v' not supported", Local, packageType)

	}
	return
}

func NewLocalRepositoryClonerImpl(packageType PackageType,
	source *ArtifactoryManager, target *ArtifactoryManager) *RepositoryClonerImpl {
	return NewRepositoryCloner(Local, packageType, source, target)
}

type BowerLocalRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewBowerLocalRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *BowerLocalRepositoryCloner {
	return &BowerLocalRepositoryCloner{
		NewLocalRepositoryClonerImpl(Bower, source, target)}
}

func (o *BowerLocalRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.BowerLocalRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			return o.Target.CreateLocalRepository().Bower(sourceRepoDetails)
		})
	}
	return
}

type ChefLocalRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewChefLocalRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *ChefLocalRepositoryCloner {
	return &ChefLocalRepositoryCloner{
		NewLocalRepositoryClonerImpl(Chef, source, target)}
}

func (o *ChefLocalRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.ChefLocalRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			return o.Target.CreateLocalRepository().Chef(sourceRepoDetails)
		})
	}
	return
}

type CocoaPodsLocalPodsRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewCocoaPodsLocalPodsRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *CocoaPodsLocalPodsRepositoryCloner {
	return &CocoaPodsLocalPodsRepositoryCloner{
		NewLocalRepositoryClonerImpl(CocoaPods, source, target)}
}

func (o *CocoaPodsLocalPodsRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.CocoapodsLocalRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			return o.Target.CreateLocalRepository().Cocoapods(sourceRepoDetails)
		})
	}
	return
}

type ConanLocalRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewConanLocalRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *ConanLocalRepositoryCloner {
	return &ConanLocalRepositoryCloner{
		NewLocalRepositoryClonerImpl(Conan, source, target)}
}

func (o *ConanLocalRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.ConanLocalRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			return o.Target.CreateLocalRepository().Conan(sourceRepoDetails)
		})
	}
	return
}

type DockerLocalRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewDockerLocalRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *DockerLocalRepositoryCloner {
	return &DockerLocalRepositoryCloner{
		NewLocalRepositoryClonerImpl(Docker, source, target)}
}

func (o *DockerLocalRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.DockerLocalRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			return o.Target.CreateLocalRepository().Docker(sourceRepoDetails)
		})
	}
	return
}

type GoLocalRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewGoLocalRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *GoLocalRepositoryCloner {
	return &GoLocalRepositoryCloner{
		NewLocalRepositoryClonerImpl(Go, source, target)}
}

func (o *GoLocalRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.GoLocalRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			return o.Target.CreateLocalRepository().Go(sourceRepoDetails)
		})
	}
	return
}

type NuGetLocalRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewNuGetLocalRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *NuGetLocalRepositoryCloner {
	return &NuGetLocalRepositoryCloner{
		NewLocalRepositoryClonerImpl(NuGet, source, target)}
}

func (o *NuGetLocalRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.NugetLocalRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			return o.Target.CreateLocalRepository().Nuget(sourceRepoDetails)
		})
	}
	return
}

type NpmLocalRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewNpmLocalRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *NpmLocalRepositoryCloner {
	return &NpmLocalRepositoryCloner{
		NewLocalRepositoryClonerImpl(Npm, source, target)}
}

func (o *NpmLocalRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.NpmLocalRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			return o.Target.CreateLocalRepository().Npm(sourceRepoDetails)
		})
	}
	return
}

type PhpComposerLocalRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewPhpComposerLocalRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *PhpComposerLocalRepositoryCloner {
	return &PhpComposerLocalRepositoryCloner{
		NewLocalRepositoryClonerImpl(PhpComposer, source, target)}
}

func (o *PhpComposerLocalRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.ComposerLocalRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			return o.Target.CreateLocalRepository().Composer(sourceRepoDetails)
		})
	}
	return
}

type PuppetLocalRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewPuppetLocalRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *PuppetLocalRepositoryCloner {
	return &PuppetLocalRepositoryCloner{
		NewLocalRepositoryClonerImpl(Puppet, source, target)}
}

func (o *PuppetLocalRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.PuppetLocalRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			return o.Target.CreateLocalRepository().Puppet(sourceRepoDetails)
		})
	}
	return
}

type PypiLocalRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewPyPiLocalRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *PypiLocalRepositoryCloner {
	return &PypiLocalRepositoryCloner{
		NewLocalRepositoryClonerImpl(PyPi, source, target)}
}

func (o *PypiLocalRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.PypiLocalRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			return o.Target.CreateLocalRepository().Pypi(sourceRepoDetails)
		})
	}
	return
}

type RubyLocalGemsRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewRubyLocalGemsRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *RubyLocalGemsRepositoryCloner {
	return &RubyLocalGemsRepositoryCloner{
		NewLocalRepositoryClonerImpl(RubyGems, source, target)}
}

func (o *RubyLocalGemsRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.GemsLocalRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			return o.Target.CreateLocalRepository().Gems(sourceRepoDetails)
		})
	}
	return
}

type GenericLocalRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewGenericLocalRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *GenericLocalRepositoryCloner {
	return &GenericLocalRepositoryCloner{
		NewLocalRepositoryClonerImpl(Generic, source, target)}
}

func (o *GenericLocalRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.GenericLocalRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			return o.Target.CreateLocalRepository().Generic(sourceRepoDetails)
		})
	}
	return
}

type MavenLocalRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewMavenLocalRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *MavenLocalRepositoryCloner {
	return &MavenLocalRepositoryCloner{
		NewLocalRepositoryClonerImpl(Maven, source, target)}
}

func (o *MavenLocalRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.MavenLocalRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			return o.Target.CreateLocalRepository().Maven(sourceRepoDetails)
		})
	}
	return
}

type HelmLocalRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewHelmLocalRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *HelmLocalRepositoryCloner {
	return &HelmLocalRepositoryCloner{
		NewLocalRepositoryClonerImpl(Helm, source, target)}
}

func (o *HelmLocalRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.HelmLocalRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			return o.Target.CreateLocalRepository().Helm(sourceRepoDetails)
		})
	}
	return
}

type GitLfsLocalRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewGitLfsLocalRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *GitLfsLocalRepositoryCloner {
	return &GitLfsLocalRepositoryCloner{
		NewLocalRepositoryClonerImpl(GitLfs, source, target)}
}

func (o *GitLfsLocalRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.GitlfsLocalRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			return o.Target.CreateLocalRepository().Gitlfs(sourceRepoDetails)
		})
	}
	return
}

type DebianLocalRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewDebianLocalRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *DebianLocalRepositoryCloner {
	return &DebianLocalRepositoryCloner{
		NewLocalRepositoryClonerImpl(GitLfs, source, target)}
}

func (o *DebianLocalRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.DebianLocalRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			return o.Target.CreateLocalRepository().Debian(sourceRepoDetails)
		})
	}
	return
}

type YumLocalRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewYumLocalRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *YumLocalRepositoryCloner {
	return &YumLocalRepositoryCloner{
		NewLocalRepositoryClonerImpl(YUM, source, target)}
}

func (o *YumLocalRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.YumLocalRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			return o.Target.CreateLocalRepository().Yum(sourceRepoDetails)
		})
	}
	return
}

type VagrantLocalRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewVagrantLocalRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *VagrantLocalRepositoryCloner {
	return &VagrantLocalRepositoryCloner{
		NewLocalRepositoryClonerImpl(YUM, source, target)}
}

func (o *VagrantLocalRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.VagrantLocalRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			return o.Target.CreateLocalRepository().Vagrant(sourceRepoDetails)
		})
	}
	return
}

type CargoLocalRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewCargoLocalRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *CargoLocalRepositoryCloner {
	return &CargoLocalRepositoryCloner{
		NewLocalRepositoryClonerImpl(Cargo, source, target)}
}

func (o *CargoLocalRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.CargoLocalRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			return o.Target.CreateLocalRepository().Cargo(sourceRepoDetails)
		})
	}
	return
}

type GradleLocalRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewGradleLocalRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *GradleLocalRepositoryCloner {
	return &GradleLocalRepositoryCloner{
		NewLocalRepositoryClonerImpl(Gradle, source, target)}
}

func (o *GradleLocalRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.GradleLocalRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			return o.Target.CreateLocalRepository().Gradle(sourceRepoDetails)
		})
	}
	return
}
