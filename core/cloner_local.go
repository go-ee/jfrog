package core

import (
	"fmt"
	"github.com/jfrog/jfrog-client-go/artifactory/services"
)

func BuildLocalRepoCloner(packageType PackageType,
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) (ret RepositoryCloner, err error) {

	switch packageType {
	case Bower:
		ret = NewBowerLocalRepositoryCloner(source, target, dryRun)
	case Chef:
		ret = NewChefLocalRepositoryCloner(source, target, dryRun)
	case CocoaPods:
		ret = NewCocoaPodsLocalPodsRepositoryCloner(source, target, dryRun)
	case Conan:
		ret = NewConanLocalRepositoryCloner(source, target, dryRun)
	case Docker:
		ret = NewDockerLocalRepositoryCloner(source, target, dryRun)
	case Go:
		ret = NewGoLocalRepositoryCloner(source, target, dryRun)
	case NuGet:
		ret = NewNuGetLocalRepositoryCloner(source, target, dryRun)
	case Npm:
		ret = NewNpmLocalRepositoryCloner(source, target, dryRun)
	case PhpComposer:
		ret = NewPhpComposerLocalRepositoryCloner(source, target, dryRun)
	case Puppet:
		ret = NewPuppetLocalRepositoryCloner(source, target, dryRun)
	case PyPi:
		ret = NewPyPiLocalRepositoryCloner(source, target, dryRun)
	case RubyGems:
		ret = NewRubyLocalGemsRepositoryCloner(source, target, dryRun)
	case Generic:
		ret = NewGenericLocalRepositoryCloner(source, target, dryRun)
	case Maven:
		ret = NewMavenLocalRepositoryCloner(source, target, dryRun)
	case Helm:
		ret = NewHelmLocalRepositoryCloner(source, target, dryRun)
	case GitLfs:
		ret = NewGitLfsLocalRepositoryCloner(source, target, dryRun)
	case Debian:
		ret = NewDebianLocalRepositoryCloner(source, target, dryRun)
	case YUM:
		ret = NewYumLocalRepositoryCloner(source, target, dryRun)
	case Vagrant:
		ret = NewVagrantLocalRepositoryCloner(source, target, dryRun)
	case Cargo:
		ret = NewCargoLocalRepositoryCloner(source, target, dryRun)
	case Gradle:
		ret = NewGradleLocalRepositoryCloner(source, target, dryRun)
	default:
		err = fmt.Errorf("repo type '%v' with package type '%v' not supported", Local, packageType)

	}
	return
}

func NewLocalRepositoryClonerImpl(packageType PackageType,
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *RepositoryClonerImpl {
	return &RepositoryClonerImpl{RepoType: Local, PackageType: packageType,
		Source: source, Target: target, DryRun: dryRun}
}

type BowerLocalRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewBowerLocalRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *BowerLocalRepositoryCloner {
	return &BowerLocalRepositoryCloner{
		NewLocalRepositoryClonerImpl(Bower, source, target, dryRun)}
}

func (o *BowerLocalRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.BowerLocalRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateLocalRepository().Bower(sourceRepoDetails)
		})
	}
	return
}

type ChefLocalRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewChefLocalRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *ChefLocalRepositoryCloner {
	return &ChefLocalRepositoryCloner{
		NewLocalRepositoryClonerImpl(Chef, source, target, dryRun)}
}

func (o *ChefLocalRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.ChefLocalRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateLocalRepository().Chef(sourceRepoDetails)
		})
	}
	return
}

type CocoaPodsLocalPodsRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewCocoaPodsLocalPodsRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *CocoaPodsLocalPodsRepositoryCloner {
	return &CocoaPodsLocalPodsRepositoryCloner{
		NewLocalRepositoryClonerImpl(CocoaPods, source, target, dryRun)}
}

func (o *CocoaPodsLocalPodsRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.CocoapodsLocalRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateLocalRepository().Cocoapods(sourceRepoDetails)
		})
	}
	return
}

type ConanLocalRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewConanLocalRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *ConanLocalRepositoryCloner {
	return &ConanLocalRepositoryCloner{
		NewLocalRepositoryClonerImpl(Conan, source, target, dryRun)}
}

func (o *ConanLocalRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.ConanLocalRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateLocalRepository().Conan(sourceRepoDetails)
		})
	}
	return
}

type DockerLocalRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewDockerLocalRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *DockerLocalRepositoryCloner {
	return &DockerLocalRepositoryCloner{
		NewLocalRepositoryClonerImpl(Docker, source, target, dryRun)}
}

func (o *DockerLocalRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.DockerLocalRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateLocalRepository().Docker(sourceRepoDetails)
		})
	}
	return
}

type GoLocalRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewGoLocalRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *GoLocalRepositoryCloner {
	return &GoLocalRepositoryCloner{
		NewLocalRepositoryClonerImpl(Go, source, target, dryRun)}
}

func (o *GoLocalRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.GoLocalRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateLocalRepository().Go(sourceRepoDetails)
		})
	}
	return
}

type NuGetLocalRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewNuGetLocalRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *NuGetLocalRepositoryCloner {
	return &NuGetLocalRepositoryCloner{
		NewLocalRepositoryClonerImpl(NuGet, source, target, dryRun)}
}

func (o *NuGetLocalRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.NugetLocalRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateLocalRepository().Nuget(sourceRepoDetails)
		})
	}
	return
}

type NpmLocalRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewNpmLocalRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *NpmLocalRepositoryCloner {
	return &NpmLocalRepositoryCloner{
		NewLocalRepositoryClonerImpl(Npm, source, target, dryRun)}
}

func (o *NpmLocalRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.NpmLocalRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateLocalRepository().Npm(sourceRepoDetails)
		})
	}
	return
}

type PhpComposerLocalRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewPhpComposerLocalRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *PhpComposerLocalRepositoryCloner {
	return &PhpComposerLocalRepositoryCloner{
		NewLocalRepositoryClonerImpl(PhpComposer, source, target, dryRun)}
}

func (o *PhpComposerLocalRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.ComposerLocalRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateLocalRepository().Composer(sourceRepoDetails)
		})
	}
	return
}

type PuppetLocalRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewPuppetLocalRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *PuppetLocalRepositoryCloner {
	return &PuppetLocalRepositoryCloner{
		NewLocalRepositoryClonerImpl(Puppet, source, target, dryRun)}
}

func (o *PuppetLocalRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.PuppetLocalRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateLocalRepository().Puppet(sourceRepoDetails)
		})
	}
	return
}

type PypiLocalRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewPyPiLocalRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *PypiLocalRepositoryCloner {
	return &PypiLocalRepositoryCloner{
		NewLocalRepositoryClonerImpl(PyPi, source, target, dryRun)}
}

func (o *PypiLocalRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.PypiLocalRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateLocalRepository().Pypi(sourceRepoDetails)
		})
	}
	return
}

type RubyLocalGemsRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewRubyLocalGemsRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *RubyLocalGemsRepositoryCloner {
	return &RubyLocalGemsRepositoryCloner{
		NewLocalRepositoryClonerImpl(RubyGems, source, target, dryRun)}
}

func (o *RubyLocalGemsRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.GemsLocalRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateLocalRepository().Gems(sourceRepoDetails)
		})
	}
	return
}

type GenericLocalRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewGenericLocalRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *GenericLocalRepositoryCloner {
	return &GenericLocalRepositoryCloner{
		NewLocalRepositoryClonerImpl(Generic, source, target, dryRun)}
}

func (o *GenericLocalRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.GenericLocalRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateLocalRepository().Generic(sourceRepoDetails)
		})
	}
	return
}

type MavenLocalRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewMavenLocalRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *MavenLocalRepositoryCloner {
	return &MavenLocalRepositoryCloner{
		NewLocalRepositoryClonerImpl(Maven, source, target, dryRun)}
}

func (o *MavenLocalRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.MavenLocalRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateLocalRepository().Maven(sourceRepoDetails)
		})
	}
	return
}

type HelmLocalRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewHelmLocalRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *HelmLocalRepositoryCloner {
	return &HelmLocalRepositoryCloner{
		NewLocalRepositoryClonerImpl(Helm, source, target, dryRun)}
}

func (o *HelmLocalRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.HelmLocalRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateLocalRepository().Helm(sourceRepoDetails)
		})
	}
	return
}

type GitLfsLocalRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewGitLfsLocalRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *GitLfsLocalRepositoryCloner {
	return &GitLfsLocalRepositoryCloner{
		NewLocalRepositoryClonerImpl(GitLfs, source, target, dryRun)}
}

func (o *GitLfsLocalRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.GitlfsLocalRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateLocalRepository().Gitlfs(sourceRepoDetails)
		})
	}
	return
}

type DebianLocalRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewDebianLocalRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *DebianLocalRepositoryCloner {
	return &DebianLocalRepositoryCloner{
		NewLocalRepositoryClonerImpl(GitLfs, source, target, dryRun)}
}

func (o *DebianLocalRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.DebianLocalRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateLocalRepository().Debian(sourceRepoDetails)
		})
	}
	return
}

type YumLocalRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewYumLocalRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *YumLocalRepositoryCloner {
	return &YumLocalRepositoryCloner{
		NewLocalRepositoryClonerImpl(YUM, source, target, dryRun)}
}

func (o *YumLocalRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.YumLocalRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateLocalRepository().Yum(sourceRepoDetails)
		})
	}
	return
}

type VagrantLocalRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewVagrantLocalRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *VagrantLocalRepositoryCloner {
	return &VagrantLocalRepositoryCloner{
		NewLocalRepositoryClonerImpl(YUM, source, target, dryRun)}
}

func (o *VagrantLocalRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.VagrantLocalRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateLocalRepository().Vagrant(sourceRepoDetails)
		})
	}
	return
}

type CargoLocalRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewCargoLocalRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *CargoLocalRepositoryCloner {
	return &CargoLocalRepositoryCloner{
		NewLocalRepositoryClonerImpl(Cargo, source, target, dryRun)}
}

func (o *CargoLocalRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.CargoLocalRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateLocalRepository().Cargo(sourceRepoDetails)
		})
	}
	return
}

type GradleLocalRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewGradleLocalRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *GradleLocalRepositoryCloner {
	return &GradleLocalRepositoryCloner{
		NewLocalRepositoryClonerImpl(Gradle, source, target, dryRun)}
}

func (o *GradleLocalRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.GradleLocalRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateLocalRepository().Gradle(sourceRepoDetails)
		})
	}
	return
}
