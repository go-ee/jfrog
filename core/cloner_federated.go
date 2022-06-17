package core

import (
	"fmt"
	"github.com/jfrog/jfrog-client-go/artifactory/services"
)

func BuildFederatedRepoCloner(packageType PackageType,
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) (ret RepositoryCloner, err error) {

	switch packageType {
	case Bower:
		ret = NewBowerFederatedRepositoryCloner(source, target, dryRun)
	case Chef:
		ret = NewChefFederatedRepositoryCloner(source, target, dryRun)
	case CocoaPods:
		ret = NewCocoaPodsFederatedPodsRepositoryCloner(source, target, dryRun)
	case Conan:
		ret = NewConanFederatedRepositoryCloner(source, target, dryRun)
	case Docker:
		ret = NewDockerFederatedRepositoryCloner(source, target, dryRun)
	case Go:
		ret = NewGoFederatedRepositoryCloner(source, target, dryRun)
	case NuGet:
		ret = NewNuGetFederatedRepositoryCloner(source, target, dryRun)
	case Npm:
		ret = NewNpmFederatedRepositoryCloner(source, target, dryRun)
	case PhpComposer:
		ret = NewPhpComposerFederatedRepositoryCloner(source, target, dryRun)
	case Puppet:
		ret = NewPuppetFederatedRepositoryCloner(source, target, dryRun)
	case PyPi:
		ret = NewPyPiFederatedRepositoryCloner(source, target, dryRun)
	case RubyGems:
		ret = NewRubyFederatedGemsRepositoryCloner(source, target, dryRun)
	case Generic:
		ret = NewGenericFederatedRepositoryCloner(source, target, dryRun)
	case Maven:
		ret = NewMavenFederatedRepositoryCloner(source, target, dryRun)
	case Helm:
		ret = NewHelmFederatedRepositoryCloner(source, target, dryRun)
	case GitLfs:
		ret = NewGitLfsFederatedRepositoryCloner(source, target, dryRun)
	case Debian:
		ret = NewDebianFederatedRepositoryCloner(source, target, dryRun)
	case YUM:
		ret = NewYumFederatedRepositoryCloner(source, target, dryRun)
	case Vagrant:
		ret = NewVagrantFederatedRepositoryCloner(source, target, dryRun)
	case Cargo:
		ret = NewCargoFederatedRepositoryCloner(source, target, dryRun)
	case Gradle:
		ret = NewGradleFederatedRepositoryCloner(source, target, dryRun)
	default:
		err = fmt.Errorf("repo type '%v' with package type '%v' not supported", Federated, packageType)

	}
	return
}

func NewFederatedRepositoryClonerImpl(packageType PackageType,
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *RepositoryClonerImpl {
	return &RepositoryClonerImpl{RepoType: Federated, PackageType: packageType,
		Source: source, Target: target, DryRun: dryRun}
}

type BowerFederatedRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewBowerFederatedRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *BowerFederatedRepositoryCloner {
	return &BowerFederatedRepositoryCloner{
		NewFederatedRepositoryClonerImpl(Bower, source, target, dryRun)}
}

func (o *BowerFederatedRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.BowerFederatedRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateFederatedRepository().Bower(sourceRepoDetails)
		})
	}
	return
}

type ChefFederatedRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewChefFederatedRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *ChefFederatedRepositoryCloner {
	return &ChefFederatedRepositoryCloner{
		NewFederatedRepositoryClonerImpl(Chef, source, target, dryRun)}
}

func (o *ChefFederatedRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.ChefFederatedRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateFederatedRepository().Chef(sourceRepoDetails)
		})
	}
	return
}

type CocoaPodsFederatedPodsRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewCocoaPodsFederatedPodsRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *CocoaPodsFederatedPodsRepositoryCloner {
	return &CocoaPodsFederatedPodsRepositoryCloner{
		NewFederatedRepositoryClonerImpl(CocoaPods, source, target, dryRun)}
}

func (o *CocoaPodsFederatedPodsRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.CocoapodsFederatedRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateFederatedRepository().Cocoapods(sourceRepoDetails)
		})
	}
	return
}

type ConanFederatedRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewConanFederatedRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *ConanFederatedRepositoryCloner {
	return &ConanFederatedRepositoryCloner{
		NewFederatedRepositoryClonerImpl(Conan, source, target, dryRun)}
}

func (o *ConanFederatedRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.ConanFederatedRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateFederatedRepository().Conan(sourceRepoDetails)
		})
	}
	return
}

type DockerFederatedRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewDockerFederatedRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *DockerFederatedRepositoryCloner {
	return &DockerFederatedRepositoryCloner{
		NewFederatedRepositoryClonerImpl(Docker, source, target, dryRun)}
}

func (o *DockerFederatedRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.DockerFederatedRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateFederatedRepository().Docker(sourceRepoDetails)
		})
	}
	return
}

type GoFederatedRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewGoFederatedRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *GoFederatedRepositoryCloner {
	return &GoFederatedRepositoryCloner{
		NewFederatedRepositoryClonerImpl(Go, source, target, dryRun)}
}

func (o *GoFederatedRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.GoFederatedRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateFederatedRepository().Go(sourceRepoDetails)
		})
	}
	return
}

type NuGetFederatedRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewNuGetFederatedRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *NuGetFederatedRepositoryCloner {
	return &NuGetFederatedRepositoryCloner{
		NewFederatedRepositoryClonerImpl(NuGet, source, target, dryRun)}
}

func (o *NuGetFederatedRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.NugetFederatedRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateFederatedRepository().Nuget(sourceRepoDetails)
		})
	}
	return
}

type NpmFederatedRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewNpmFederatedRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *NpmFederatedRepositoryCloner {
	return &NpmFederatedRepositoryCloner{
		NewFederatedRepositoryClonerImpl(Npm, source, target, dryRun)}
}

func (o *NpmFederatedRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.NpmFederatedRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateFederatedRepository().Npm(sourceRepoDetails)
		})
	}
	return
}

type PhpComposerFederatedRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewPhpComposerFederatedRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *PhpComposerFederatedRepositoryCloner {
	return &PhpComposerFederatedRepositoryCloner{
		NewFederatedRepositoryClonerImpl(PhpComposer, source, target, dryRun)}
}

func (o *PhpComposerFederatedRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.ComposerFederatedRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateFederatedRepository().Composer(sourceRepoDetails)
		})
	}
	return
}

type PuppetFederatedRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewPuppetFederatedRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *PuppetFederatedRepositoryCloner {
	return &PuppetFederatedRepositoryCloner{
		NewFederatedRepositoryClonerImpl(Puppet, source, target, dryRun)}
}

func (o *PuppetFederatedRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.PuppetFederatedRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateFederatedRepository().Puppet(sourceRepoDetails)
		})
	}
	return
}

type PyPiFederatedRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewPyPiFederatedRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *PyPiFederatedRepositoryCloner {
	return &PyPiFederatedRepositoryCloner{
		NewFederatedRepositoryClonerImpl(PyPi, source, target, dryRun)}
}

func (o *PyPiFederatedRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.PuppetFederatedRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateFederatedRepository().Puppet(sourceRepoDetails)
		})
	}
	return
}

type RubyFederatedGemsRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewRubyFederatedGemsRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *RubyFederatedGemsRepositoryCloner {
	return &RubyFederatedGemsRepositoryCloner{
		NewFederatedRepositoryClonerImpl(RubyGems, source, target, dryRun)}
}

func (o *RubyFederatedGemsRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.GemsFederatedRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateFederatedRepository().Gems(sourceRepoDetails)
		})
	}
	return
}

type GenericFederatedRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewGenericFederatedRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *GenericFederatedRepositoryCloner {
	return &GenericFederatedRepositoryCloner{
		NewFederatedRepositoryClonerImpl(Generic, source, target, dryRun)}
}

func (o *GenericFederatedRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.GenericFederatedRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateFederatedRepository().Generic(sourceRepoDetails)
		})
	}
	return
}

type MavenFederatedRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewMavenFederatedRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *MavenFederatedRepositoryCloner {
	return &MavenFederatedRepositoryCloner{
		NewFederatedRepositoryClonerImpl(Maven, source, target, dryRun)}
}

func (o *MavenFederatedRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.MavenFederatedRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateFederatedRepository().Maven(sourceRepoDetails)
		})
	}
	return
}

type HelmFederatedRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewHelmFederatedRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *HelmFederatedRepositoryCloner {
	return &HelmFederatedRepositoryCloner{
		NewFederatedRepositoryClonerImpl(Helm, source, target, dryRun)}
}

func (o *HelmFederatedRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.HelmFederatedRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateFederatedRepository().Helm(sourceRepoDetails)
		})
	}
	return
}

type GitLfsFederatedRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewGitLfsFederatedRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *GitLfsFederatedRepositoryCloner {
	return &GitLfsFederatedRepositoryCloner{
		NewFederatedRepositoryClonerImpl(GitLfs, source, target, dryRun)}
}

func (o *GitLfsFederatedRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.GitlfsFederatedRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateFederatedRepository().Gitlfs(sourceRepoDetails)
		})
	}
	return
}

type DebianFederatedRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewDebianFederatedRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *DebianFederatedRepositoryCloner {
	return &DebianFederatedRepositoryCloner{
		NewFederatedRepositoryClonerImpl(GitLfs, source, target, dryRun)}
}

func (o *DebianFederatedRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.DebianFederatedRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateFederatedRepository().Debian(sourceRepoDetails)
		})
	}
	return
}

type YumFederatedRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewYumFederatedRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *YumFederatedRepositoryCloner {
	return &YumFederatedRepositoryCloner{
		NewFederatedRepositoryClonerImpl(YUM, source, target, dryRun)}
}

func (o *YumFederatedRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.YumFederatedRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateFederatedRepository().Yum(sourceRepoDetails)
		})
	}
	return
}

type VagrantFederatedRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewVagrantFederatedRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *VagrantFederatedRepositoryCloner {
	return &VagrantFederatedRepositoryCloner{
		NewFederatedRepositoryClonerImpl(YUM, source, target, dryRun)}
}

func (o *VagrantFederatedRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.VagrantFederatedRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateFederatedRepository().Vagrant(sourceRepoDetails)
		})
	}
	return
}

type CargoFederatedRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewCargoFederatedRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *CargoFederatedRepositoryCloner {
	return &CargoFederatedRepositoryCloner{
		NewFederatedRepositoryClonerImpl(Cargo, source, target, dryRun)}
}

func (o *CargoFederatedRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.CargoFederatedRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateFederatedRepository().Cargo(sourceRepoDetails)
		})
	}
	return
}

type GradleFederatedRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewGradleFederatedRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager, dryRun bool) *GradleFederatedRepositoryCloner {
	return &GradleFederatedRepositoryCloner{
		NewFederatedRepositoryClonerImpl(Gradle, source, target, dryRun)}
}

func (o *GradleFederatedRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.GradleFederatedRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.clone(repoKey, func() error {
			return o.Target.CreateFederatedRepository().Gradle(sourceRepoDetails)
		})
	}
	return
}
